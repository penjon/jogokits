package codec

import (
	"net"
)

type Decoder interface {
	Decode(net.Conn) ([]byte, bool, error)
}

type Encoder interface {
	Encode([]byte) ([]byte, error)
}
