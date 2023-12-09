package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Certificate struct {
	Id         string `json:"id"`
	Driver_Fio string `json:"driver_fio"`
	Number     string `json:"number"`
	Age_Action string `json:"age_action"`
	Category   string `json:"category"`
}

func (c *Chaincode) SetCertificate(ctx contractapi.TransactionContextInterface, fio, number, age_action, category string) error {
	require := requireCategory(false, category, "")
	if !require {
		return fmt.Errorf("invalid category certificate")
	}

	driver, err := c.GetDriver(ctx, fio)
	if err != nil {
		return err
	}

	defer driver.Save(ctx)

	cert := new(Certificate)
	//если функция не выдаст ошибку, значит такой сертификат уже существует
	if err := Get(ctx, CERTIFICATE+number, cert); err == nil {
		return fmt.Errorf("ttis certificate already exists")
	}

	certificate := Certificate{
		Id:         CERTIFICATE + number,
		Driver_Fio: driver.Fio,
		Number:     number,
		Age_Action: age_action,
		Category:   category,
	}

	if err := Set(ctx, CERTIFICATE+number, certificate); err != nil {
		return err
	}

	driver.Certificate = &certificate

	return nil
}

// func (c *Client) SetCertificate(fio, number, age_action, category string) error {
// 	_, err := c.Transact("setCertificate", fio, number, age_action, category)
// 	return err
// }
