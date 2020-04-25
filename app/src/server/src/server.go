package main

import (
	"fmt"
	"log"
  "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	conn, err := ethclient.Dial("https://rinkeby.infura.io/v3/f4851ba4a9e546e7a1f16bf8e14f0a92")
	if err != nil {
		log.Fatalf("Failed to connect to Rinkeby: %v", err)
	}

	contract, err := NewIATIActivities(common.HexToAddress("0x1b61bE4D8f631afE6F24F595Df73D5275050B231"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("GetNumActivities Failed: %v", err)
	}
	smallNum := num.Int64()
	fmt.Println("Number activities:", smallNum)

	var i int64 = 0
	for ; i < smallNum; i++ {
		ref, err := contract.GetReference(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatalf("GetReference Failed: %v", err)
		}
		pass := fmt.Sprintf("%x", ref)
		fmt.Println(pass)

		orgActivities, err := contract.GetActivities(&bind.CallOpts{}, ref)
		if err != nil {
			log.Fatalf("GetActivities Failed: %v", err)
		}
		fmt.Println("Activities:", orgActivities)
	}
}
