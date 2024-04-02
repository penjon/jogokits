package utils

import (
	"bytes"
	"sync"
)

type ByteBuffPool struct {
	pools *sync.Pool
}

var pool *ByteBuffPool

func GetBuffPool() *ByteBuffPool {
	if nil == pool {
		pool = &ByteBuffPool{}
		pool.pools = &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 4096))
			},
		}
	}
	return pool
}

func (p *ByteBuffPool) Get() *bytes.Buffer {
	buff := p.pools.Get().(*bytes.Buffer)
	buff.Reset()
	return buff
}
func (p *ByteBuffPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.pools.Put(b)
}
