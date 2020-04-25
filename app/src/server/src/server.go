package main

import (
	"fmt"
	"log"
  "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	fmt.Printf("Number activities: %v\n\n", smallNum)

	var i int64 = 0
	for ; i < smallNum; i++ {
		ref, err := contract.GetReference(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatalf("GetReference Failed: %v", err)
		}
		fmt.Printf("Activities Reference: %x\n", ref)

		activities, err := contract.GetActivities(&bind.CallOpts{}, ref)
		if err != nil {
			log.Fatalf("GetActivities Failed: %v", err)
		}

		thisVersion := activities.Version
		sliceVersion := thisVersion[:]
		version := hexutil.Encode(sliceVersion)
		stringVersion, _ := hexutil.Decode(version)

		thisTime := activities.GeneratedTime
		sliceTime := thisTime[:]
		time := hexutil.Encode(sliceTime)
		stringTime, _ := hexutil.Decode(time)

		thisLink := activities.LinkedData
		sliceLink := thisLink[:]
		link := hexutil.Encode(sliceLink)
		stringLink, _ := hexutil.Decode(link)

		fmt.Printf("Version: %s\n", stringVersion)
		fmt.Printf("Generated Time: %s\n", stringTime)
		fmt.Printf("Linked Data: %s\n\n", stringLink)
	}
}
