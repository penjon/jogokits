package network

import (
	"fmt"
	"git.ingcreations.com/ingcreations-golang/gokits/codec"
	"git.ingcreations.com/ingcreations-golang/gokits/utils"
	"testing"
)

func TestStartServe(t *testing.T) {
	StartServe(8000, onConnected)
}

var index uint = 1
var clients = make(map[string]*TcpClient)

func onConnected(client *TcpClient) {
	client.SetHandlers(onReceive, onErr)
	client.Decoder = &codec.LengthCodec{
		//Options: &codec.Options{
		//	DefaultBuffSize: 2048,
		//	HeadSize:        4,
		//},
	}
	client.Encoder = &codec.LengthCodec{}
	id := fmt.Sprintf("%d", utils.GetTimeMillis())
	client.ID = id
	clients[id] = client
	index++
}

func onReceive(c *TcpClient, data []byte) {
	fmt.Println(c.ID)
	fmt.Println(string(data))
}

func onErr(c *TcpClient, err error) {
	fmt.Println(err.Error())
	fmt.Println("关闭[" + c.ID + "]")
	//c := clients[c.ID]
	c.Close()
	delete(clients, c.ID)
}
