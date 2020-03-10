package main

import (
	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/ogady/find_the_right_answer/config"
	"github.com/ogady/find_the_right_answer/interface/graph"
	"github.com/ogady/find_the_right_answer/interface/graph/generated"
)

// Defining the Graphql handler

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/")
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

func main() {
	err := config.InitConf()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/", graphqlHandler())
	r.GET("/playground", playgroundHandler())
	r.Run(":80")
}
