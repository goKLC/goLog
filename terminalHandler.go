package goLog

import (
	"encoding/json"
	"fmt"
)

type TerminalHandler struct {
}

func NewTerminalHandler() *TerminalHandler {
	return &TerminalHandler{}
}

func (th *TerminalHandler) Write(log Log) {
	var err error
	var context []byte

	fmt.Printf("[%v] %v: %v", log.date, log.level, log.message)
	fmt.Println("")

	if len(log.context) != 0 {
		context, err = json.MarshalIndent(log.context, "	", "  ")
		fmt.Print("	" + string(context))
		fmt.Println("")
	}

	if err != nil {
		fmt.Println(err)
	}
}
