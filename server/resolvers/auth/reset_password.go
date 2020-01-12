package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type resetPasswordInput struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required,stringlength(8|1000)"`
	// ConfirmPassword    string `valid:"required,stringlength(8|1000)"`
	PasswordResetToken string `valid:"required"`
}

// ResetPassword resolver
func ResetPassword(ctx *context.Context, gc *gin.Context, p *prisma.Client, email, password, passwordResetToken string) (*models.User, error) {
	// 1. validate input
	email, passwordResetToken, password =
		govalidator.Trim(email, ""),
		govalidator.Trim(passwordResetToken, ""),
		govalidator.Trim(password, "")
		// govalidator.Trim(confirmPassword, "")

	isValid, _ := govalidator.ValidateStruct(&resetPasswordInput{
		Email:    email,
		Password: password,
		// ConfirmPassword:    confirmPassword,
		PasswordResetToken: passwordResetToken,
	})

	// if strings.Compare(password, confirmPassword) != 0 {
	// 	return nil, passwordsMismatch
	// }

	if !isValid {
		return nil, invalidinputFields
	}

	// 2. check if token expired
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	if user.PasswordResetToken == nil || len(*user.PasswordResetToken) == 0 {
		return nil, invalidPasswordResetToken
	}

	expiresAt, err := time.Parse(time.RFC3339, *user.PasswordResetTokenExpiresAt)

	if err != nil {
		return nil, err
	}

	if expiresAt.Before(time.Now()) {
		return nil, invalidPasswordResetToken
	}

	// 3. Hash password
	roundsHash := defaultPasswordHashRounds
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_ROUNDS")); err == nil {
		roundsHash = r
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), roundsHash)
	if err != nil {
		return nil, err
	}

	// 4. update user record
	user, err = p.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Email: &email,
		},
		Data: prisma.UserUpdateInput{
			Password:           prisma.Str(string(hash)),
			PasswordResetToken: prisma.Str(""),
		},
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	// 5. issue auth token
	// 6. set cookie
	if err := issueAuthToken(gc, user.ID); err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      models.UserRoles(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
