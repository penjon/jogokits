package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	//"github.com/valyala/gozstd"
	"io"
)

/*func CompressByZstd(data []byte) []byte {
	return gozstd.Compress(nil, data)
}

func DecompressByZstd(data []byte) ([]byte, error) {
	return gozstd.Decompress(nil, data)
}*/

func CompressByGzip(data []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := gzip.NewWriter(&buff)
	defer func() {
		_ = writer.Close()
	}()
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Flush(); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func DecompressByGzip(data []byte) ([]byte, error) {
	reader, e := gzip.NewReader(bytes.NewReader(data))
	if e != nil {
		return nil, fmt.Errorf("read erorr[%v]", e.Error())
	}
	defer func() {
		_ = reader.Close()
	}()
	var content []byte
	var tmp = make([]byte, 128)

	for {
		n, err := reader.Read(tmp)
		if err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			break
		}
		content = append(content, tmp[:n]...)
	}

	return content, nil
}
