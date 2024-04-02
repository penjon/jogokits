package network

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpError struct {
	StatusCode int
	Err        error
}

func (i *HttpError) Error() string {
	return fmt.Sprintf("StatusCode[%d],Error message[%s]", i.StatusCode, i.Err.Error())
}

func Error(code int, err string) *HttpError {
	return &HttpError{
		StatusCode: code,
		Err:        errors.New(err),
	}
}

type RequestCallback func(string, *HttpError)

func DoGetAsync(url string, timeout uint, callback RequestCallback) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	go doRequest(request, timeout, callback)
	return nil
}

func doRequest(request *http.Request, timeout uint, callback RequestCallback) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	response, err := http.DefaultClient.Do(request.WithContext(ctx))
	defer func() {
		if nil != cancel {
			cancel()
		}
	}()

	if err != nil {
		callback("", Error(http.StatusGatewayTimeout, err.Error()))
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		callback("", Error(0, err.Error()))
		return
	}

	if response.StatusCode != http.StatusOK {
		callback("", Error(response.StatusCode, string(data)))
		return
	}

	callback(string(data), nil)

	defer func() {
		if nil != response {
			_ = response.Body.Close()
		}
	}()
}
