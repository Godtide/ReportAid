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
	XMLName   	        xml.Name    `xml:"reportaid:organisation-list"`
    Namespace           string      `xml:"xmlns:reportaid,attr"`
	OrganisationsRef	string      `xml:"reportaid:organisations-ref,attr"`
    Ref			        []string     `xml:"reportaid:organisation-ref"`
}

// organisationTotal - get the total number of organisations
type organisationTotal struct {
	XMLName                xml.Name   `xml:"reportaid:organisation-total"`
    Namespace              string     `xml:"xmlns:reportaid,attr"`
    OrganisationsRef      string      `xml:"reportaid:organisations-ref,attr"`
	Total 		           int64      `xml:"reportaid:total"`
}

// IATIOrganisation - the XML for the organisation file
type IATIOrganisation struct {
	XMLName   	    xml.Name         `xml:"iati-organisation"`
    Namespace       string          `xml:"xmlns:reportaid,attr"`
	Lang 	        string           `xml:"xml:lang,attr"`
	Currency 		string          `xml:"default-currency,attr"`
	Time            string         `xml:"last-updated-datetime,attr"`
    Name            string         `xml:"name>narrative"`
    ReportingOrg    ReportingOrg
    Ref			    []string       `xml:"reportaid:organisation-ref"`
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
    totalXML := &organisationTotal{Namespace: configs.URLReportAidNamespace, OrganisationsRef: ref, Total: num}
    thisXML, _ := xml.MarshalIndent(totalXML, "", "")
    return thisXML
}

// Organisation get specific organisation
func Organisation (contracts *contracts.Contracts, organisationsRef [32]byte, organisationRef [32]byte) (*IATIOrganisation, string) {

    organisation, err := contracts.OrganisationContract.GetOrganisation(&bind.CallOpts{}, organisationsRef, organisationRef)
    if err != nil {
        return nil, configs.ErrorOrganisationRef
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
        return nil, configs.ErrorOrgsRef
    }
    reportingOrgType := organisation.ReportingOrg.OrgType
    reportingOrgIsSecondary := organisation.ReportingOrg.IsSecondary

    organisationXML := &IATIOrganisation{
                                            Namespace: configs.URLReportAidNamespace,
                                            Lang: lang,
                                            Currency: curr,
                                            Name: orgName,
                                            ReportingOrg: ReportingOrg{
                                                                        Ref: reportingOrgRef,
                                                                        Type: reportingOrgType,
                                                                        Secondary: reportingOrgIsSecondary,
                                                                        Narrative: reportOrgName,
                                                                    },
                                        }
    return organisationXML, ""
}

// OrganisationList - return array of all organisation
func OrganisationList (contracts *contracts.Contracts, organisationsRef [32]byte) ([]string, string){

    var organisationRefs []string
    num := organisationNum(contracts.OrganisationContract, organisationsRef)
    var i int64
    for ; i < num; i++ {
        ref, err := contracts.OrganisationContract.GetOrganisationReference(&bind.CallOpts{}, organisationsRef, big.NewInt(i))
        if err != nil {
            return nil, configs.ErrorOrganisationList
        }
        thisRef := fmt.Sprintf("%x", ref)
        organisationRefs = append(organisationRefs, thisRef)
    }
    return organisationRefs, ""
}

// OrganisationListXML - return XML of all organisation
func OrganisationListXML (contracts *contracts.Contracts, organisationsRef [32]byte) ([]byte){

    log := LogInit()
    thisOrganisationsRef := fmt.Sprintf("%x", organisationsRef)
    organisationRefs, errString := OrganisationList(contracts, organisationsRef)
    if organisationRefs == nil {
        thisError := Log(configs.ErrorOrganisationList + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    organisationRefsXML := &organisationList{Namespace: configs.URLReportAidNamespace, OrganisationsRef: thisOrganisationsRef}
    organisationRefsXML.Ref = organisationRefs
    thisXML, err := xml.MarshalIndent(organisationRefsXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
