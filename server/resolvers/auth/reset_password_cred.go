package auth

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
)

// ResetPasswordCred resolver
func ResetPasswordCred(ctx *context.Context, p *prisma.Client, passwordResetToken string) (string, error) {
	// 1. check if token exists
	user, err := p.User(prisma.UserWhereUniqueInput{
		PasswordResetToken: prisma.Str(govalidator.Trim(passwordResetToken, "")),
	}).Exec(*ctx)

	if err != nil {
		return "", err
	}

	expiresAt, err := time.Parse(time.RFC3339, *user.PasswordResetTokenExpiresAt)

	if err != nil {
		return "", err
	}

	if expiresAt.Before(time.Now()) {
		return "", invalidPasswordResetToken
	}

	return user.Email, err
}
