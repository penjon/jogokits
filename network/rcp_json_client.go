package network

import (
	"fmt"
	"git.ingcreations.com/ingcreations-golang/gokits/logs"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type RpcOption struct {
	//最大重连次数
	ReconnectMaxTime int
	//重连间隔时间
	ReconnectWaitTime int
}
type RpcJsonClient struct {
	addr   string
	Option *RpcOption
	client *rpc.Client
}

func (i *RpcJsonClient) Connect(addr string) error {
	i.addr = addr
	tryCount := 0
	maxTry := 3
	tryWait := 5
	if nil != i.Option {
		if i.Option.ReconnectMaxTime > 0 {
			maxTry = i.Option.ReconnectMaxTime
		}
		if i.Option.ReconnectWaitTime > 0 {
			tryWait = i.Option.ReconnectWaitTime
		}
	}
	for {
		if tryCount >= maxTry {
			return net.ErrClosed
		}
		tryCount++
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			i.client = jsonrpc.NewClient(conn)
			return nil
		}
		ticker := time.NewTicker(time.Duration(tryWait) * time.Second)
		//重连失败,直接返回错误
		logs.Get().Info("开始进入第[%v]次重连等待,等待时间[%v]秒", tryCount, tryWait)
		for {
			<-ticker.C
			ticker.Stop()
			logs.Get().Info("第[%v]次重连等待结束", tryCount)
			break
		}
	}
	return nil
}
func (i *RpcJsonClient) Call(serviceName, method string, requestParam any, responseParam any) error {
	service := fmt.Sprintf("%s.%s", serviceName, method)
	for {
		if err := i.client.Call(service, requestParam, responseParam); err != rpc.ErrShutdown {
			return err
		}
		if err := i.Connect(i.addr); err != nil {
			return err
		}
	}
}

func Call[T any](client *RpcJsonClient, serviceName, method string, requestParam any) (*T, error) {
	var r T
	if err := client.Call(serviceName, method, requestParam, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
