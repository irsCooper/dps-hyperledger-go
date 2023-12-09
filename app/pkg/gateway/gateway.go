package gateway


import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var (
	Contract *gateway.Contract

	Organisations map[uint8]*Org
)

type Org struct {
	Number uint8
	Name   string
	Wallet *gateway.Wallet
}

func Init() {
	Organisations = make(map[uint8]*Org)
	os.RemoveAll("bank")
	os.RemoveAll("users")
	os.RemoveAll("keystore")
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	org1, org2 := NewOrg(1, "Bank"), NewOrg(2, "Users")

	u1, u2 := []string{"bank"}, []string{"Иванов_Иван_Иванович", "Семенов_Семен_Семенович", "Петров_Петр_Петрович"}

	for _, user := range u1 {
		org1.RegisterUser(user)
	}
	for _, user := range u2 {
		org2.RegisterUser(user)
	}

	// for i := 1; i < 9; i++ {
	// 	org2.RegisterUser(fmt.Sprint(i))
	// }

	// Настроить пути при запуске
	ccpPath := filepath.Join(
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(org1.Wallet, "bank"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}

	// Подключение к сети
	nw, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %s", err)
	}
	log.Println("network -> " + fmt.Sprint(nw))

	// Получаем контракт
	Contract = nw.GetContract("basic")
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	log.Println("Contract ->" + fmt.Sprint(Contract))

}

func NewOrg(number uint8, name string) *Org {
	w, err := gateway.NewFileSystemWallet(strings.ToLower(name))
	if err != nil {
		log.Fatal(err)
	}

	org := Org{Number: number, Name: name, Wallet: w}

	Organisations[number] = &org

	return &org
}

func GetOrg(number uint8) (*Org, error) {
	org := Organisations[number]
	if org == nil {
		return nil, fmt.Errorf("invalid number organisations") 
	}
	return org, nil
}

// RegisterUser registration user in the organization
// @param: user to register
func (o *Org) RegisterUser(user string) error {
	credPath := filepath.Join(
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		fmt.Sprintf("org%d.example.com", o.Number),
		"users",
		fmt.Sprintf("User1@org%d.example.com", o.Number),
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(o.Name, string(cert), string(key))

	return o.Wallet.Put(user, identity)
}
