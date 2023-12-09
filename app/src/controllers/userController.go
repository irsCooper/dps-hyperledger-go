package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irsCooper/dps-hyperledger-go/app/pkg/gateway"
	"github.com/irsCooper/dps-hyperledger-go/app/src/utils"
	"github.com/irsCooper/dps-hyperledger-go/chaincode-go/chaincode"
)


func SignIn(c *gin.Context) {
	fio := c.Request.FormValue("fio")
	age_stage := c.Request.FormValue("age_stage")

	_, err := gateway.Contract.SubmitTransaction("SignIn", fio, age_stage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to submit transaction: " + err.Error(),
		})
	}

	orgUsers, err := gateway.GetOrg(2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to get organization: " + err.Error(),
		})
	}

	if err := orgUsers.RegisterUser(fio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to create certificate: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"successfully submit transaction",
	})
}

func Login(c *gin.Context) {
	fio := c.Request.FormValue("fio")

	driverJSON, err := gateway.Contract.EvaluateTransaction("Login", fio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to evaluate transaction: " + err.Error(),
		})
	}

	driver := new(chaincode.Driver)
	if err := utils.FromByteToStruct(driverJSON, driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to unmarshall struct: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"succesfully evaluate transaction",
		"driver": driver,
	})
}
