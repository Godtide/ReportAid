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
	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// ActivitiesList - a list if all activities
type activitiesList struct {
	XMLName   	xml.Name `xml:"iati-activities-ids"`
	ID			[]string `xml:"id"`
}

// Activities - the XML for the activities file
type iatiActivities struct {
	XMLName   	xml.Name `xml:"iati-activities"`
	Version 	string `xml:"version,attr"`
	Time 		string `xml:"generated-datetime,attr"`
	LinkedData  string `xml:"linked-data-default,attr"`
}

// TotalActivities - get the total number of activities
type totalActivities struct {
	XMLName   	xml.Name `xml:"iati-total-activities"`
	Total 		int64 `xml:"total"`
}

// GetNumActivites - Get the total number of activities
func numActivites(contract *activities.IATIActivities) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorNumActivities, err)
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// NumActivites - get total activities
func NumActivites (c *ethclient.Client, contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    num := numActivites(contracts.ActivitiesContract)
    totalXML := &totalActivities{Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", " ")
    w.Write(thisXML)
}

// ActivitiesID get specific acvtitie
func ActivitiesID (c *ethclient.Client, contracts *contracts.Contracts, ref [32]byte, w http.ResponseWriter, r *http.Request) {

    activities, err := contracts.ActivitiesContract.GetActivities(&bind.CallOpts{}, ref)
    if err != nil {
        log.Fatalf("%s: %v", configs.ErrorActivities, err)
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

    activitiesXML := &iatiActivities{
                                        Version: string(stringVersion),
                                        Time: string(stringTime),
                                        LinkedData: string(stringLink),
                                    }
    thisXML, _ := xml.MarshalIndent(activitiesXML, " ", " ")
    w.Write(thisXML)
	//fmt.Printf("Activities Reference: %x\n", ref)
}


// ListActivities - list all activities
func ListActivities (c *ethclient.Client, contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    num := numActivites(contracts.ActivitiesContract)
    var i int64
    activitiesIDs := &activitiesList{}
    for ; i < num; i++ {
        ref, err := contracts.ActivitiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
            log.Fatalf("%s: %v", configs.ErrorActivitiesID, err)
        }
        thisRef := fmt.Sprintf("%x", ref)
        activitiesIDs.ID = append(activitiesIDs.ID, thisRef)
    }
    thisXML, _ := xml.MarshalIndent(activitiesIDs, " ", " ")
    w.Write(thisXML)
}
