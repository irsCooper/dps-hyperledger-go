package chaincode

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type Chaincode struct {
	contractapi.Contract
}

// type Client struct {
// 	Call, Transact func(string, ...string) ([]byte, error)
// }

// func NewClient(call, transact func(string, ...string) ([]byte, error)) *Client {
// 	return &Client{
// 		Call:     call,
// 		Transact: transact,
// 	}
// }
