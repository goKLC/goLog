package goLog

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailHandler struct {
	From   string
	To     []string
	Auth   smtp.Auth
	Host   string
	Levels []Level
}

func NewEmailHandler() *EmailHandler {

	return &EmailHandler{}
}

func (eh *EmailHandler) Write(log Log) {
	find := false

	for _, level := range eh.Levels {
		if log.level == level {
			find = true
			break
		}
	}

	if !find {
		return
	}

	message := "To: %s \r\n" +
		"From: %s \r\n" +
		"Subject: We have a problem\r\n" +
		"\r\n" +
		"<!DOCTYPE html>" +
		"<html lang=\"en\">" +
		"<head>" +
		"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>" +
		"</head>" +
		"<body>" +
		"<table>" +
		"<tr>" +
		"<td>Date</td>" +
		"<td>%s</td>" +
		"</tr>" +
		"<tr>" +
		"<td>Level</td>" +
		"<td>%s</td>" +
		"</tr>" +
		"<tr>" +
		"<td>Message</td>" +
		"<td>%s</td>" +
		"</tr>" +
		"<tr>" +
		"<td>Data</td>" +
		"<td>%s</td>" +
		"</tr>" +
		"</table>" +
		"</body>" +
		"</html>\r\n"

	var context []byte
	if len(log.context) != 0 {
		context, _ = json.Marshal(log.context)
	}

	message = fmt.Sprintf(message, strings.Join(eh.To, ", "), eh.From, log.date, log.level, log.message, context)

	err := smtp.SendMail(eh.Host, eh.Auth, eh.From, eh.To, []byte(message))

	if err != nil {
		fmt.Println(err.Error())
	}
}
