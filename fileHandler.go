package goLog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const HOURLY PartitionRange = "Hourly"
const DAILY PartitionRange = "Daily"
const MONTHLY PartitionRange = "Monthly"
const YEARLY PartitionRange = "Yearly"

type FileHandler struct {
	Partition      bool
	PartitionRange PartitionRange
	FileName       string
	Path           string
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		Partition:      false,
		PartitionRange: DAILY,
		FileName:       "goLog.log",
		Path:           "log",
	}
}

func (fh *FileHandler) Write(log Log) {
	var path, _ = os.Getwd()
	var fileName = fh.FileName
	var folderPath = fmt.Sprintf("%s/%s", path, fh.Path)

	if fh.Partition {
		var partitionFormat string

		switch fh.PartitionRange {
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
	var context []byte

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

	if len(log.context) != 0 {
		context, err = json.MarshalIndent(log.context, "	", "  ")

		_, err = file.WriteString("	" + string(context))
		_, err = file.WriteString(fmt.Sprintln(""))
	}

	err = file.Close()

	if err != nil {
		fmt.Println(err)
	}
}
