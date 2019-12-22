package main

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/dmitrychurkin/hotelier/server/prisma-generated/prisma-client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/rs/cors"
)

type contextKey struct {
	name string
}

// GinContextKey ...
var ginContextKey = &contextKey{"GinContextKey"}

func main() {
	var (
		resolvers = &Resolver{Prisma: prisma.New(nil)}
		server    = gin.Default()
	)
	server.RedirectTrailingSlash = true
	server.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"*"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"x-auth-token"},
	}))
	server.Use(contextToContextMiddleware())

	server.POST("/query", graphqlHandler(Config{Resolvers: resolvers}))
	server.GET("/", playgroundHandler())

	log.Fatal(server.Run())
}

func graphqlHandler(config Config) gin.HandlerFunc {
	h := handler.GraphQL(NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func contextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// Development only middleware
// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		cors.New(cors.Options{
// 			AllowedOrigins: []string{"*"},
// 			AllowedHeaders: []string{"*"},
// 		}).HandlerFunc(c.Writer, c.Request)

// 		c.Next()
// 	}
// }

// TODO: set up request rate limiter middleware

// GinContextFromContext - generic context receiver
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ginContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}
