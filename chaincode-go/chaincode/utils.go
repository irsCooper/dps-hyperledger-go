package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "golang.org/x/crypto/bcrypt"
)

func Get(ctx contractapi.TransactionContextInterface, id string, asset interface{}) error {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return err
	}

	if assetJSON == nil {
		return fmt.Errorf("asset %s doesn't exist from world state", id)
	}

	return json.Unmarshal(assetJSON, asset)
}

func Set(ctx contractapi.TransactionContextInterface, id string, asset interface{}) error {
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("error marshalling asset %s", asset)
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func newId(ctx contractapi.TransactionContextInterface) (time.Time, string, error) {
	now, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return time.Now(), "", fmt.Errorf("invalid timestamp: %v", err)
	}
	_now := now.AsTime()

	return _now, fmt.Sprint(_now.Unix()*1e3 + int64(now.Nanos)/1e6), nil
}


//concair - нужно ли проверять совпадение с категорией транспорта
func requireCategory(concair bool, certCategory, transportCategory string) (bool) {
	if concair {
		return certCategory == transportCategory
	} 

	if certCategory == "A" || certCategory == "B" || certCategory == "C" {
		return true
	}

	return false
}


//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

//хайпер не любит пробелы в получемых аргументах, по этому
//ПЕРЕНЕСТИ ФУНКЦИЮ НА ФРОНТ ЧАСТЬ

// func requireDriverFio(fio string) (string, error) {
// 	split_fio := strings.Split(fio, " ")
// 	if len(split_fio) != 3 {
// 		return "", fmt.Errorf("Invalid Fio specification")
// 	}
// 	return split_fio[0] + "_" + split_fio[1] + "_" + split_fio[2], nil
// }

//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

