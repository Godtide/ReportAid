package xml

import (
	"fmt"
	"log"
    "net/http"
  	"math/big"
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/configs"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Activities is a struct for holding the XML for the activities file
type Activities struct {
	Version 		string `xml:"version"`
	Time 				string `xml:"time"`
	LinkedData  string `xml:"linkedData"`
}

// GetActivites - output activies XML
func GetActivites(w http.ResponseWriter, r *http.Request) {

	conn, err := ethclient.Dial(string(configs.RinkebyAddr))
	if err != nil {
		log.Fatalf("Failed to connect to Rinkeby: %v", err)
	}

	contract, err := contracts.NewIATIActivities(common.HexToAddress(string(configs.ActivitiesContractAddr)), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate activities contract: %v", err)
	}

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("GetNumActivities Failed: %v", err)
	}
	smallNum := num.Int64()
	//fmt.Printf("Number activities: %v\n\n", smallNum)

	var i int64
	for ; i < smallNum; i++ {
		ref, err := contract.GetReference(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatalf("GetReference Failed: %v", err)
		}
		//fmt.Printf("Activities Reference: %x\n", ref)

		activities, err := contract.GetActivities(&bind.CallOpts{}, ref)
		if err != nil {
			log.Fatalf("GetActivities Failed: %v", err)
		}

		thisVersion := activities.Version
		sliceVersion := thisVersion[:]
		trimmedVersion := common.TrimRightZeroes(sliceVersion)
		version := hexutil.Encode(trimmedVersion)
		stringVersion, _ := hexutil.Decode(version)

		thisTime := activities.GeneratedTime
		sliceTime := thisTime[:]
		trimmedTime := common.TrimRightZeroes(sliceTime)
		time := hexutil.Encode(trimmedTime)
		stringTime, _ := hexutil.Decode(time)

		thisLink := activities.LinkedData
		sliceLink := thisLink[:]
		trimmedLink := common.TrimRightZeroes(sliceLink)
		link := hexutil.Encode(trimmedLink)
		stringLink, _ := hexutil.Decode(link)

		/* fmt.Printf("Version: %s\n", stringVersion)
		fmt.Printf("Generated Time: %s\n", stringTime)
		fmt.Printf("Linked Data: %s\n\n", stringLink)*/

		activitiesXML := &Activities{Version: string(stringVersion), Time: string(stringTime), LinkedData: string(stringLink)}
		thisXML, _ := xml.MarshalIndent(activitiesXML, "", " ")

		fmt.Fprintf(w, "%s\n\n", thisXML)
	}
}
