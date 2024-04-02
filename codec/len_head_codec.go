package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Options struct {
	DefaultBuffSize int
	HeadSize        int
}

type LengthCodec struct {
	Options *Options
}

// 返回参数说明
// []byte 	读取的数据
// bool   	是否还有剩余数据
// error		异常
func (c *LengthCodec) Decode(conn net.Conn) ([]byte, bool, error) {
	headSize := 4
	if c.Options != nil {
		headSize = c.Options.HeadSize
	}
	buffSize := 1024
	if c.Options != nil {
		buffSize = c.Options.DefaultBuffSize
	}
	lenBytes := make([]byte, headSize)
	if _, err := conn.Read(lenBytes); err != nil {
		return nil, false, err
	}

	buff := bytes.NewBuffer(lenBytes)
	var size int32
	if err := binary.Read(buff, binary.BigEndian, &size); err != nil {
		return nil, false, err
	}
	bodySize := int(size)
	fmt.Printf("数据包长度[%d]", bodySize)

	buff.Reset()
	readAll := 0
	length := buffSize
	for {
		length = buffSize
		//根据当前已经读取的数据量确定这次读取的大小
		if readAll+length > int(bodySize) {
			length = int(bodySize) - readAll
		}
		buffBytes := make([]byte, length)
		size, err := conn.Read(buffBytes)

		buff.Write(buffBytes[0:size])
		readAll += size
		if err == io.EOF {
			continue
		}
		fmt.Printf("等待数据 %d/%d\n", readAll, bodySize)
		//已经读取到所有数据
		if readAll >= bodySize {
			break
		}
	}
	return buff.Bytes(), false, nil
}

func (c *LengthCodec) Encode(source []byte) ([]byte, error) {
	totalLen := int32(len(source))
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, totalLen); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, source); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil

}
