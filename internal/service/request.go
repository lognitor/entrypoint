package service

import "time"

type Request struct {
	Level   string    `json:"level"`
	Prefix  string    `json:"prefix"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
