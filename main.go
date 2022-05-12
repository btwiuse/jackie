package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/swagin"
	// "github.com/long2ice/swagin/security"
)

func getPort() string {
	const DEFAULT_PORT = "8085"
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return ":" + port
}

func main() {
	// Use customize Gin engine
	r := gin.New()

	// Registering func(c *gin.Context) is accepted,
	// but the OpenAPI generator will ignore the operation and it won't appear in the specification.
	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	*/

	app := swagin.NewFromEngine(r, NewSwagger())
	// subApp := swagin.NewFromEngine(r, NewSwagger())
	// apiV1 := swagin.NewFromEngine(r, NewSwagger())

	// apiV1XuperGroup := apiV1.Group("/xuper", swagin.Tags("XuperChain"))
	apiV1XuperGroup := app.Group("/api/v1/xuper", swagin.Tags("XuperChainV1"))
	// apiV1XuperGroup.GET("/hello", apiV1XuperHello)
	apiV1XuperGroup.POST("/keypair/new", apiV1XuperKeypairNew)
	apiV1XuperGroup.GET("/balance", apiV1XuperAdminBalance)
	apiV1XuperGroup.GET("/balance/:address", apiV1XuperBalance)

	/*
		You can use default Gin engin:
			app := swagin.New(NewSwagger())
			subApp := swagin.New(NewSwagger())
	*/

	// subApp.GET("/noModel", noModel)
	// app.Mount("/sub", subApp)
	// app.Mount("/api/v1", apiV1)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	app.GET("/health", health)
	/*
		queryGroup := app.Group("/query", swagin.Tags("Query"))
		queryGroup.GET("/list", queryList)
		queryGroup.GET("/:id", queryPath)
		queryGroup.DELETE("", query)

		app.GET("/noModel", noModel)

		formGroup := app.Group("/form", swagin.Tags("Form"), swagin.Security(&security.Bearer{}))
		formGroup.POST("/encoded", formEncode)
		formGroup.PUT("", body)
		formGroup.POST("/file", file)
	*/

	port := getPort()

	log.Printf("Now you can visit http://127.0.0.1%v/docs or http://127.0.0.1%v/redoc to see the api docs. Have fun!", port, port)
	log.Fatalln(app.Run(port))
}
