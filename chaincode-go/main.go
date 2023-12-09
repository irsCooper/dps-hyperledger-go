package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/irsCooper/dps-hyperledger-go/chaincode-go/chaincode"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&chaincode.Chaincode{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
