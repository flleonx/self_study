package p2p

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

type (
	// TCPPeer represents the remote node over a TCP established connection.
	TCPPeer struct {
		// The underlying connection of the peer. Which in this case
		// is a TCP connection.
		net.Conn

		// if we dial and retrieve a conn -> outbound == true
		// if we accept and retrieve a conn -> outbound == false
		outbound bool

		wg *sync.WaitGroup
	}

	TCPTransportOps struct {
		ListenAddr    string
		HandshakeFunc HandhsakeFunc
		Decoder       Decoder
		OnPeer        func(Peer) error
	}

	TCPTransport struct {
		TCPTransportOps
		listener net.Listener
		rpcch    chan RPC
	}
)

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		outbound: outbound,
		wg:       &sync.WaitGroup{},
	}
}

func (p *TCPPeer) CloseStream() {
	p.wg.Done()
}

// Send implements the Peer interface.
func (p *TCPPeer) Send(b []byte) error {
	_, err := p.Conn.Write(b)
	return err
}

func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
		rpcch:           make(chan RPC, 1024),
	}
}

// Adrr implements the Transport interface.
func (t *TCPTransport) Addr() string {
	return t.ListenAddr
}

// Dial implements the Transport interface.
func (t *TCPTransport) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	go t.handleConn(conn, true)

	return nil
}

// Close implements the transport interface.
func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

// Consumes implements the Transport interface, which will return read-only channel
// for reading the incoming messages received from another peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", t.listener.Addr())

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}

		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn, false)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn, outbound bool) {
	var err error
	defer func() {
		fmt.Printf("TCP dropping peer connection: %s\n", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, outbound)

	if err := t.HandshakeFunc(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	for {
		rpc := RPC{}
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			fmt.Printf("TCP decode error: %s\n", err)
			return
		}

		rpc.From = conn.RemoteAddr().String()
		if rpc.Stream {
			peer.wg.Add(1)
			fmt.Printf("[%s] incoming stream, waiting...\n", conn.RemoteAddr())
			peer.wg.Wait()
			fmt.Printf("[%s] stream closed, resuming read loop...\n", conn.RemoteAddr())
			continue
		}

		t.rpcch <- rpc
	}
}
