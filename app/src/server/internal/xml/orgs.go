package xml

import (
    "fmt"
	"encoding/xml"
  	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/orgs"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// OrgsList - a list if all orgs
type orgsList struct {
	XMLName   	xml.Name `xml:"iati-orgs-ids"`
	ID			[]string `xml:"id"`
}

// Orgs - the XML for the orgs file
type iatiOrgs struct {
	XMLName   	xml.Name 	`xml:"iati-orgs"`
	Org 		[]*iatiorg 	`xml:"orgs"`
}

type iatiorg struct {
	Ref         string `xml:"ref,attr"`
	ID			string `xml:"iati-identifier"`
	Name   		string `xml:"name>narrative"`
}

// TotalOrgs - get the total number of orgs
type orgsTotal struct {
	XMLName   	xml.Name `xml:"iati-orgs-total"`
	Total 		int64 `xml:"orgs-total"`
}

// ReportingOrg - organisation reporting organisation
type ReportingOrg struct {
    XMLName   	xml.Name `xml:"reporting-org"`
    Ref         string `xml:"ref,attr"`
    Type        uint8 `xml:"type,attr"`
    Secondary   bool `xml:"secondary-reporter,attr"`
    Narrative   string `xml:"narrative"`
}

// OrgName - get the org name for the reference
func OrgName (contract *contracts.Contracts, orgRef [32]byte) (string) {

    name, err := contract.OrgsContract.GetOrgName(&bind.CallOpts{}, orgRef)
	if err != nil {
		return ""
	}
	return name

}

// OrgID - get the IATI ref' for the org'
func OrgID (contract *contracts.Contracts, orgRef [32]byte) (string) {

    id, err := contract.OrgsContract.GetOrgIdentifier(&bind.CallOpts{}, orgRef)
	if err != nil {
		return ""
	}
	return id

}

// GetNumOrgs - Get the total number of orgs
func orgsNum(contract *orgs.IATIOrgs) (int64) {

	num, err := contract.GetNumOrgs(&bind.CallOpts{})
	if err != nil {
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// OrgsNum - get total orgs
func OrgsNum (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := orgsNum(contracts.OrgsContract)
    totalXML := &orgsTotal{Total: num}
    thisXML, err := xml.MarshalIndent(totalXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrgsNum + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        thisXML, _ := xml.MarshalIndent(log, "", "")
        return thisXML
    }
    return thisXML
}

// OrgsID - get specific org
func OrgsID (contracts *contracts.Contracts, orgsRef [32]byte, activityRef [32]byte) ([]byte) {

    log := LogInit()
    orgs, err := contracts.OrgsContract.GetOrg(&bind.CallOpts{}, orgsRef)
    if err != nil {
        thisError := Log(configs.ErrorOrgs, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }


    thisOrgRef := fmt.Sprintf("%x", orgsRef)
    name := orgs.Name
    identifier := orgs.Identifier

	orgXML := &iatiorg {
							Ref: thisOrgRef,
							ID: identifier,
							Name: name,
						}

    orgsXML := &iatiOrgs{}
    orgsXML.Org = append(orgsXML.Org, orgXML)
    thisXML, err := xml.MarshalIndent(orgsXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrgs + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}

// OrgsList - list all orgs
func OrgsList (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := orgsNum(contracts.OrgsContract)
    var i int64
    orgsIDs := &orgsList{}
    for ; i < num; i++ {
        ref, err := contracts.OrgsContract.GetOrgReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorOrgsID, err)
            log.Errors = append(log.Errors, thisError)
            thisXML, _ := xml.MarshalIndent(log, "", "")
            return thisXML
        }
        thisRef := fmt.Sprintf("%x", ref)
        orgsIDs.ID = append(orgsIDs.ID, thisRef)
    }
    thisXML, err := xml.MarshalIndent(orgsIDs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrgsList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
