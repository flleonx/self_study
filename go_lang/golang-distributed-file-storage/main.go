package main

import (
	"bytes"
	"fmt"
	"fss/p2p"
	"io"
	"log"
	"time"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(),
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":50100")
	s2 := makeServer(":50200", ":50100")
	s3 := makeServer(":50300", ":50200", ":50100")

	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(500 * time.Millisecond)

	go func() {
		log.Fatal(s2.Start())
	}()
	time.Sleep(500 * time.Millisecond)

	go func() {
		err := s3.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("picture_%d.png", i)
		data := bytes.NewReader([]byte("my big data file here!"))
		if err := s3.Store(key, data); err != nil {
			log.Fatal(err)
		}

		if err := s3.store.Delete(s3.ID, key); err != nil {
			log.Fatal(err)
		}

		r, err := s3.Get(key)
		if err != nil {
			log.Fatal(err)
		}

		b, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("result: ", string(b))
	}
}
