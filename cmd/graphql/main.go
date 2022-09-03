package main

import (
	"log"
	"net/http"
	"os"

	"github.com/wander4747/adopet-backend/pkg/graph/directives"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/wander4747/adopet-backend/pkg/graph/generated"
	"github.com/wander4747/adopet-backend/pkg/graph/resolver"
	"github.com/wander4747/adopet-backend/pkg/service"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	services := service.NewService()

	config := generated.Config{Resolvers: &resolver.Resolver{
		Services: services,
	}}
	config.Directives.Validation = directives.Validate
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
