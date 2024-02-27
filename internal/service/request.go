package service

import (
	"github.com/guregu/null/v5"
	"time"
)

type Request struct {
	Level     string      `json:"level"`
	Prefix    string      `json:"prefix"`
	Message   string      `json:"message"`
	Lang      null.String `json:"lang"`
	UserAgent null.String `json:"user_agent"`
	Trace     null.String `json:"trace"`
	IP        null.String `json:"ip"`
	Time      time.Time   `json:"time"`
}
