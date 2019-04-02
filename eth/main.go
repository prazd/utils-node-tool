package main

import (
	"log"
	"os"

	_ "github.com/button-tech/utils-node-tool/eth/docs"
	"github.com/button-tech/utils-node-tool/eth/handlers"
	"github.com/button-tech/utils-node-tool/eth/handlers/storage"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prazd/round-robin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {

	// @title Swagger BUTTON Node API
	// @version 1.0
	// @description This is BUTTON Node API responseModels documentation

	// @contact.name API Support
	// @contact.email nk
	// ap@buttonwallet.com

	// @license.name MIT
	// @license.url https://opensource.org/licenses/MIT

	// @BasePath /

	// must add addresses to slice
	var EthNodes = []string{}

	rr, err := roundrobin.New(EthNodes)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// Round Robin middleware
	r.Use(func(c *gin.Context) {
		storage.EthNodeAddress.Set(rr.Next())
	})

	gin.SetMode(gin.ReleaseMode)

	r.GET("/eth/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/eth/balance/:address", handlers.GetBalance)

	r.GET("/eth/transactionFee", handlers.GetTxFee)

	r.GET("/eth/gasPrice/", handlers.GetGasPrice)

	r.GET("/eth/tokenBalance/:sc-address/:address", handlers.GetTokenBalance)

	r.POST("/eth/balances", handlers.GetBalances)

	r.POST("/eth/tokenBalances", handlers.GetTokenBalances)

	if err = r.Run(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
