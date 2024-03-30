package structs

import (
	"time"
)

type Log struct {
	Level   string        `json:"level"`
	Prefix  string        `json:"prefix"`
	IP      string        `json:"ip"`
	Agent   string        `json:"agent"`
	Message any           `json:"message"`
	Trace   []Frame       `json:"trace"`
	Source  FrameWithCode `json:"source"`
	Time    time.Time     `json:"time"`
}

type Frame struct {
	Path string `json:"path"`
	Line int    `json:"line"`
	Func string `json:"func"`
}

type FrameWithCode struct {
	Frame
	Code []string `json:"code"`
}
