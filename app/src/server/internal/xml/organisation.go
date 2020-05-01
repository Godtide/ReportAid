package xml

import (
    "fmt"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/organisation"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
)

type organisationList struct {
	XMLName   	    xml.Name `xml:"iati-organisation-ids"`
	OrganisationsID	string `xml:"organisations-id"`
    ID			    []string `xml:"organisation-ids>id"`
}

// IATIOrganisation - the XML for the organisation file
type IATIOrganisation struct {
	XMLName   	    xml.Name `xml:"iati-organisation"`
	Lang 	        string `xml:"xml:lang,attr"`
	Currency 		string `xml:"default-currency,attr"`
	Time            string `xml:"last-updated-datetime,attr"`
    Name           string `xml:"name>narrative"`
    ReportingOrg    ReportingOrg
}

// organisationTotal - get the total number of organisations
type organisationTotal struct {
	XMLName                xml.Name    `xml:"iati-organisation-total"`
    IATIOrganisations      string      `xml:"iati-organisations-id"`
	Total 		           int64       `xml:"organisation-total"`
}

// organisationNum - Get the total number of organisations
func organisationNum(contract *organisation.IATIOrganisation, organisationsRef [32]byte) (int64) {

	num, err := contract.GetNumOrganisations(&bind.CallOpts{}, organisationsRef)
	if err != nil {
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// OrganisationNum - get total organisation
func OrganisationNum (contracts *contracts.Contracts, organisationsRef [32]byte) ([]byte) {

    num := organisationNum(contracts.OrganisationContract, organisationsRef)
    ref := fmt.Sprintf("%x", organisationsRef)
    totalXML := &organisationTotal{IATIOrganisations: ref, Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", "")
    return thisXML
}

// OrganisationID get specific acvtitie
func OrganisationID (contracts *contracts.Contracts, organisationsRef [32]byte, organisationRef [32]byte) (*IATIOrganisation, string) {

    organisation, err := contracts.OrganisationContract.GetOrganisation(&bind.CallOpts{}, organisationsRef, organisationRef)
    if err != nil {
        return nil, configs.ErrorOrganisationID
    }

    lang := utils.GetString(organisation.Lang)
    curr := utils.GetString(organisation.Currency)

    orgName := OrgName(contracts, organisation.OrgRef)
    if orgName == "" {
        return nil, configs.ErrorOrgsName
    }

    reportOrgName := OrgName(contracts, organisation.ReportingOrg.OrgRef)
    if reportOrgName == "" {
        return nil, configs.ErrorOrgsName
    }
    reportingOrgRef := OrgID(contracts, organisation.ReportingOrg.OrgRef)
    if reportingOrgRef == "" {
        return nil, configs.ErrorOrgsID
    }
    reportingOrgType := organisation.ReportingOrg.OrgType
    reportingOrgIsSecondary := organisation.ReportingOrg.IsSecondary

    organisationXML := &IATIOrganisation{
                                    Lang: lang,
                                    Currency: curr,
                                    Name: orgName,
                                    ReportingOrg: ReportingOrg{Ref: reportingOrgRef, Type: reportingOrgType, Secondary: reportingOrgIsSecondary, Narrative: reportOrgName},
                                }
    return organisationXML, ""
}


// OrganisationList - list all organisations
func OrganisationList (contracts *contracts.Contracts, organisationsRef [32]byte) ([]byte){

    log := LogInit()
    thisOrganisationsRef := fmt.Sprintf("%x", organisationsRef)
    organisationIDs := &organisationList{OrganisationsID: thisOrganisationsRef}
    num := organisationNum(contracts.OrganisationContract, organisationsRef)
    var i int64
    for ; i < num; i++ {
        ref, err := contracts.OrganisationContract.GetOrganisationReference(&bind.CallOpts{}, organisationsRef, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorOrganisationList, err)
            log.Errors = append(log.Errors, thisError)
            error, _ := xml.MarshalIndent(log, "", "")
            return error
        }
        thisRef := fmt.Sprintf("%x", ref)
        organisationIDs.ID = append(organisationIDs.ID, thisRef)
    }
    thisXML, err := xml.MarshalIndent(organisationIDs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
