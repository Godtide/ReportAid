package xml

import (
	"fmt"
	"log"
    "net/http"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// Activities - the XML for the activities file
type Activities struct {
	Version 		string `xml:"version"`
	Time 				string `xml:"time"`
	LinkedData  string `xml:"linkedData"`
}

// NumActivities - get the total number of activities
type NumActivities struct {
	Total 		int64 `xml:"total"`
}

// GetNumActivites - Get the total number of activities
func numActivites(contract *contracts.IATIActivities) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("GetNumActivities Failed: %v", err)
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// NumActivites - get total activities
func NumActivites (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

	num := numActivites(contract)
	activitiesXML := &NumActivities{Total: num}
	thisXML, _ := xml.MarshalIndent(activitiesXML, "", " ")

	fmt.Fprintf(w, "%s\n\n", thisXML)
}

// AllActivites - output activies XML
func AllActivites (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

	num := numActivites(contract)

	var i int64
	for ; i < num; i++ {
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

// ActivitesContract - get activities contract address
func ActivitesContract (c *ethclient.Client) (*contracts.IATIActivities){

	contract, err := contracts.NewIATIActivities(common.HexToAddress(string(configs.ActivitiesContractAddr)), c)
	if err != nil {
		log.Fatalf("Failed to instantiate activities contract: %v", err)
		return nil
	}
	return contract
}
