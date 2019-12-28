package goLog

import (
	"sync"
	"time"
)

const DEBUG Level = "Debug"
const INFO Level = "Info"
const NOTICE Level = "Notice"
const WARNING Level = "Warning"
const ERROR Level = "Error"
const CRITICAL Level = "Critical"
const ALERT Level = "Alert"
const EMERGENCY Level = "Emergency"

type Level string
type PartitionRange string

type Log struct {
	date    string
	level   Level
	message string
	context map[string]interface{}
}

type LogInterface interface {
	Debug(message string, context map[string]interface{})
	Info(message string, context map[string]interface{})
	Notice(message string, context map[string]interface{})
	Warning(message string, context map[string]interface{})
	Error(message string, context map[string]interface{})
	Critical(message string, context map[string]interface{})
	Alert(message string, context map[string]interface{})
	Emergency(message string, context map[string]interface{})
}

type HandlerInterface interface {
	Write(log Log)
}

type Config struct {
	TimeFormat    string
	PrintTerminal bool
	Handler       []HandlerInterface
}

var config *Config

var mux = &sync.RWMutex{}

func New() (*Log, *Config) {
	config = &Config{
		TimeFormat: "2006-01-02 15:04:05",
		Handler:    make([]HandlerInterface, 0),
	}

	return &Log{}, config
}

func (c *Config) AddHandler(handler HandlerInterface) {
	c.Handler = append(c.Handler, handler)
}

func (l *Log) Debug(message string, context map[string]interface{}) {
	addLog(l, message, context, DEBUG)
}

func (l *Log) Info(message string, context map[string]interface{}) {
	addLog(l, message, context, INFO)
}

func (l *Log) Notice(message string, context map[string]interface{}) {
	addLog(l, message, context, NOTICE)
}

func (l *Log) Error(message string, context map[string]interface{}) {
	addLog(l, message, context, ERROR)
}

func (l *Log) Warning(message string, context map[string]interface{}) {
	addLog(l, message, context, WARNING)
}

func (l *Log) Critical(message string, context map[string]interface{}) {
	addLog(l, message, context, CRITICAL)
}

func (l *Log) Alert(message string, context map[string]interface{}) {
	addLog(l, message, context, ALERT)
}

func (l *Log) Emergency(message string, context map[string]interface{}) {
	addLog(l, message, context, EMERGENCY)
}

func addLog(l *Log, message string, context map[string]interface{}, level Level) {
	mux.Lock()
	log := Log{
		date:    time.Now().Format(config.TimeFormat),
		level:   level,
		message: message,
		context: context,
	}
	mux.Unlock()

	var wg sync.WaitGroup

	wg.Add(len(config.Handler))

	for _, handler := range config.Handler {
		go func(log Log) {
			wg.Done()
			handler.Write(log)
		}(log)
	}
}
