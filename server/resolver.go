package main

import (
	"context"
	"errors"
	"fmt"

	prisma "github.com/dmitrychurkin/hotelier/server/prisma-generated/prisma-client"
)

// Resolver with prisma client
type Resolver struct {
	Prisma *prisma.Client
}

// Mutation ...
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query ...
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*prisma.User, error) {
	return LoginHandler(ctx, r.Prisma, email, password)
}
func (r *mutationResolver) Signup(ctx context.Context, email string, firstName *string, lastName *string, password string, confirmPassword string) (*prisma.User, error) {
	return SignupHandler(ctx, r.Prisma, email, firstName, lastName, password, confirmPassword)
}
func (r *mutationResolver) SendPasswordResetLink(ctx context.Context, email string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) ResetPassword(ctx context.Context, email string, password string, confirmPassword string, passwordResetToken string) (*prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) UserByID(ctx context.Context, id string) (*prisma.User, error) {
	// panic("not implemented")
	contextData := ctx.Value(dataCtxKey).(*contextData)
	contextData.Res.Header().Set("Access-Control-Expose-Headers", "*")
	contextData.Res.Header().Set("x-access-token", "12345")
	contextData.Res.Header().Set("x-refresh-token", "54321")
	fmt.Println(contextData.Req.Header.Get("x-test-context"))
	return nil, errors.New("Work in progress :)")
}
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	// panic("not implemented")
	contextData := ctx.Value(dataCtxKey).(*contextData)
	fmt.Printf("%+v\n", contextData)
	// Loop over header names
	for name, values := range contextData.Req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	// //Iterate over all header fields
	// for k, v := range contextData.Req.Header {
	// 	fmt.Fprintf(contextData.Res, "Header field %q, Value %q\n", k, v)
	// }
	// fmt.Fprintf(contextData.Res, "Host = %q\n", contextData.Req.Host)
	// fmt.Fprintf(contextData.Res, "RemoteAddr= %q\n", contextData.Req.RemoteAddr)
	return nil, errors.New("Invalid request")
}
