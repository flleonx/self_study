package main

import (
	"bytes"
	"fmt"
	"fss/p2p"
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
	s1 := makeServer(":3000")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(time.Second * 2)

	go s2.Start()
	time.Sleep(time.Second * 2)

	data := bytes.NewReader([]byte("my big data file here!"))
	s2.Store(fmt.Sprintf("coolpicture.jpg"), data)
	time.Sleep(5 * time.Millisecond)

	// r, err := s2.Get("coolpicture.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// b, err := io.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("result: ", string(b))
}
