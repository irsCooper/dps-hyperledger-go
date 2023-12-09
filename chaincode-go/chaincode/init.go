package chaincode

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

func (c *Chaincode) Init(ctx contractapi.TransactionContextInterface) {
	NewDriver("bank", Role{DPS: false, Bank: true}, "0").Save(ctx)
	NewDriver("Иванов_Иван_Иванович", Role{DPS: true, Bank: false}, "2").Save(ctx)
	NewDriver("Семенов_Семен_Семенович", Role{DPS: false, Bank: false}, "5").Save(ctx)
	NewDriver("Петров_Петр_Петрович", Role{DPS: false, Bank: false}, "9").Save(ctx)
}