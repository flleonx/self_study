package p2p

import (
	"encoding/gob"
	"io"
)

type (
	Decoder interface {
		Decode(io.Reader, *RPC) error
	}

	GOBDecoder struct{}

	DefaultDecoder struct{}
)

func (d GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

func (d DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	peekBuf := make([]byte, 1)
	if _, err := r.Read(peekBuf); err != nil {
		return err
	}

	// In case of a stream we are not decoding what is being sent over the network.
	// We just set Stream true that way is possible to handle that in our logic.
	stream := peekBuf[0] == IncomingStream
	if stream {
		msg.Stream = true
		return nil
	}

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
