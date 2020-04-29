package main

import (
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ogady/find_the_right_answer/config"
	"github.com/ogady/find_the_right_answer/interface/graph"
	"github.com/ogady/find_the_right_answer/interface/graph/generated"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
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
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	tracer.Start(tracer.WithEnv(os.Getenv("ENV")))
	defer tracer.Stop()

	err := config.InitConf()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(gintrace.Middleware("FTRA-API"), cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"*",
		},

		MaxAge: 24 * time.Hour,
	}))

	r.POST("/", graphqlHandler())
	r.GET("/playground", playgroundHandler())
	r.Run(":80")
}
