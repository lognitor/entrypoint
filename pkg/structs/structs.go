package structs

import (
	"time"
)

type Log struct {
	Level   string        `json:"level" db:"level"`
	Prefix  string        `json:"prefix" db:"prefix"`
	IP      string        `json:"ip" db:"ip"`
	Agent   string        `json:"agent" db:"agent"`
	Message any           `json:"message" db:"message"`
	Trace   []Frame       `json:"trace" db:"trace"`
	Source  FrameWithCode `json:"source" db:"source"`
	Time    time.Time     `json:"time" db:"created_at"`
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
