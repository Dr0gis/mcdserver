package utils

import "fmt"

type Logs struct {
	enable bool
}

func NewLogs(enable bool) Logs {
	return Logs{enable: enable}
}

func (logs Logs) Print(text ...interface{}) {
	if logs.enable {
		fmt.Println(text)
	}
}