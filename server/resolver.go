package main

import (
	"context"

	"github.com/dmitrychurkin/hotelier/server/graph/generated"
	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
	"github.com/dmitrychurkin/hotelier/server/resolvers/auth"
)

// Resolver with prisma client
type Resolver struct {
	Prisma *prisma.Client
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query ...
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*models.User, error) {
	// return Login(ctx, r.Prisma, email, password)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.Login(&ctx, gc, r.Prisma, email, password)
}
func (r *mutationResolver) Signup(ctx context.Context, email string, firstName *string, lastName *string, password string) (*models.User, error) {
	// return Signup(ctx, r.Prisma, email, firstName, lastName, password, confirmPassword)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.Signup(&ctx, gc, r.Prisma, email, firstName, lastName, password)
}
func (r *mutationResolver) SendPasswordResetLink(ctx context.Context, email string, path string) (*bool, error) {
	// return SendPasswordResetLink(ctx, r.Prisma, email, path)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.SendPasswordResetLink(&ctx, gc, r.Prisma, email, path)
}
func (r *mutationResolver) ResetPassword(ctx context.Context, email string, password string, passwordResetToken string) (*models.User, error) {
	// return ResetPassword(ctx, r.Prisma, email, password, confirmPassword, passwordResetToken)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.ResetPassword(&ctx, gc, r.Prisma, email, password, passwordResetToken)
}
func (r *mutationResolver) Logout(ctx context.Context) (*bool, error) {
	// return Logout(ctx)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.Logout(&ctx, gc)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserByID(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	// return User(ctx, r.Prisma)
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.User(&ctx, gc, r.Prisma)
}
func (r *queryResolver) ResetPasswordCred(ctx context.Context, passwordResetToken string) (string, error) {
	// return ResetPasswordCreds(ctx, r.Prisma, passwordResetToken)
	return auth.ResetPasswordCred(&ctx, r.Prisma, passwordResetToken)
}
