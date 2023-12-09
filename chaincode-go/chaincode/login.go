package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (c *Chaincode) SignIn(ctx contractapi.TransactionContextInterface, fio, age_stage string) error {
	_, err := c.GetDriver(ctx, fio)
	if err == nil {
		return fmt.Errorf("this driver already exists: %v", err)
	}

	if err != nil {
		return err
	}

	driver := Driver{
		Id: DRIVER+fio,
		Role: Role{DPS: false, Bank: false},
		Fio: fio,
		Age_Stage: age_stage,
		Balance_ProfiCoin: ProfiCoin(50),
	}

	return driver.Save(ctx)
}

// func (c *Client) SignIn(fio string) error {
// 	_, err := c.Transact("signIn", fio)
// 	return err
// }



func (c *Chaincode) Login(ctx contractapi.TransactionContextInterface, fio string) (*Driver, error) {
	driver, err := c.GetDriver(ctx, fio)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

// func (c *Client) Login(fio string) (*Driver, error) {
// 	driverJSON, err := c.Call("login", fio)
// 	if err != nil {
// 		return nil, err
// 	}

// 	driver := new(Driver)
// 	return driver, json.Unmarshal(driverJSON, driver)
// }