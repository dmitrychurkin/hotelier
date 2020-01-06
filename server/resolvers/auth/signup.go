package auth

import (
	"context"
	"os"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type signupInput struct {
	Email    string `valid:"required,email"`
	Password string `valid:"required,stringlength(8|1000)"`
	// ConfirmPassword string  `valid:"required,stringlength(8|1000)"`
	FirstName *string `valid:"stringlength(0|1000)"`
	LastName  *string `valid:"stringlength(0|1000)"`
}

// Signup resolver
func Signup(ctx *context.Context, gc *gin.Context, p *prisma.Client, email string, firstName *string, lastName *string, password string) (*models.User, error) {
	// 1. validate input
	email, password, fName, lName :=
		govalidator.Trim(email, ""),
		govalidator.Trim(password, ""),
		// govalidator.Trim(confirmPassword, ""),
		govalidator.Trim(*firstName, ""),
		govalidator.Trim(*lastName, "")

	isValid, _ := govalidator.ValidateStruct(&signupInput{
		Email:    email,
		Password: password,
		// ConfirmPassword: confirmPassword,
		FirstName: &fName,
		LastName:  &lName,
	})

	// if strings.Compare(password, confirmPassword) != 0 {
	// 	return nil, passwordsMismatch
	// }

	if !isValid {
		return nil, invalidinputFields
	}

	// 2. Hash password
	roundsHash := defaultPasswordHashRounds
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_ROUNDS")); err == nil {
		roundsHash = r
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), roundsHash)
	if err != nil {
		return nil, err
	}

	// 3. save user into DB
	user, err := p.CreateUser(prisma.UserCreateInput{
		Email:     email,
		Password:  string(hash),
		FirstName: &fName,
		LastName:  &lName,
	}).Exec(*ctx)

	if err != nil {
		return nil, err
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
