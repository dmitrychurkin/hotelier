package main

import (
	"context"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dmitrychurkin/hotelier/server/prisma-generated/prisma-client"
	"golang.org/x/crypto/bcrypt"
)

const reEmail = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// LoginHandler resolver
func LoginHandler(ctx context.Context, p *prisma.Client, email string, password string) (*prisma.User, error) {
	// 1. validate input
	email, password = strings.TrimSpace(email), strings.TrimSpace(password)
	if len(email) == 0 ||
		len(email) > 1000 ||
		len(password) == 0 ||
		len(password) > 1000 {
		return nil, errors.New("We don't have records with associated credentials")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("Invalid email pattern")
	}

	// 2. get user from DB
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("We don't have records with associated credentials")
	}

	// 3. check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("We don't have records with associated credentials")
	}

	// 4. issue auth token
	// 5. set headers
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// SignupHandler resolver
func SignupHandler(ctx context.Context, p *prisma.Client, email string, firstName *string, lastName *string, password string, confirmPassword string) (*prisma.User, error) {
	// 1. validate input
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	confirmPassword = strings.TrimSpace(confirmPassword)
	fName := *firstName
	fName = strings.TrimSpace(fName)
	lName := *lastName
	lName = strings.TrimSpace(lName)
	if len(email) == 0 ||
		len(email) > 1000 ||
		len(password) == 0 ||
		password != confirmPassword {
		return nil, errors.New("Some of input was invalid")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("Invalid email pattern")
	}

	// 2. Hash password
	var (
		passwordHashRounds = os.Getenv("PASSWORD_HASH_ROUNDS")
		roundsHash         = 14
	)

	r, err := strconv.Atoi(passwordHashRounds)
	if err == nil {
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
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// 4. issue auth token
	// 5. set headers
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func issueAuthToken(ctx context.Context, user *prisma.User) error {
	var (
		authTokenSecret     = os.Getenv("AUTH_TOKEN_SECRET")
		authTokenLifetime   = os.Getenv("AUTH_TOKEN_LIFETIME")
		authTokenHeaderName = os.Getenv("AUTH_TOKEN_HEADER_NAME")
		authTokenMin        = 5
	)

	if len(authTokenSecret) == 0 {
		authTokenSecret = "jwt_access_secret|123"
	}

	d, err := strconv.Atoi(authTokenLifetime)
	if err == nil {
		authTokenMin = d
	}

	authTokenClaims := &jwt.StandardClaims{
		Subject:   user.ID,
		ExpiresAt: time.Now().Add(time.Duration(authTokenMin) * time.Minute).Unix(),
	}

	authToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authTokenClaims).SignedString([]byte(authTokenSecret))

	if err != nil {
		return err
	}

	if len(authTokenHeaderName) == 0 {
		authTokenHeaderName = "x-auth-token"
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	res := gc.Writer

	res.Header().Set("Access-Control-Expose-Headers", authTokenHeaderName)
	res.Header().Set(authTokenHeaderName, authToken)

	return nil
}
