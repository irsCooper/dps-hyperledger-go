package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irsCooper/dps-hyperledger-go/app/pkg/gateway"
	"github.com/irsCooper/dps-hyperledger-go/app/src/controllers"
)

func main() {
	gateway.Init()

	router := gin.Default()

	router.GET("/login", controllers.Login)
	router.GET("/onefine", controllers.GetOneFine)
	router.GET("/allmyfines", controllers.GetAllMyFines)


	router.POST("/signin", controllers.SignIn)
	router.POST("/setcertificate", controllers.SetCertificate)
	router.POST("/settransport", controllers.SetTransport)
	router.POST("/setfine", controllers.SetFine)

	router.Run(":8080")
}