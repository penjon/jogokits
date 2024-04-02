package utils

import (
	"reflect"
	"sync"
)

type Service interface {
	GetServiceName() string
}

var services = make(map[string]interface{})
var lock sync.Mutex

func GetService[T Service]() T {
	var r T
	service, ok := services[r.GetServiceName()]
	if !ok {
		lock.Lock()
		defer lock.Unlock()
		t := reflect.TypeOf(r)
		ptr := reflect.New(t.Elem())
		r = ptr.Interface().(T)
		service = r
		services[r.GetServiceName()] = service
	}
	return service.(T)
}
