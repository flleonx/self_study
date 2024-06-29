package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"fss/p2p"
	"io"
	"log"
	"sync"
	"time"
)

type (
	FileServerOpts struct {
		ID                string
		EncKey            []byte
		StorageRoot       string
		PathTransformFunc PathTransformFunc
		Transport         p2p.Transport
		BootstrapNodes    []string
	}

	FileServer struct {
		FileServerOpts

		peerLock sync.Mutex
		peers    map[string]p2p.Peer

		store  *Store
		quitch chan struct{}
	}

	Message struct {
		Payload any
	}

	MessageStoreFile struct {
		ID   string
		Key  string
		Size int64
	}

	MessageGetFile struct {
		ID  string
		Key string
	}
)

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts{
		Root:              opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc,
	}

	if len(opts.ID) == 0 {
		opts.ID = generateID()
	}

	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(storeOpts),
		quitch:         make(chan struct{}),
		peers:          make(map[string]p2p.Peer),
	}
}

func (s *FileServer) broadcast(msg *Message) error {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(&msg); err != nil {
		fmt.Printf("error encoding: %s\n", err)
		return err
	}

	peers := []io.Writer{}
	for _, peer := range s.peers {
		peers = append(peers, peer)
	}

	mw := io.MultiWriter(peers...)
	if _, err := mw.Write([]byte{p2p.IncomingMessage}); err != nil {
		fmt.Printf("error sending flag: %s\n", err)
		return err
	}
	if _, err := mw.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (s *FileServer) Get(key string) (io.Reader, error) {
	if s.store.Has(s.ID, key) {
		fmt.Printf(
			"[%s] serving file (%s) from local disk\n",
			s.Transport.Addr(),
			key,
		)

		_, r, err := s.store.Read(s.ID, key)
		return r, err
	}

	fmt.Printf(
		"[%s] dont have file (%s) locally, fetching from network...\n",
		s.Transport.Addr(),
		key,
	)
	msg := Message{
		Payload: MessageGetFile{
			ID:  s.ID,
			Key: hashKey(key),
		},
	}

	if err := s.broadcast(&msg); err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	for _, peer := range s.peers {
		// First read the file size so we can limit the amount of bytes that
		// we read from the connection, so it will not keep hanging.

		var fileSize int64
		if err := binary.Read(peer, binary.LittleEndian, &fileSize); err != nil {
			fmt.Printf(
				"[%s] error decoding fileSize from (%s): %s\n",
				s.Transport.Addr(),
				peer.RemoteAddr(),
				err,
			)
		}

		n, err := s.store.WriteDecrypt(s.ID, s.EncKey, key, io.LimitReader(peer, fileSize))
		if err != nil {
			return nil, err
		}

		fmt.Printf(
			"[%s] received (%d) bytes over the network from (%s)\n",
			s.Transport.Addr(),
			n,
			peer.RemoteAddr(),
		)

		peer.CloseStream()
	}

	_, r, err := s.store.Read(s.ID, key)
	return r, err
}

func (s *FileServer) Store(key string, r io.Reader) error {
	// 1. Store this file to disk
	// 2. broadcast this file to all known peers in the network
	var (
		fileBuffer = new(bytes.Buffer)
		tee        = io.TeeReader(r, fileBuffer)
	)

	size, err := s.store.Write(s.ID, key, tee)
	if err != nil {
		return err
	}

	msg := Message{
		Payload: MessageStoreFile{
			ID:   s.ID,
			Key:  hashKey(key),
			Size: size + 16,
		},
	}

	if err := s.broadcast(&msg); err != nil {
		return err
	}

	peers := []io.Writer{}
	for _, peer := range s.peers {
		peers = append(peers, peer)
	}

	time.Sleep(5 * time.Millisecond)
	mw := io.MultiWriter(peers...)
	_, _ = mw.Write([]byte{p2p.IncomingStream})
	n, err := copyEncrypt(s.EncKey, fileBuffer, mw)
	if err != nil {
		return err
	}

	fmt.Printf(
		"[%s] received and written (%d) bytes to disk\n",
		s.Transport.Addr(),
		n,
	)

	return nil
}

func (s *FileServer) Stop() {
	close(s.quitch)
}

func (s *FileServer) OnPeer(p p2p.Peer) error {
	s.peerLock.Lock()
	defer s.peerLock.Unlock()

	s.peers[p.RemoteAddr().String()] = p
	log.Printf("connected with remote %s", p.RemoteAddr())
	return nil
}

func (s *FileServer) loop() {
	defer func() {
		log.Println("file server stopped due to error or user quit action")
		s.Transport.Close()
	}()

	for {
		select {
		case rpc := <-s.Transport.Consume():
			var msg Message
			if err := gob.NewDecoder(bytes.NewReader(rpc.Payload)).Decode(&msg); err != nil {
				log.Println("decoding error: ", err)
			}

			if err := s.handleMessage(rpc.From, &msg); err != nil {
				log.Println("handle message error: ", err)
			}
		case <-s.quitch:
			return
		}
	}
}

func (s *FileServer) handleMessage(from string, msg *Message) error {
	switch v := msg.Payload.(type) {
	case MessageStoreFile:
		return s.handleMessageStoreFile(from, v)
	case MessageGetFile:
		return s.handleMessageGetFile(from, v)
	}

	return nil
}

func (s *FileServer) handleMessageGetFile(from string, msg MessageGetFile) error {
	if !s.store.Has(msg.ID, msg.Key) {
		return fmt.Errorf(
			"[%s] need to serve file (%s) but it does not exist on disk",
			s.Transport.Addr(),
			msg.Key,
		)
	}

	fmt.Printf(
		"[%s] serving file (%s) over the network\n",
		s.Transport.Addr(),
		msg.Key,
	)
	fileSize, r, err := s.store.Read(msg.ID, msg.Key)
	if err != nil {
		return err
	}

	if rc, ok := r.(io.ReadCloser); ok {
		fmt.Println("closing readCloser")
		defer rc.Close()
	}

	peer, ok := s.peers[from]
	if !ok {
		return fmt.Errorf("peer %s not in map", from)
	}

	// First send the "IncomingStream" byte to the peer and then
	// send the file size as an int64.
	if err := peer.Send([]byte{p2p.IncomingStream}); err != nil {
		fmt.Printf("error sending incoming stream message %s", err)
		return err
	}
	if err := binary.Write(peer, binary.LittleEndian, fileSize); err != nil {
		fmt.Printf("error sending file size %s", err)
		return err
	}

	n, err := io.Copy(peer, r)
	if err != nil {
		return err
	}

	fmt.Printf("[%s] written (%d) bytes over the network to %s\n", s.Transport.Addr(), n, from)

	return nil
}

func (s *FileServer) handleMessageStoreFile(from string, msg MessageStoreFile) error {
	peer, ok := s.peers[from]
	if !ok {
		return fmt.Errorf("peer (%s) could not be found in the peer list", from)
	}

	n, err := s.store.Write(msg.ID, msg.Key, io.LimitReader(peer, msg.Size))
	if err != nil {
		log.Println("error writing message", err)
		return err
	}

	fmt.Printf("written %d bytes to disk\n", n)

	peer.CloseStream()

	return nil
}

func (s *FileServer) bootstrapNetwork() error {
	for _, addr := range s.BootstrapNodes {
		go func(addr string) {
			log.Printf("[%s] attempting to connect with remote: %s\n", s.Transport.Addr(), addr)
			if err := s.Transport.Dial(addr); err != nil {
				log.Println("dial error: ", err)
			}
		}(addr)
	}

	return nil
}

func (s *FileServer) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	err := s.bootstrapNetwork()
	if err != nil {
		panic(err)
	}

	s.loop()

	return nil
}

func init() {
	gob.Register(MessageStoreFile{})
	gob.Register(MessageGetFile{})
}
