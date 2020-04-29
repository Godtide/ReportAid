package xml

import (
    "fmt"
	"log"
    "net/http"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
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
    Activity    IATIActivity
}

// TotalActivities - get the total number of activities
type totalActivities struct {
	XMLName   	xml.Name `xml:"iati-total-activities"`
	Total 		int64 `xml:"total"`
}

// GetNumActivites - Get the total number of activities
func numActivities(contract *activities.IATIActivities) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorActivitiesNum, err)
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// NumActivites - get total activities
func NumActivites (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    num := numActivities(contracts.ActivitiesContract)
    totalXML := &totalActivities{Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", " ")
    w.Write(thisXML)
}

// ActivitiesID - get specific activitie
func ActivitiesID (contracts *contracts.Contracts, activitiesRef [32]byte, activityRef [32]byte, w http.ResponseWriter, r *http.Request) {

    activities, err := contracts.ActivitiesContract.GetActivities(&bind.CallOpts{}, activitiesRef)
    if err != nil {
        log.Fatalf("%s: %v", configs.ErrorActivities, err)
    }

    version := utils.GetString(activities.Version)
    time := utils.GetString(activities.GeneratedTime)
    link := utils.GetString(activities.LinkedData)
    activity := ActivityID(contracts, activitiesRef, activityRef, w, r)

    activitiesXML := &iatiActivities{
                                        Version: string(version),
                                        Time: string(time),
                                        LinkedData: string(link),
                                        Activity: *activity,
                                    }
    thisXML, _ := xml.MarshalIndent(activitiesXML, " ", " ")
    w.Write(thisXML)
	//fmt.Printf("Activities Reference: %x\n", ref)
}


// ListActivities - list all activities
func ListActivities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    num := numActivities(contracts.ActivitiesContract)
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
