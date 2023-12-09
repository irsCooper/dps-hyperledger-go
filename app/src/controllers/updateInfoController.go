package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irsCooper/dps-hyperledger-go/app/pkg/gateway"
)

func SetCertificate(c *gin.Context) {
	fio := c.Request.FormValue("fio")
	number := c.Request.FormValue("number")
	age_action := c.Request.FormValue("age_action")
	category := c.Request.FormValue("category")

	_, err := gateway.Contract.SubmitTransaction("SetCertificate", fio, number, age_action, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to submit transaction: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully submit transaction",
	})
}

func SetTransport(c *gin.Context) {
	fio := c.Request.FormValue("fio")
	category := c.Request.FormValue("category")
	price := c.Request.FormValue("price")
	age := c.Request.FormValue("age")

	_, err := gateway.Contract.SubmitTransaction("SetTransport", fio, category, price, age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to submit transaction: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully submit transaction",
	})
}


func SetFine(c *gin.Context) {
	fio := c.Request.FormValue("fio")
	numberCertificate := c.Request.FormValue("numberCertificate")

	_, err := gateway.Contract.SubmitTransaction("SetFine", fio, numberCertificate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to submit transaction: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully submit transaction",
	})
}