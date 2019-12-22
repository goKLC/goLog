package goLog

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
)

type log struct {
	gorm.Model
	Type string `gorm:"type:varchar(100);"`
	Log  string `gorm:"type:text;"`
	Data string `gorm:"type:text;"`
}

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler(DB *gorm.DB) *DBHandler {
	if !DB.HasTable("logs") {
		DB.AutoMigrate(&log{})
	}

	return &DBHandler{DB: DB}
}

func (dbh DBHandler) Write(l Log) {
	var context []byte
	var err error

	if len(l.context) != 0 {
		context, err = json.MarshalIndent(l.context, "	", "  ")
	}

	logModel := &log{
		Type: string(l.level),
		Log:  l.message,
		Data: string(context),
	}

	if err != nil {
		fmt.Println(err)
	}

	dbh.DB.Create(logModel)
}
