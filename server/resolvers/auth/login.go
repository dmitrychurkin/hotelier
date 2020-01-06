package auth

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginInput struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required,stringlength(8|1000)"`
}

// Login resolver
func Login(ctx *context.Context, gc *gin.Context, p *prisma.Client, email string, password string) (*models.User, error) {
	// 1. validate input
	email, password = govalidator.Trim(email, ""), govalidator.Trim(password, "")
	if isValid, _ := govalidator.ValidateStruct(&loginInput{Email: email, Password: password}); !isValid {
		return nil, userNotFound
	}

	// 2. get user from DB
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	// 3. check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, userNotFound
	}

	// 4. issue auth token
	// 5. set cookie
	if err := issueAuthToken(ctx, gc, user); err != nil {
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
