package auth

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/dmitrychurkin/hotelier/server/mailer"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

type sendPasswordResetLinkInput struct {
	Email string `valid:"required,email"`
	Path  string `valid:"required,url"`
}

// SendPasswordResetLink resolver
func SendPasswordResetLink(ctx *context.Context, gc *gin.Context, p *prisma.Client, email, path string) (*bool, error) {
	// 1. validate input
	email, path = govalidator.Trim(email, ""), govalidator.Trim(path, "")
	if isValid, _ := govalidator.ValidateStruct(&sendPasswordResetLinkInput{Email: email, Path: path}); !isValid {
		return nil, userNotFound
	}

	// 2. get user from DB
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	resendTokenTimespan := defaultResendTokenTimespan
	if r, err := strconv.Atoi(os.Getenv("RESEND_TOKEN_TIMESPAN")); err == nil {
		resendTokenTimespan = r
	}

	// 2. Check if elapsed time since last email over resendTokenTimespan
	if user.PasswordResetToken != nil && len(*user.PasswordResetToken) > 0 {
		if t, err := time.Parse(time.RFC3339, *user.PasswordResetTokenCreatedAt); err == nil {
			if el := int(time.Since(t).Minutes()); el < resendTokenTimespan {
				return nil, allowSendEmailAfter(el)
			}
		}
	}

	passwordResetTokenLength := defaultPasswordResetTokenLength
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_RESET_TOKEN_LENGTH")); err == nil {
		passwordResetTokenLength = r
	}

	// 3. Generate password reset token
	passwordResetToken, err := generateURLHash(passwordResetTokenLength)
	if err != nil {
		return nil, err
	}

	// 4. Store into db
	now := time.Now()
	updatedUser, err := p.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			ID: &user.ID,
		},
		Data: prisma.UserUpdateInput{
			PasswordResetToken:          &passwordResetToken,
			PasswordResetTokenCreatedAt: prisma.Str(now.Format(time.RFC3339)),
			PasswordResetTokenExpiresAt: prisma.Str(now.Add(time.Minute * 30).Format(time.RFC3339)),
		},
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	// 5. get request url
	if path[len(path)-1:] != "/" {
		path += "/"
	}

	// 6. construct reset url
	link := path + *updatedUser.PasswordResetToken

	// 7. Send email
	host := gc.Request.Host
	if i := strings.Index(host, ":"); i != -1 {
		host = host[:i]
	}

	from := "no-reply@" + host
	if !govalidator.IsEmail(from) {
		from = ""
	}

	emailAgent := &mailer.ResetPasswordData{Link: link}
	success, err := emailAgent.Send(&mailer.MailRequest{
		From: from,
		To:   []string{user.Email},
	})
	return &success, err
}
