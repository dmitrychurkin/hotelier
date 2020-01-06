package auth

import (
	"context"
	"net/http"

	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/gqlerror"
)

// User resolver
func User(ctx *context.Context, gc *gin.Context, p *prisma.Client) (*models.User, error) {
	// 1. get jwt token claims
	claims, err := parseAuthToken(ctx, gc)
	if err != nil {
		return nil, err
	}
	if claims == nil {
		return nil, &gqlerror.Error{
			Message: "Unautorized",
			Extensions: map[string]interface{}{
				"code": http.StatusUnauthorized,
			},
		}
	}

	// 2. query user
	user, err := p.User(prisma.UserWhereUniqueInput{
		ID: &claims.Subject,
	}).Exec(*ctx)

	if err != nil {
		return nil, err
	}

	// 3. issue auth token
	// 4. set cookie
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
