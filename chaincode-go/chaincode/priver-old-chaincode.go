package chaincode

// v1.0 оставлен на память

// import (
// 	// "encoding/json"
// 	// "fmt"
// 	// "time"

// 	"github.com/hyperledger/fabric-contract-api-go/contractapi"
// )



// const bank_Id = "bank@bank"

// func (c *Chaincode) Init(ctx contractapi.TransactionContextInterface) error {
// 	drivers := []Driver{
// 		{
// 			Id:                "bank@bank",
// 			Role:              Role{DPS: false, Bank: true},
// 			Fio:               "Bank",
// 			Balance_ProfiCoin: ProfiCoin(1000),
// 		},

// 		{
// 			Id:                "driver@Иванов_Иван_Иванович",
// 			Role:              Role{DPS: true, Bank: false},
// 			Fio:               "Иванов_Иван_Иванович",
// 			Age_Stage:         "2",
// 			Balance_ProfiCoin: ProfiCoin(50),
// 		},

// 		{
// 			Id:                "driver@Семенов_Семен_Семенович",
// 			Fio:               "Семенов_Семен_Семенович",
// 			Age_Stage:         "5",
// 			Balance_ProfiCoin: ProfiCoin(50),
// 		},

// 		{
// 			Id:                "driver@Петров_Петр_Петрович",
// 			Fio:               "Петров_Петр_Петрович",
// 			Age_Stage:         "10",
// 			Balance_ProfiCoin: ProfiCoin(50),
// 		},
// 	}

// 	for _, driver := range drivers {
// 		if err := Set(ctx, driver.Id, driver); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // функция получает фио в формате: фамилия_имя_отчество
// // func (c *Chaincode) Registration(ctx contractapi.TransactionContextInterface, fio, age_stage string) error {
// // 	driver := Driver{
// // 		Id:                "driver@" + fio,
// // 		Fio:               fio,
// // 		Age_Stage:         age_stage,
// // 		Balance_ProfiCoin: ProfiCoin(50),
// // 	}

// // 	return Set(ctx, driver.Id, driver)
// // }

// func (c *Chaincode) Authorization(ctx contractapi.TransactionContextInterface, fio string) (*Driver, error) {
// 	var id string

// 	if fio == "Bank" {
// 		id = bank_Id
// 	} else {
// 		id = "driver@" + fio
// 	}

// 	driver := new(Driver)

// 	if err := Get(ctx, id, driver); err != nil {
// 		return nil, err
// 	}

// 	return driver, nil
// }

// // +
// // func (c *Chaincode) SetCertificate(ctx contractapi.TransactionContextInterface, fio, number, age_action, category string) error {
// // 	driver, err := c.Authorization(ctx, fio)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	var _category string

// // 	if category == "A" || category == "B" || category == "C" {
// // 		_category = category
// // 	} else {
// // 		return fmt.Errorf("Invalid certificate category")
// // 	}

// // 	test_certificate := new(Certificate)

// // 	//если мы получили ошибку из вызываемой функции, значит такого сертификата ещё не добавлено в систему и мы можем продолжать выполнение функции
// // 	if err = Get(ctx, "certificate@"+number, test_certificate); err == nil {
// // 		return fmt.Errorf("this is not your certificate")
// // 	}

// // 	certificate := Certificate{
// // 		Id:         "certificate@" + number,
// // 		Driver_ID:  driver.Id,
// // 		Number:     number,
// // 		Age_Action: age_action,
// // 		Category:   _category,
// // 	}

// // 	if err = Set(ctx, certificate.Id, certificate); err != nil {
// // 		return err
// // 	}

// // 	newDriver := Driver{
// // 		Id:                driver.Id,
// // 		Role:              driver.Role,
// // 		Fio:               driver.Fio,
// // 		Certificate:       &certificate,
// // 		Transport:         driver.Transport,
// // 		Age_Stage:         driver.Age_Stage,
// // 		// Amount_Fine:       driver.Amount_Fine,
// // 		Balance_ProfiCoin: driver.Balance_ProfiCoin,
// // 	}

// // 	return Set(ctx, newDriver.Id, newDriver)
// // }

// // +
// // func (c *Chaincode) SetTransport(ctx contractapi.TransactionContextInterface, fio, category string, price, age uint) error {
// // 	driver, err := c.Authorization(ctx, fio)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	var _category string

// // 	if category == "A" || category == "B" || category == "C" {
// // 		_category = category
// // 	} else {
// // 		return fmt.Errorf("Invalid trancport category certificate ")
// // 	}

// // 	if driver.Certificate.Category != _category {
// // 		return fmt.Errorf("Invalid category certificate driver")
// // 	}

// // 	transport := Transport{
// // 		Category: category,
// // 		Price:    ProfiCoin(price),
// // 		Age:      age,
// // 	}

// // 	newDriver := Driver{
// // 		Id:                driver.Id,
// // 		Role:              driver.Role,
// // 		Fio:               driver.Fio,
// // 		Certificate:       driver.Certificate,
// // 		Transport:         &transport,
// // 		Age_Stage:         driver.Age_Stage,
// // 		// Amount_Fine:       driver.Amount_Fine,
// // 		Balance_ProfiCoin: driver.Balance_ProfiCoin,
// // 	}

