package utils

import "time"

type Value struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	Serial     string      `json:"serial,omitempty"`
	ServerTime int64       `json:"serverTime,omitempty"`
}

type ErrorValue struct {
	Code int    `json:"code"`
	Data string `json:"data,omitempty"`
	Err  error
}

func GetTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeMillsByZone(local *time.Location) int64 {
	return time.Now().In(local).UnixNano() / 1e6
}
