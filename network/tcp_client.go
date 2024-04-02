package network

import (
	"git.ingcreations.com/ingcreations-golang/gokits/codec"
	"net"
)

//const (
//	BUFF_SIZE = 512
//)

type OnReceiver func(*TcpClient, []byte)
type OnError func(*TcpClient, error)

type TcpClient struct {
	conn     net.Conn
	receiver OnReceiver
	err      OnError
	ID       string
	Decoder  codec.Decoder
	Encoder  codec.Encoder
}

func (c *TcpClient) SetHandlers(receiver OnReceiver, err OnError) {
	c.receiver = receiver
	c.err = err
}

func (c *TcpClient) Write(data []byte) error {
	if nil != c.Encoder {
		//经过编码器处理
		buff, err := c.Encoder.Encode(data)
		if err != nil {
			return err
		}
		_, err = c.conn.Write(buff)
		return err
	}
	_, err := c.conn.Write(data)
	return err
}

func (c *TcpClient) Close() {
	_ = c.conn.Close()
}

func (c *TcpClient) start() {
	go func() {
		//buf := make([]byte,BUFF_SIZE)
		//reader := bytes.NewBuffer(make([]byte,BUFF_SIZE))
		//reader.Reset()
		//reader := bufio.NewReader(c.conn)
		for {
			data, _, err := c.Decoder.Decode(c.conn)
			if err != nil {
				if c.err != nil {
					c.err(c, err)
				}
				return
			}
			if c.receiver != nil {
				c.receiver(c, data)
			}
		}
	}()
}

func NewClient(conn net.Conn) *TcpClient {
	coder := &codec.DefaultCodec{}
	client := &TcpClient{conn: conn}
	client.Encoder = coder
	client.Decoder = coder
	return client
}