// // 	return Set(ctx, newDriver.Id, newDriver)
// // }


// // +
// // попробовать раскоментить и ещё раз потестить
// // func (c *Chaincode) SetFine(ctx contractapi.TransactionContextInterface, dps_fio, number_certificate string) error {
// // 	dps, err := c.Authorization(ctx, dps_fio)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	if !dps.Role.DPS {
// // 		return fmt.Errorf("You don't have enough rights to execute")
// // 	}

// // 	// certificate_driver := new(Certificate)

// // 	// if err = getAssetFromWorldState(ctx, "certificate@"+number_certificate, certificate_driver); err != nil {
// // 	// 	return fmt.Errorf("Failed to get certificate: %v", err)
// // 	// }

// // 	// driver := new(Driver)

// // 	// if err = getAssetFromWorldState(ctx, certificate_driver.Driver_ID, driver); err != nil {
// // 	// 	return fmt.Errorf("Failed to get driver: %v", err)
// // 	// }

// // 	// newDriver := Driver{
// // 	// 	Id:                driver.Id,
// // 	// 	Role:              driver.Role,
// // 	// 	Fio:               driver.Fio,
// // 	// 	Certificate:       driver.Certificate,
// // 	// 	Transport:         driver.Transport,
// // 	// 	Age_Stage:         driver.Age_Stage,
// // 	// 	Amount_Fine:       driver.Amount_Fine + 1,
// // 	// 	Balance_ProfiCoin: driver.Balance_ProfiCoin,
// // 	// }

// // 	// if err = setAssetFromWorldState(ctx, newDriver.Id, newDriver); err != nil {
// // 	// 	return fmt.Errorf("Fataled to set new driver: %v", err)
// // 	// }

// // 	now, id, err := genNewId(ctx)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	fine := Fine{
// // 		Id:    "fine@" + number_certificate + "@" + id,
// // 		Price: ProfiCoin(10),
// // 		Data:  now,
// // 	}

// // 	return Set(ctx, fine.Id, fine)
// // }

// // +
// func (c *Chaincode) GetOneMyFine(ctx contractapi.TransactionContextInterface, id_fine string) (*Fine, error) {
// 	fine := new(Fine)

// 	if err := Get(ctx, id_fine, fine); err != nil {
// 		return nil, err
// 	}

// 	return fine, nil
// }

// // +
// // func (c *Chaincode) GetAllMyFines(ctx contractapi.TransactionContextInterface, number_certificate string) ([]*Fine, error) {
// // 	iterator, err := ctx.GetStub().GetStateByRange("fine@"+number_certificate+"@", "")
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	defer iterator.Close()

// // 	var fines []*Fine

// // 	for iterator.HasNext() {
// // 		request, err := iterator.Next()
// // 		if err != nil {
// // 			return nil, err
// // 		}

// // 		var fine Fine

// // 		if err = json.Unmarshal(request.Value, &fine); err != nil {
// // 			return nil, err
// // 		}

// // 		fines = append(fines, &fine)
// // 	}

// // 	return fines, nil
// // }

// // +
// // func (c *Chaincode) ToPayOneFine(ctx contractapi.TransactionContextInterface, id_fine, fio string) error {
// // 	driver, err := c.Authorization(ctx, fio)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	fine, err := c.GetOneMyFine(ctx, id_fine)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	var price_fine ProfiCoin

// // 	//ко времени, в которое был получен штраф, добавим 5 минут
// // 	term_to_pay := fine.Data.Second() + int(time.Second)*60*5

// // 	if time.Now().Second() <= term_to_pay {
// // 		price_fine = ProfiCoin(5)
// // 	} else {
// // 		price_fine = ProfiCoin(10)
// // 	}

// // 	if driver.Balance_ProfiCoin < price_fine {
// // 		return fmt.Errorf("Invalid money")
// // 	}

// // 	newDriver := Driver{
// // 		Id:                driver.Id,
// // 		Role:              driver.Role,
// // 		Fio:               driver.Fio,
// // 		Certificate:       driver.Certificate,
// // 		Transport:         driver.Transport,
// // 		Age_Stage:         driver.Age_Stage,
// // 		// Amount_Fine:       driver.Amount_Fine - 1,
// // 		Balance_ProfiCoin: driver.Balance_ProfiCoin - price_fine,
// // 	}

// // 	if err = Set(ctx, newDriver.Id, newDriver); err != nil {
// // 		return nil
// // 	}

// // 	bank, err := c.Authorization(ctx, "Bank")
// // 	if err != nil {
// // 		return err
// // 	}

// // 	newBank := Driver{
// // 		Id:                bank.Id,
// // 		Role:              bank.Role,
// // 		Fio:               bank.Fio,
// // 		Certificate:       bank.Certificate,
// // 		Transport:         bank.Transport,
// // 		Age_Stage:         bank.Age_Stage,
// // 		Balance_ProfiCoin: bank.Balance_ProfiCoin + price_fine,
// // 	}

// // 	if err = Set(ctx, newBank.Id, newBank); err != nil {
// // 		return nil
// // 	}

// // 	return ctx.GetStub().DelState(id_fine)
// // }
