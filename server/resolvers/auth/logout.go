package auth

import (
	"context"
	"github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"os"
)

// Logout resolver
func Logout(ctx *context.Context, gc *gin.Context) (*bool, error) {
	authTokenCookieName := os.Getenv("AUTH_TOKEN_COOKIE_NAME")
	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)

	return prisma.Bool(true), nil
}
