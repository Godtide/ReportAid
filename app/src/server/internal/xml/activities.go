package xml

import (
    "fmt"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
)

// ActivitiesList - a list if all activities
type activitiesList struct {
	XMLName   	xml.Name `xml:"reportaid:activities-list"`
    Namespace   string   `xml:"xmlns:reportaid,attr"`
	Ref			[]string `xml:"reportaid:activities-ref"`
}

// TotalActivities - get the total number of activities
type activitiesTotal struct {
	XMLName   	xml.Name `xml:"reportaid:activities-total"`
    Namespace   string   `xml:"xmlns:reportaid,attr"`
	Total 		int64    `xml:"reportaid:total"`
}

// Activities - the IATI XML for the activities file
type iatiActivities struct {
	XMLName   	xml.Name    `xml:"iati-activities"`
    Namespace   string      `xml:"xmlns:reportaid,attr"`
	Version 	string      `xml:"version,attr"`
	Time 		string      `xml:"generated-datetime,attr"`
	LinkedData  string      `xml:"linked-data-default,attr"`
    Activity    []*IATIActivity
    Ref         string      `xml:"reportaid:activities-ref"`
}

// GetNumActivities - Get the total number of activities
func activitiesNum(contract *activities.IATIActivities) (int64) {

    num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// ActivitiesNum - get total activities
func ActivitiesNum (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := activitiesNum(contracts.ActivitiesContract)
    totalXML := &activitiesTotal{Namespace: configs.URLReportAidNamespace, Total: num}
    thisXML, err := xml.MarshalIndent(totalXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorActivitiesNum + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        thisXML, _ := xml.MarshalIndent(log, "", "")
        return thisXML
    }
    return thisXML
}

func activitiesHeader (contracts *contracts.Contracts, activitiesRef [32]byte) (*iatiActivities, string) {

    thisActivitiesRef := fmt.Sprintf("%x", activitiesRef)

    activities, err := contracts.ActivitiesContract.GetActivities(&bind.CallOpts{}, activitiesRef)
    if err != nil {
        return nil, configs.ErrorActivities
    }

    version := utils.GetString(activities.Version)
    time := utils.GetString(activities.GeneratedTime)
    link := utils.GetString(activities.LinkedData)

    activitiesXML := &iatiActivities{
                                        Namespace: configs.URLReportAidNamespace,
                                        Version: string(version),
                                        Time: string(time),
                                        LinkedData: string(link),
                                        Ref: thisActivitiesRef,
                                    }
    return activitiesXML, ""
}

// Activities - get specific activities
func Activities (contracts *contracts.Contracts, activitiesRef [32]byte, activityRef [32]byte) ([]byte) {

    log := LogInit()

    activitiesXML, errString := activitiesHeader(contracts, activitiesRef)
    if activitiesXML == nil {
        thisError := Log(configs.ErrorActivities + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    var activity, activityErrString = Activity(contracts, activitiesRef, activityRef)
    if activity == nil {
        thisError := Log(configs.ErrorActivity + " - " + activityErrString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    activitiesXML.Activity = append(activitiesXML.Activity, activity)

    thisXML, err := xml.MarshalIndent(activitiesXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorActivity + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}

// ActivitiesAll - get specific activities
func ActivitiesAll (contracts *contracts.Contracts, activitiesRef [32]byte) ([]byte) {

    log := LogInit()

    activitiesXML, errString := activitiesHeader(contracts, activitiesRef)
    if activitiesXML == nil {
        thisError := Log(configs.ErrorActivities + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    activityList, errString := ActivityList(contracts, activitiesRef)
    if activityList == nil {
        thisError := Log(configs.ErrorActivities + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    for i := 0; i < len(activityList); i++ {

    	convertedActivityRef := common.HexToHash(activityList[i])
        var activity, errString = Activity(contracts, activitiesRef, convertedActivityRef)
        if activity == nil {
            thisError := Log(configs.ErrorActivity + " - " + errString, nil)
            log.Errors = append(log.Errors, thisError)
            error, _ := xml.MarshalIndent(log, "", "")
            return error
        }
        activitiesXML.Activity = append(activitiesXML.Activity, activity)
    }

    thisXML, err := xml.MarshalIndent(activitiesXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorActivity + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
	//fmt.Printf("Activities Reference: %x\n", ref)
}

// ActivitiesList - activities list
func ActivitiesList (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    activitiesRefs := &activitiesList{Namespace: configs.URLReportAidNamespace}

    num := activitiesNum(contracts.ActivitiesContract)
    var i int64
    for ; i < num; i++ {
        ref, err := contracts.ActivitiesContract.GetReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorActivitiesRef, err)
            log.Errors = append(log.Errors, thisError)
            thisXML, _ := xml.MarshalIndent(log, "", "")
            return thisXML
        }
        thisRef := fmt.Sprintf("%x", ref)
        activitiesRefs.Ref = append(activitiesRefs.Ref, thisRef)
    }
    thisXML, err := xml.MarshalIndent(activitiesRefs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorActivitiesList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
