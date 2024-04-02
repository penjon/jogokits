package codec

import (
	"bytes"
	"github.com/penjon/jogokits/logs"
	"net"
)

const (
	BUFF_SIZE = 512
)

type DefaultCodec struct {
}

func (c *DefaultCodec) Decode(conn net.Conn) ([]byte, bool, error) {
	buf := make([]byte, BUFF_SIZE)
	reader := new(bytes.Buffer)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			return nil, false, err
		}

		if size > 0 {
			reader.Write(buf[:size])
			if size >= BUFF_SIZE {
				logs.Get().Debug("继续读取..")
				continue
			}
		}

		data := reader.Bytes()
		return data, false, nil
	}
}

func (c *DefaultCodec) Encode(source []byte) ([]byte, error) {
	return source, nil
}
