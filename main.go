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
	{
		apiV1XuperGroup.POST("/keypair/new", apiV1XuperKeypairNew)
		apiV1XuperGroup.POST("/account/new", apiV1XuperAccountNew)
		apiV1XuperGroup.GET("/balance", apiV1XuperAdminBalance)
		apiV1XuperGroup.GET("/balance/:address", apiV1XuperBalance)
		apiV1XuperGroup.POST("/faucet/:address", apiV1XuperFaucet)
		apiV1XuperGroup.GET("/transaction/:id", apiV1XuperQueryTx)
		apiV1XuperGroup.POST("/contract/deploy", apiV1XuperContractDeploy)
		apiV1XuperGroup.POST("/contract/invoke", apiV1XuperContractInvoke)
		apiV1XuperGroup.POST("/contract/query", apiV1XuperContractQuery)
		apiV1XuperGroup.GET("/addrconv/x2e/:address", apiV1XuperAddrconvX2e)
		apiV1XuperGroup.GET("/addrconv/e2x/:address", apiV1XuperAddrconvE2x)
		apiV1XuperGroup.POST("/collection/new", apiV1XuperCollectionNew)
		apiV1XuperGroup.POST("/collection/:cname/mint", apiV1XuperContractInvoke)
		apiV1XuperGroup.POST("/collection/:cname/transfer", apiV1XuperContractInvoke)
		apiV1XuperGroup.GET("/collection/:cname/info", apiV1XuperContractInvoke)
		// apiV1XuperGroup.GET("/collection/:cname/name", apiV1XuperContractInvoke)
		// apiV1XuperGroup.GET("/collection/:cname/description", apiV1XuperContractQuery)
		// apiV1XuperGroup.GET("/collection/:cname/owner", apiV1XuperContractQuery)
		apiV1XuperGroup.GET("/collection/:cname/tokenUri/:id", apiV1XuperContractQuery)
		// total issuance
		apiV1XuperGroup.GET("/collection/:cname/token/:id", apiV1XuperContractQuery)
		apiV1XuperGroup.GET("/collection/:cname/balanceOf/:address/:id", apiV1XuperContractQuery)
		// apiV1XuperGroup.GET("/template/:name/bin", apiV1XuperTemplateBin)
		// apiV1XuperGroup.GET("/template/:name/sol", apiV1XuperTemplateSol)
		// apiV1XuperGroup.GET("/template/:name/abi", apiV1XuperTemplateAbi)
		// apiV1XuperGroup.GET("/template/:name/bin", apiV1XuperTemplateBin)
	}

	apiV1Group := app.Group("/api/v1", swagin.Tags("JackieChainV1"))
	{
		apiV1Group.POST("/:chain/keypair/new", apiV1XuperKeypairNew)
		apiV1Group.POST("/:chain/account/new", apiV1XuperAccountNew)
		apiV1Group.GET("/:chain/balance", apiV1XuperAdminBalance)
		apiV1Group.GET("/:chain/balance/:address", apiV1XuperBalance)
		apiV1Group.POST("/:chain/faucet/:address", apiV1XuperFaucet)
		apiV1Group.GET("/:chain/transaction/:id", apiV1XuperQueryTx)
		apiV1Group.POST("/:chain/contract/deploy", apiV1XuperContractDeploy)
		apiV1Group.POST("/:chain/contract/invoke", apiV1XuperContractInvoke)
		apiV1Group.POST("/:chain/contract/query", apiV1XuperContractQuery)
	}
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
