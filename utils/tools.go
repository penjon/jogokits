package utils

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

//获取当前运行应用的名称
func GetExecAppName() string {
	path,_ := exec.LookPath(os.Args[0])
	file := filepath.Base(path)
	return file
}

func GetLogNameByAppName() string {
	appName := GetExecAppName()
	return fmt.Sprintf("%s_%d_%d.log",appName,GetTimeMillis(),rand.Int63n(100))
}

//获取当前第一个获取的局域网IP
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.IsGlobalUnicast() {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}