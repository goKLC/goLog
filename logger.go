package goLog

import (
	"encoding/json"
	"fmt"
	"os"
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

type Context map[string]interface{}
type Level string

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
	Path          string
	FileName      string
	TimeFormat    string
	PrintTerminal bool
}

var config *Config

func New() (*Log, *Config) {
	config = &Config{
		Path:          "",
		FileName:      "goLog.log",
		TimeFormat:    "2006-01-02 15:04:05",
		PrintTerminal: false,
	}

	return &Log{}, config
}

func (l *Log) Debug(message string, context Context) {
	log(l, message, context, DEBUG)
}

func (l *Log) Info(message string, context Context) {
	log(l, message, context, INFO)
}

func (l *Log) Notice(message string, context Context) {
	log(l, message, context, NOTICE)
}

func (l *Log) Error(message string, context Context) {
	log(l, message, context, ERROR)
}

func (l *Log) Warning(message string, context Context) {
	log(l, message, context, WARNING)
}

func (l *Log) Critical(message string, context Context) {
	log(l, message, context, CRITICAL)
}

func (l *Log) Alert(message string, context Context) {
	log(l, message, context, ALERT)
}

func (l *Log) Emergency(message string, context Context) {
	log(l, message, context, EMERGENCY)
}

func log(l *Log, message string, context Context, level Level) {
	l.date = time.Now().Format(config.TimeFormat)
	l.message = message
	l.context = context
	l.level = level

	write(l)
}

func write(log *Log) {
	var path, _ = os.Getwd()
	var folderPath = fmt.Sprintf("%s/%s", path, config.Path)
	var filePath = fmt.Sprintf("%s/%s", folderPath, config.FileName)
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
}
