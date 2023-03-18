package commons

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type OpsLogger interface {
	Error(message string)
	Info(message string)
	Debug(message string)
	Warning(message string)
}

type opsLogger struct {
	filename string
}

func NewLogger() OpsLogger {
	return &opsLogger{
		filename: "logs/logger - " + (time.Now()).Format("2006-01-02") + ".log",
	}
}

func (r opsLogger) Logger() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	// var filename string = "logs/logger - " + (time.Now()).Format("2006-01-02") + ".log"
	f, err := os.OpenFile(r.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	log.SetFormatter(Formatter)
	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}

func (r opsLogger) Error(message string) {
	log.Error(message)
}

func (r opsLogger) Info(message string) {
	log.Info(message)
}

func (r opsLogger) Debug(message string) {
	log.Debug(message)
}

func (r opsLogger) Warning(message string) {
	log.Warning(message)
}
