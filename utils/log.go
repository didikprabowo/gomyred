package utils

import (
	"fmt"
	"log"
	"os"
)

type (
	Logger struct {
		Info  string
		Error error
	}
)

// NewLogError
func (l Logger) NewLogError() {
	logFile, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("%s => %s", l.Info, l.Error))

}
