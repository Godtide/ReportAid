package xml

import (
    "fmt"
	"log"
    "net/http"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/activity"
	"github.com/ReportAid/app/src/server/internal/configs"
)

type activityList struct {
	XMLName   	    xml.Name `xml:"iati-activity-ids"`
	ActivitiesID	string `xml:"activities-id"`
    ID			    []string `xml:"activity-ids>id"`
}

// Activities - the XML for the activities file
type iatiActivity struct {
	XMLName   	    xml.Name `xml:"iati-activity"`
	Lang 	        string `xml:"xml:lang,attr"`
	Currency 		string `xml:"default-currency,attr"`
	Time            string `xml:"last-updated-datetime,attr"`
    Humanitarian    int64 `xml:"humanitarian,attr"`
	LinkedData      string `xml:"linked-data-uri,attr"`
    Budget          int64 `xml:"budget-not-provided,attr"`
}

// TotalActivities - get the total number of activities
type totalActivity struct {
	XMLName   	      xml.Name   `xml:"iati-total-activity"`
    IATIActivities    string      `xml:"iati-id"`
	Total 		      int64       `xml:"total"`
}

// GetNumActivites - Get the total number of activities
func numActivity(contract *activity.IATIActivity, activitiesRef [32]byte) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{}, activitiesRef)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorNumActivity, err)
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// NumActivity - get total activity
func NumActivity (contracts *contracts.Contracts, activitiesRef [32]byte, w http.ResponseWriter, r *http.Request) {

    num := numActivity(contracts.ActivityContract, activitiesRef)
    ref := fmt.Sprintf("%x", activitiesRef)
    totalXML := &totalActivity{IATIActivities: ref, Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", "")
    w.Write(thisXML)
}

// ActivityID get specific acvtitie
func ActivityID (contracts *contracts.Contracts, activitiesRef [32]byte, activityRef [32]byte, w http.ResponseWriter, r *http.Request) {

    activity, err := contracts.ActivityContract.GetActivity(&bind.CallOpts{}, activitiesRef, activityRef)
    if err != nil {
        log.Fatalf("%s: %v", configs.ErrorActivities, err)
    }

    /*Lang 	        string `xml:"xml:lang,attr"`
	Currency 		string `xml:"default-currency,attr"`
	Time            string `xml:"last-updated-datetime,attr"`
    Humanitarian    int64 `xml:"humanitarian,attr"`
	LinkedData      string `xml:"linked-data-uri,attr"`
    Budget          int64 `xml:"budget-not-provided,attr"`*/

    /*bool humanitarian;
    uint8 hierarchy;
    uint8 status;
    uint8 budgetNotProvided;
    ReportingOrg reportingOrg;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bytes32 linkedData;
    bytes32 identifier;
    string title;
    string description;*/

    thisLang := activity.Lang
    sliceLang := thisLang[:]
    trimmedLang := common.TrimRightZeroes(sliceLang)
    lang := hexutil.Encode(trimmedLang)
    stringLang, _ := hexutil.Decode(lang)

    thisTime := activity.LastUpdated
    sliceTime := thisTime[:]
    trimmedTime := common.TrimRightZeroes(sliceTime)
    time := hexutil.Encode(trimmedTime)
    stringTime, _ := hexutil.Decode(time)

    thisLink := activity.LinkedData
    sliceLink := thisLink[:]
    trimmedLink := common.TrimRightZeroes(sliceLink)
    link := hexutil.Encode(trimmedLink)
    stringLink, _ := hexutil.Decode(link)

    /* fmt.Printf("Version: %s\n", stringVersion)
    fmt.Printf("Generated Time: %s\n", stringTime)
    fmt.Printf("Linked Data: %s\n\n", stringLink)*/

    activityXML := &iatiActivity{
                                    Lang: string(stringLang),
                                    Time: string(stringTime),
                                    LinkedData: string(stringLink),
                                }
    thisXML, _ := xml.MarshalIndent(activityXML, " ", " ")
    w.Write(thisXML)
	//fmt.Printf("Activities Reference: %x\n", ref)
}


// ListActivity - list all activities
func ListActivity (contracts *contracts.Contracts, activitiesRef [32]byte, w http.ResponseWriter, r *http.Request) {

    thisActivitiesRef := fmt.Sprintf("%x", activitiesRef)
    activityIDs := &activityList{ActivitiesID: thisActivitiesRef}
    num := numActivity(contracts.ActivityContract, activitiesRef)
    var i int64
    for ; i < num; i++ {
        ref, err := contracts.ActivityContract.GetReference(&bind.CallOpts{}, activitiesRef, big.NewInt(i))
        if err != nil {
            log.Fatalf("%s: %v", configs.ErrorActivityID, err)
        }
        thisRef := fmt.Sprintf("%x", ref)
        activityIDs.ID = append(activityIDs.ID, thisRef)
    }
    thisXML, _ := xml.MarshalIndent(activityIDs, configs.XMLTabDouble, configs.XMLTabDouble)
    w.Write(thisXML)
}
