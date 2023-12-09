package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ProfiCoin uint

type Role struct {
	DPS  bool `json:"dps"`
	Bank bool `json:"bank"`
}

type Driver struct {
	Id                string       `json:"id"`
	Role              Role         `json:"role"`
	Fio               string       `json:"fio"`
	Certificate       *Certificate `json:"driver_certificate"`
	Transport         *Transport   `json:"transport"`
	Age_Stage         string       `json:"age_stage"`
	Balance_ProfiCoin ProfiCoin    `json:"balance_proficoin"`
}

func NewDriver(fio string, role Role, age_stage string) (*Driver) {
	if fio == "bank" {
		return &Driver{
			Id:                BANK + fio,
			Role:              role,
			Fio:               fio,
			Age_Stage:         age_stage,
			Balance_ProfiCoin: ProfiCoin(1000),
		}
	}

	return &Driver{
		Id:                DRIVER + fio,
		Role:              role,
		Fio:               fio,
		Age_Stage:         age_stage,
		Balance_ProfiCoin: ProfiCoin(50),
	}
}

func (c *Chaincode) GetDriver(ctx contractapi.TransactionContextInterface, fio string) (*Driver, error) {
	driver := new(Driver)

	if fio == "bank" || fio == "Bank" {
		return driver, Get(ctx, BANK+fio, driver)
	}

	return driver, Get(ctx, DRIVER+fio, driver)
}

func (d *Driver) Save(ctx contractapi.TransactionContextInterface) error {
	return Set(ctx, d.Id, d)
}

// func (c *Client) GetDriver(fio string) (*Driver, error) {
// 	driver := new(Driver)

// 	data, err := c.Call("getDriver", fio)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return driver, json.Unmarshal(data, &driver)
// }


