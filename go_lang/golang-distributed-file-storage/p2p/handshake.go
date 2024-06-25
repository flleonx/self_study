package p2p

type (
	// HandshakeFunc ?
	HandhsakeFunc func(Peer) error
)

func NOPHandshakeFunc(Peer) error {
	return nil
}
