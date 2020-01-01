package main

import (
	"context"

	"github.com/dmitrychurkin/hotelier/server/graph/generated"
	"github.com/dmitrychurkin/hotelier/server/models"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
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
	return Login(ctx, r.Prisma, email, password)
}
func (r *mutationResolver) Signup(ctx context.Context, email string, firstName *string, lastName *string, password string, confirmPassword string) (*models.User, error) {
	return Signup(ctx, r.Prisma, email, firstName, lastName, password, confirmPassword)
}
func (r *mutationResolver) SendPasswordResetLink(ctx context.Context, email string, path string) (*bool, error) {
	return SendPasswordResetLink(ctx, r.Prisma, email, path)
}
func (r *mutationResolver) ResetPassword(ctx context.Context, email string, password string, confirmPassword string, passwordResetToken string) (*models.User, error) {
	return ResetPassword(ctx, r.Prisma, email, password, confirmPassword, passwordResetToken)
}
func (r *mutationResolver) Logout(ctx context.Context) (*bool, error) {
	return Logout(ctx)
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
	return User(ctx, r.Prisma)
}
func (r *queryResolver) ResetPasswordCreds(ctx context.Context, passwordResetToken string) (*models.ResetPasswordCreds, error) {
	return ResetPasswordCreds(ctx, r.Prisma, passwordResetToken)
}
