package xml

import (
    "fmt"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/activity"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
)

type activityList struct {
	XMLName   	    xml.Name `xml:"reportaid:activity-list"`
    Namespace       string   `xml:"xmlns:reportaid,attr"`
	ActivitiesRef	string   `xml:"reportaid:activities-ref,attr"`
    Ref			    []string `xml:"reportaid:activity-ref"`
}

// TotalActivities - get the total number of activities
type activityTotal struct {
	XMLName   	      xml.Name  `xml:"reportaid:activity-total"`
    Namespace         string    `xml:"xmlns:reportaid,attr"`
    ActivitiesRef     string    `xml:"reportaid:activities-ref,attr"`
	Total 		      int64     `xml:"reportaid:total"`
}

// IATIActivity - the XML for the activities file
type IATIActivity struct {
	XMLName   	    xml.Name        `xml:"iati-activity"`
    Namespace       string          `xml:"xmlns:reportaid,attr"`
	Lang 	        string          `xml:"xml:lang,attr"`
	Currency 		string          `xml:"default-currency,attr"`
	Time            string          `xml:"last-updated-datetime,attr"`
    Humanitarian    bool            `xml:"humanitarian,attr"`
	LinkedData      string          `xml:"linked-data-uri,attr"`
    Hierarchy       uint8           `xml:"hierarchy,attr"`
    Budget          uint8           `xml:"budget-not-provided,attr"`
    ReportingOrg    ReportingOrg
    Status          Status
    IATIIdentifier  string          `xml:"iati-identifier"`
    Title           string          `xml:"title>narrative"`
    Desc            Desc
    Ref             string          `xml:"reportaid:activity-ref"`
}

// Status - activity status
type Status struct {
    XMLName   xml.Name  `xml:"activity-status"`
    Code      uint8     `xml:"code,attr"`
}

// GetNumActivites - Get the total number of activities
func activityNum(contract *activity.IATIActivity, activitiesRef [32]byte) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{}, activitiesRef)
	if err != nil {
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// ActivityNum - get total activity
func ActivityNum (contracts *contracts.Contracts, activitiesRef [32]byte) ([]byte) {

    num := activityNum(contracts.ActivityContract, activitiesRef)
    ref := fmt.Sprintf("%x", activitiesRef)
    totalXML := &activityTotal{Namespace: configs.URLReportAidNamespace, ActivitiesRef: ref, Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", "")
    return thisXML
}

// ActivityID get specific acvtitie
func ActivityID (contracts *contracts.Contracts, activitiesRef [32]byte, activityRef [32]byte) (*IATIActivity, string) {

    activity, err := contracts.ActivityContract.GetActivity(&bind.CallOpts{}, activitiesRef, activityRef)
    if err != nil {
        return nil, configs.ErrorActivityID
    }

    lang := utils.GetString(activity.Lang)
    curr := utils.GetString(activity.Currency)
    time := utils.GetString(activity.LastUpdated)
    thisHumanitarian := activity.Humanitarian
    link := utils.GetString(activity.LinkedData)
    hierarchy := activity.Hierarchy
    budget := activity.BudgetNotProvided

    //fmt.Println("%x", activity.ReportingOrg.OrgRef)

    reportOrgName := OrgName(contracts, activity.ReportingOrg.OrgRef)
    if reportOrgName == "" {
        return nil, configs.ErrorOrgsName
    }
    reportingOrgRef := OrgID(contracts, activity.ReportingOrg.OrgRef)
    if reportingOrgRef == "" {
        return nil, configs.ErrorOrgsID
    }
    reportingOrgType := activity.ReportingOrg.OrgType
    reportingOrgIsSecondary := activity.ReportingOrg.IsSecondary

    status := activity.Status
    id := utils.GetString(activity.Identifier)
    title := activity.Title
    descType := uint8(1) // i think there's only one type at the moment
    desc := activity.Description

    thisActivityRef := fmt.Sprintf("%x", activityRef)

    activityXML := &IATIActivity{
                                    Namespace: configs.URLReportAidNamespace,
                                    Lang: lang,
                                    Currency: curr,
                                    Time: time,
                                    LinkedData: link,
                                    Humanitarian: thisHumanitarian,
                                    Hierarchy: hierarchy,
                                    Budget: budget,
                                    ReportingOrg: ReportingOrg{Ref: reportingOrgRef, Type: reportingOrgType, Secondary: reportingOrgIsSecondary, Narrative: reportOrgName},
                                    Status: Status{Code: status},
                                    IATIIdentifier: id,
                                    Title: title,
                                    Desc: Desc{Type: descType, Narrative: desc},
                                    Ref: thisActivityRef,
                                }

    return activityXML, ""
}


// ActivityList - list all activities
func ActivityList (contracts *contracts.Contracts, activitiesRef [32]byte) ([]byte){

    log := LogInit()
    thisActivitiesRef := fmt.Sprintf("%x", activitiesRef)
    activityIDs := &activityList{Namespace: configs.URLReportAidNamespace, ActivitiesRef: thisActivitiesRef}
    num := activityNum(contracts.ActivityContract, activitiesRef)
    var i int64
    for ; i < num; i++ {
        ref, err := contracts.ActivityContract.GetReference(&bind.CallOpts{}, activitiesRef, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorActivityList, err)
            log.Errors = append(log.Errors, thisError)
            error, _ := xml.MarshalIndent(log, "", "")
            return error
        }
        thisRef := fmt.Sprintf("%x", ref)
        activityIDs.Ref = append(activityIDs.Ref, thisRef)
    }
    thisXML, err := xml.MarshalIndent(activityIDs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorActivityList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
