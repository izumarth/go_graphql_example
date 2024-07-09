package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/izumarth/go-graphql-example/graph"
	"github.com/izumarth/go-graphql-example/graph/services"
	"github.com/izumarth/go-graphql-example/internal"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const (
	defaultPort = "8080"
	dbFile      = "./db/mygraphql.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbFile))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	services := services.New(db)

	srv := handler.NewDefaultServer(
		internal.NewExecutableSchema(
			internal.Config{
				Resolvers: &graph.Resolver{
					Srv:     services,
					Loaders: graph.NewLoaders(services),
				},
				Complexity: graph.ComplexityConfig(),
			},
		),
	)
	srv.Use(extension.FixedComplexityLimit(50))

	boil.DebugMode = true

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
