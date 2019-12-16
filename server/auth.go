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

	// 4. issue access and refresh token
	// 5. set headers
	if err := issueTokens(ctx, user); err != nil {
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

	if len(passwordHashRounds) != 0 {
		r, err := strconv.Atoi(passwordHashRounds)
		if err == nil {
			roundsHash = r
		}
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

	// 4. issue access and refresh token
	// 5. set headers
	if err := issueTokens(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func issueTokens(ctx context.Context, user *prisma.User) error {
	var (
		accessTokenKey      = os.Getenv("ACCESS_TOKEN_SECRET")
		accessTokenLifetime = os.Getenv("ACCESS_TOKEN_LIFETIME")
		accessTokenMin      = 5
	)

	if len(accessTokenKey) == 0 {
		accessTokenKey = "jwt_access_secret|123"
	}

	if len(accessTokenLifetime) != 0 {
		d, err := strconv.Atoi(accessTokenLifetime)
		if err == nil {
			accessTokenMin = d
		}
	}

	accessTokenClaims := &jwt.StandardClaims{
		Subject:   user.ID,
		ExpiresAt: time.Now().Add(time.Duration(accessTokenMin) * time.Minute).Unix(),
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(accessTokenKey))

	if err != nil {
		return err
	}

	var (
		refreshTokenKey      = os.Getenv("REFRESH_TOKEN_SECRET")
		refreshTokenLifetime = os.Getenv("REFRESH_TOKEN_LIFETIME")
		refreshTokenMin      = 300
	)

	if len(refreshTokenKey) == 0 {
		refreshTokenKey = "jwt_refresh_secret|987"
	}

	if len(refreshTokenLifetime) != 0 {
		d, err := strconv.Atoi(refreshTokenLifetime)
		if err == nil {
			refreshTokenMin = d
		}
	}

	refreshTokenClaims := &jwt.StandardClaims{
		Subject:   user.ID,
		ExpiresAt: time.Now().Add(time.Duration(refreshTokenMin) * time.Minute).Unix(),
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(refreshTokenKey))

	if err != nil {
		return err
	}

	ctxData := ctx.Value(dataCtxKey).(*contextData)
	ctxData.Res.Header().Set("Access-Control-Expose-Headers", "x-access-token,x-refresh-token")
	ctxData.Res.Header().Set("x-access-token", accessToken)
	ctxData.Res.Header().Set("x-refresh-token", refreshToken)

	return nil
}
