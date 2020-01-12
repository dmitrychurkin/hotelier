package auth

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func parseAuthToken(gc *gin.Context) (*jwt.StandardClaims, error) {
	var (
		authTokenSecret     = os.Getenv("AUTH_TOKEN_SECRET")
		authTokenCookieName = os.Getenv("AUTH_TOKEN_COOKIE_NAME")
	)

	if len(authTokenSecret) == 0 {
		authTokenSecret = defaultTokenSecret
	}
	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	// 1. get token from req cookie
	t, err := gc.Cookie(authTokenCookieName)

	if err != nil {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, err
	}

	if len(t) == 0 {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, nil
	}

	// 2. parse token claims
	claims := &jwt.StandardClaims{}
	jwtToken, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(authTokenSecret), nil
	})

	if err != nil {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, err
	}

	if !jwtToken.Valid {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, nil
	}

	return claims, err
}

func issueAuthToken(gc *gin.Context, userID string) error {
	var (
		authTokenSecret     = os.Getenv("AUTH_TOKEN_SECRET")
		authTokenLifetime   = os.Getenv("AUTH_TOKEN_LIFETIME")
		authTokenCookieName = os.Getenv("AUTH_TOKEN_COOKIE_NAME")
		authTokenMin        = defaultTokenLifetime
	)

	if len(authTokenSecret) == 0 {
		authTokenSecret = defaultTokenSecret
	}

	d, err := strconv.Atoi(authTokenLifetime)
	if err == nil {
		authTokenMin = d
	}

	authTokenClaims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(time.Duration(authTokenMin) * time.Minute).Unix(),
	}

	authToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authTokenClaims).SignedString([]byte(authTokenSecret))

	if err != nil {
		return err
	}

	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	gc.SetCookie(authTokenCookieName, authToken, authTokenMin*60, "", "", false, true)

	return nil
}

func generateURLHash(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes), err
}
