package mailer

import (
	"os"
	"path/filepath"
)

const (
	from  = "no-reply@hotel.com"
	title = "Reset Password Email"
)

// ResetPasswordData ...
type ResetPasswordData struct {
	Link string
}

// Send ...
func (resetPasswordData *ResetPasswordData) Send(mailRequest *MailRequest) (bool, error) {
	cwd, _ := os.Getwd()
	htmlContent, err := ParseTemplate(filepath.Join(cwd, "mailer", "templates", "./resetPassword.html"), resetPasswordData)
	if err != nil {
		return false, err
	}

	textContent, err := ParseTemplate(filepath.Join(cwd, "mailer", "templates", "./resetPassword.tmpl"), resetPasswordData)
	if err != nil {
		return false, err
	}

	if len(mailRequest.From) == 0 {
		mailRequest.From = from
	}
	if len(mailRequest.Title) == 0 {
		mailRequest.Title = title
	}

	mailRequest.HTMLMessage = htmlContent
	mailRequest.PlainTextMessage = textContent

	return mailRequest.SendMail()
}
