package network

import (
	"fmt"
	"testing"
)

func TestDoPostAsync(t *testing.T) {
	_ = DoGetAsync("http://localhost:8096/health", 15, func(resp string, err *HttpError) {
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(resp)
	})
	select {}
}
