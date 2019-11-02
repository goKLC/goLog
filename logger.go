package goLog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const INFO Priority = "Info"
const ERROR Priority = "Error"
const WARNING Priority = "Warning"
const SUCCESS Priority = "Success"

type Data map[string]interface{}
type Priority string

type Log struct {
	date     string
	priority Priority
	message  string
	data     Data
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

func (l *Log) Info(message string, data Data) {
	log(l, message, data, INFO)
}

func (l *Log) Error(message string, data Data) {
	log(l, message, data, ERROR)
}

func (l *Log) Warning(message string, data Data) {
	log(l, message, data, WARNING)
}

func (l *Log) Success(message string, data Data) {
	log(l, message, data, SUCCESS)
}

func log(l *Log, message string, data Data, priority Priority) {
	l.date = time.Now().Format(config.TimeFormat)
	l.message = message
	l.data = data
	l.priority = priority

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
	_, err = file.WriteString(fmt.Sprintf("%v: ", log.priority))
	_, err = file.WriteString(fmt.Sprint(log.message))
	_, err = file.WriteString(fmt.Sprintln(""))

	if config.PrintTerminal {
		fmt.Printf("[%v] %v: %v", log.date, log.priority, log.message)
		fmt.Println("")
	}

	if len(log.data) != 0 {
		data, _ := json.MarshalIndent(log.data, "	", "  ")

		_, err = file.WriteString("	" + string(data))
		_, err = file.WriteString(fmt.Sprintln(""))

		if config.PrintTerminal {
			fmt.Print("	" + string(data))
			fmt.Println("")
		}
	}

	if err != nil {
		fmt.Println(err)
	}

	_ = file.Close()
}
