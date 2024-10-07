// main.go
package main

import (
	"log"

	"barr.io/api/routes"
	"barr.io/db"
	"barr.io/graph"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Initialize the database
	db.ConnectDatabase()

	// Set up the router
	routes.SetupRouter()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Fatal(http.ListenAndServe(":4000", nil))
}
