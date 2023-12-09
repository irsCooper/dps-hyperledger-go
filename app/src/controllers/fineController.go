package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irsCooper/dps-hyperledger-go/app/pkg/gateway"
	"github.com/irsCooper/dps-hyperledger-go/app/src/utils"
	"github.com/irsCooper/dps-hyperledger-go/chaincode-go/chaincode"
)

func GetOneFine(c *gin.Context) {
	id := c.Request.FormValue("id")

	fineJSON, err := gateway.Contract.EvaluateTransaction("GetOneFine", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to evaluate transaction: " + err.Error(),
		})
	}

	fine := new(chaincode.Fine)
	if err := utils.FromByteToStruct(fineJSON, fine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to unmarshall struct: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully evaluate transaction",
		"fine":  fine,
	})
}


func GetAllMyFines(c *gin.Context) {
	numberCertificate := c.Request.FormValue("numberCertificate")

	finesJSON, err := gateway.Contract.EvaluateTransaction("GetAllMyFines", numberCertificate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to evaluate transaction: " + err.Error(),
		})
	}

	fines := new([]chaincode.Fine)
	if err := utils.FromByteToStruct(finesJSON, fines); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to unmarshall struct: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully evaluate transaction",
		"fines":  fines,
	})
}