package network

import (
	"fmt"
	"net"
)

type Handler func(client *TcpClient)

var handler Handler

func StartServe(port uint, callback Handler) {
	handler = callback
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go initClient(conn)
	}
}

func initClient(conn net.Conn) {
	client := NewClient(conn)
	handler(client)
	client.start()
}
