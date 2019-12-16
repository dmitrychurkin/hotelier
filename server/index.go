package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/dmitrychurkin/hotelier/server/prisma-generated/prisma-client"
	"github.com/rs/cors"
)

// DataCtxKey ...
var dataCtxKey = &contextKey{"contextData"}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: client,
	}

	apply := func(middlewares ...http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dataCtxKey, &contextData{
				Req: r,
				Res: w,
			})

			r = r.WithContext(ctx)

			for _, m := range middlewares {
				m(w, r)
			}
		}
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query",
		apply(
			cors.New(cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedHeaders: []string{"*"},
			}).HandlerFunc,
			handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})),
		),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ContextData ...
type contextData struct {
	Req *http.Request
	Res http.ResponseWriter
}

type contextKey struct {
	name string
}
