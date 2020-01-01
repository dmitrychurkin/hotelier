package mailer

import (
	"bytes"
	"gopkg.in/mailgun/mailgun-go.v1"
	"html/template"
	"os"
)

// MailRequest ...
type MailRequest struct {
	From             string
	Title            string
	Subject          string
	HTMLMessage      string
	PlainTextMessage string
	To               []string
}

// SendMail ...
func (mailRequest *MailRequest) SendMail() (bool, error) {
	mg := mailgun.NewMailgun(
		os.Getenv("MAILGUN_DOMAIN"),
		os.Getenv("MAILGUN_KEY"),
		os.Getenv("MAILGUN_PUB_KEY"),
	)
	message := mg.NewMessage(
		mailRequest.From,
		mailRequest.Title,
		mailRequest.PlainTextMessage,
		mailRequest.To...,
	)
	message.SetHtml(mailRequest.HTMLMessage)

	if _, _, err := mg.Send(message); err != nil {
		return false, err
	}
	return true, nil
}

// ParseTemplate ...
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
