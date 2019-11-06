package goLog

import (
	"encoding/json"
	"fmt"
	"os"
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

const HOURLY PartitionRange = "Hourly"
const DAILY PartitionRange = "Daily"
const MONTHLY PartitionRange = "Monthly"
const YEARLY PartitionRange = "Yearly"

type Context map[string]interface{}
type Level string
type PartitionRange string

type Log struct {
	date    string
	level   Level
	message string
	context Context
}

type LogInterface interface {
	Debug(message string, context Context)
	Info(message string, context Context)
	Notice(message string, context Context)
	Warning(message string, context Context)
	Error(message string, context Context)
	Critical(message string, context Context)
	Alert(message string, context Context)
	Emergency(message string, context Context)
}

type Config struct {
	Path           string
	FileName       string
	TimeFormat     string
	PrintTerminal  bool
	Partition      bool
	PartitionRange PartitionRange
}

var config *Config

var mux = &sync.RWMutex{}

func New() (*Log, *Config) {
	config = &Config{
		Path:           "",
		FileName:       "goLog.log",
		TimeFormat:     "2006-01-02 15:04:05",
		PrintTerminal:  false,
		Partition:      false,
		PartitionRange: DAILY,
	}

	return &Log{}, config
}

func (l *Log) Debug(message string, context Context) {
	addLog(l, message, context, DEBUG)
}

func (l *Log) Info(message string, context Context) {
	addLog(l, message, context, INFO)
}

func (l *Log) Notice(message string, context Context) {
	addLog(l, message, context, NOTICE)
}

func (l *Log) Error(message string, context Context) {
	addLog(l, message, context, ERROR)
}

func (l *Log) Warning(message string, context Context) {
	addLog(l, message, context, WARNING)
}

func (l *Log) Critical(message string, context Context) {
	addLog(l, message, context, CRITICAL)
}

func (l *Log) Alert(message string, context Context) {
	addLog(l, message, context, ALERT)
}

func (l *Log) Emergency(message string, context Context) {
	addLog(l, message, context, EMERGENCY)
}

func addLog(l *Log, message string, context Context, level Level) {
	mux.Lock()
	log := Log{
		date:    time.Now().Format(config.TimeFormat),
		level:   level,
		message: message,
		context: context,
	}
	mux.Unlock()
	go write(log)
}

func write(log Log) {
	mux.Lock()

	var path, _ = os.Getwd()
	var fileName = config.FileName
	var folderPath = fmt.Sprintf("%s/%s", path, config.Path)

	if config.Partition {
		var partitionFormat string

		switch config.PartitionRange {
		case HOURLY:
			partitionFormat = "2006-01-02-15"
			break
		case DAILY:
			partitionFormat = "2006-01-02"
			break
		case MONTHLY:
			partitionFormat = "2006-01"
			break
		case YEARLY:
			partitionFormat = "2006"
			break
		}

		fileName = fmt.Sprintf("%s-%s", time.Now().Format(partitionFormat), fileName)
	}

	var filePath = fmt.Sprintf("%s/%s", folderPath, fileName)
	var file *os.File
	var err error

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		_ = os.MkdirAll(folderPath, 0777)
	}

	fmt.Println(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
	} else {
		file, _ = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	}

	_, err = file.WriteString(fmt.Sprintf("[%v] ", log.date))
	_, err = file.WriteString(fmt.Sprintf("%v: ", log.level))
	_, err = file.WriteString(fmt.Sprint(log.message))
	_, err = file.WriteString(fmt.Sprintln(""))

	if config.PrintTerminal {
		fmt.Printf("[%v] %v: %v", log.date, log.level, log.message)
		fmt.Println("")
	}

	if len(log.context) != 0 {
		context, _ := json.MarshalIndent(log.context, "	", "  ")

		_, err = file.WriteString("	" + string(context))
		_, err = file.WriteString(fmt.Sprintln(""))

		if config.PrintTerminal {
			fmt.Print("	" + string(context))
			fmt.Println("")
		}
	}

	if err != nil {
		fmt.Println(err)
	}

	_ = file.Close()
	mux.Unlock()
}
