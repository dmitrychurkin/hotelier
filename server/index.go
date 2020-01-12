package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/dmitrychurkin/hotelier/server/graph/generated"
	"github.com/dmitrychurkin/hotelier/server/prisma-client"
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
		resolvers     = &Resolver{Prisma: prisma.New(nil)}
		server        = gin.Default()
		allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
	)

	if len(allowedOrigin) == 0 {
		allowedOrigin = "http://localhost:3000"
	}

	server.RedirectTrailingSlash = true

	server.Use(
		// comment if using Playground
		cors.New(cors.Config{
			AllowOrigins:     []string{allowedOrigin},
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type"},
		}),
		contextToContextMiddleware(),
	)
	// server.Use(contextToContextMiddleware())

	server.POST("/query", graphqlHandler(generated.Config{Resolvers: resolvers}))
	server.GET("/", playgroundHandler())

	log.Fatal(server.Run())
}

func graphqlHandler(config generated.Config) gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(config))

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
