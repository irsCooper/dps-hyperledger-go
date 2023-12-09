package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


type Transport struct {
	Category string    `json:"category"`
	Price    ProfiCoin `json:"price"`
	Age      uint      `json:"age"`
}

func (c *Chaincode) SetTransport(ctx contractapi.TransactionContextInterface, fio, category string, price, age uint) error {
	driver, err := c.GetDriver(ctx, fio)
	if err != nil {
		return err
	}

	defer driver.Save(ctx)
	
	require := requireCategory(true, driver.Certificate.Category, category)
	if !require {
		return fmt.Errorf("invalid transport category")
	}

	transport := Transport{
		Category: category,
		Price: ProfiCoin(price),
		Age: age,
	}

	driver.Transport = &transport

	return nil
}


// func (c *Client) SetTransport(fio, category, price, age string) error {
// 	_, err := c.Transact("setTransport", fio, category, price, age)
// 	return err
// }