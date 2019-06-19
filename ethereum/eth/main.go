package main

import (
	"github.com/button-tech/utils-node-tool/ethereum/eth/handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	gin.SetMode(gin.ReleaseMode)

	eth := r.Group("/eth")

	{
		eth.GET("/balance/:address", handlers.GetBalance)

		eth.GET("/transactionFee", handlers.GetTxFee)

		eth.GET("/gasPrice", handlers.GetGasPrice)

		eth.GET("/tokenBalance/:sc-address/:address", handlers.GetTokenBalance)

		eth.POST("/estimateGas", handlers.GetEstimateGas)

		eth.POST("/balances", handlers.GetBalances)

		eth.POST("/tokenBalances", handlers.GetTokenBalances)
	}

	if err := r.Run(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
