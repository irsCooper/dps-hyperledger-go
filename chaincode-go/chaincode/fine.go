package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Fine struct {
	Id    string    `json:"id"`
	Price ProfiCoin `json:"price"`
	Data  time.Time `json:"data"`
}

func newFine(ctx contractapi.TransactionContextInterface, numberCertificate string) (*Fine, error) {
	timeNow, id, err := newId(ctx)
	if err != nil {
		return nil, err
	}
	return &Fine{
		Id: FINE + numberCertificate + "@" + id,
		Price: ProfiCoin(10),
		Data: timeNow,
	}, nil
}

func (c *Chaincode) SetFine(ctx contractapi.TransactionContextInterface, fio, numberCertificate string) error {
	dps, err := c.GetDriver(ctx, fio)
	if err != nil {
		return err
	}

	if !dps.Role.DPS {
		return fmt.Errorf("invalid role dps")
	}

	certificate := new(Certificate)
	if err := Get(ctx, CERTIFICATE+numberCertificate, certificate); err != nil {
		return fmt.Errorf("invalid certificate: %v", err)
	}

	fine, err := newFine(ctx, numberCertificate)
	if err != nil {
		return err
	}

	if err := Set(ctx, fine.Id, fine); err != nil {
		return err
	}

	return nil
}

// func (c *Client) SetFine(fio, numberCertificate string) error {
// 	_, err := c.Transact("setFine", fio, numberCertificate)
// 	return err
// }







func (c *Chaincode) GetOneFine(ctx contractapi.TransactionContextInterface, id string) (*Fine, error) {
	fine := new(Fine)
	return fine, Get(ctx, id, fine)
}

// func (c *Client) GetOneFine(id string) (*Fine, error) {
// 	fineJSON, err := c.Call("getOneFine", id)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	fine := new(Fine)
// 	return fine, json.Unmarshal(fineJSON, &fine)
// }







func (c *Chaincode) GetAllMyFines(ctx contractapi.TransactionContextInterface, numberCertificate string) ([]*Fine, error) {
	iterator, err := ctx.GetStub().GetStateByRange(FINE+numberCertificate+"@", "")
	if err != nil {
		return nil, err
	}

	defer iterator.Close()

	var fines []*Fine

	for iterator.HasNext() {
		request, err := iterator.Next()
		if err != nil {
			return nil, err 
		}

		var fine Fine
		if err := json.Unmarshal(request.Value, &fine); err != nil {
			return nil, err
		}

		fines = append(fines, &fine)
	}

	return fines, nil
}

// func (c *Client) GetAllMyFines(numberCertificate string) ([]Fine, error) {
// 	finesJSON, err := c.Call("getAllMyFines", numberCertificate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fines := new([]Fine)
// 	return *fines, json.Unmarshal(finesJSON, fines)
// }














func (c *Chaincode) PayFine(ctx contractapi.TransactionContextInterface, fio, idFine string) error {
	fine := new(Fine)
	err := Get(ctx, idFine, fine)
	if err != nil {
		return err
	}

	driver, err := c.GetDriver(ctx, fio)
	if err != nil {
		return err
	}

	defer driver.Save(ctx)

	timeNow, _, err := newId(ctx)
	if err != nil {
		return err
	}

	var price ProfiCoin

	if time.Duration(fine.Data.Minute()) + time.Minute*5 > time.Duration(timeNow.Minute()) {
		price = ProfiCoin(5)
	} else {
		price = ProfiCoin(10)
	}

	if driver.Balance_ProfiCoin < price {
		return fmt.Errorf("invalid money")
	}

	bank, err := c.GetDriver(ctx, "bank") 
	if err != nil {
		return err
	}

	defer bank.Save(ctx)

	driver.Balance_ProfiCoin -= price
	bank.Balance_ProfiCoin   += price

	return ctx.GetStub().DelState(idFine)
}

// func (c *Client) PayFine(fio, idFine string) error {
// 	_, err := c.Transact("payFine", fio, idFine)
// 	return err
// }
