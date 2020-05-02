package xml

import (
    "fmt"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/organisations"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
)

// OrganisationsList - a list if all organisations
type organisationsList struct {
	XMLName   	xml.Name `xml:"reportaid:organisations-list"`
    Namespace   string   `xml:"xmlns:reportaid,attr"`
	Ref			[]string `xml:"reportaid:organisations-ref"`
}

// TotalOrganisations - get the total number of organisations
type organisationsTotal struct {
	XMLName   	xml.Name   `xml:"reportaid:organisations-total"`
    Namespace   string     `xml:"xmlns:reportaid,attr"`
	Total 		int64      `xml:"reportaid:total"`
}

// Organisations - the XML for the organisations file
type iatiOrganisations struct {
	XMLName   	    xml.Name                `xml:"iati-organisations"`
    Namespace       string                  `xml:"xmlns:reportaid,attr"`
	Version 	    string                  `xml:"version,attr"`
	Time 		    string                  `xml:"generated-datetime,attr"`
    Organisation    []*IATIOrganisation
    Ref             string                  `xml:"reportaid:organisations-ref"`
}

// GetNumOrganisations - Get the total number of organisations
func organisationsNum(contract *organisations.IATIOrganisations) (int64) {

	num, err := contract.GetNumOrganisations(&bind.CallOpts{})
	if err != nil {
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// OrganisationsNum - get total organisations
func OrganisationsNum (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := organisationsNum(contracts.OrganisationsContract)
    totalXML := &organisationsTotal{Namespace: configs.URLReportAidNamespace, Total: num}
    thisXML, err := xml.MarshalIndent(totalXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationsNum + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        thisXML, _ := xml.MarshalIndent(log, "", "")
        return thisXML
    }
    return thisXML
}

func organisationsHeader (contracts *contracts.Contracts, organisationsRef [32]byte) (*iatiOrganisations, string) {

    thisOrganisationsRef := fmt.Sprintf("%x", organisationsRef)

    organisations, err := contracts.OrganisationsContract.GetOrganisations(&bind.CallOpts{}, organisationsRef)
    if err != nil {
        return nil, configs.ErrorOrganisations
    }

    version := utils.GetString(organisations.Version)
    time := utils.GetString(organisations.GeneratedTime)

    organisationsXML := &iatiOrganisations{
                                            Namespace: configs.URLReportAidNamespace,
                                            Version: string(version),
                                            Time: string(time),
                                            Ref: thisOrganisationsRef,
                                    }
    return organisationsXML, ""
}

// Organisations - get specific activitie
func Organisations (contracts *contracts.Contracts, organisationsRef [32]byte, organisationRef [32]byte) ([]byte) {

    log := LogInit()

    organisationsXML, errString := organisationsHeader(contracts, organisationsRef)
    if organisationsXML == nil {
        thisError := Log(configs.ErrorOrganisations + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    var organisation, organisationErrString = Organisation(contracts, organisationsRef, organisationRef)
    if organisation == nil {
        thisError := Log(configs.ErrorOrganisation + " - " + organisationErrString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    organisationsXML.Organisation = append(organisationsXML.Organisation, organisation)

    thisXML, err := xml.MarshalIndent(organisationsXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisation + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}

// OrganisationsAll - get specific organisations
func OrganisationsAll (contracts *contracts.Contracts, organisationsRef [32]byte) ([]byte) {

    log := LogInit()

    organisationsXML, errString := organisationsHeader(contracts, organisationsRef)
    if organisationsXML == nil {
        thisError := Log(configs.ErrorOrganisations + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    organisationList, errString := OrganisationList(contracts, organisationsRef)
    if organisationList == nil {
        thisError := Log(configs.ErrorOrganisations + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    for i := 0; i < len(organisationList); i++ {

    	convertedOrganisationRef := common.HexToHash(organisationList[i])
        var organisation, errString = Organisation(contracts, organisationsRef, convertedOrganisationRef)
        if organisation == nil {
            thisError := Log(configs.ErrorOrganisation + " - " + errString, nil)
            log.Errors = append(log.Errors, thisError)
            error, _ := xml.MarshalIndent(log, "", "")
            return error
        }
        organisationsXML.Organisation = append(organisationsXML.Organisation, organisation)
    }

    thisXML, err := xml.MarshalIndent(organisationsXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisation + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
	//fmt.Printf("Organisations Reference: %x\n", ref)
}


// OrganisationsList - list all organisations
func OrganisationsList (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := organisationsNum(contracts.OrganisationsContract)
    var i int64
    organisationsRefs := &organisationsList{Namespace: configs.URLReportAidNamespace}
    for ; i < num; i++ {
        ref, err := contracts.OrganisationsContract.GetOrganisationsReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorOrganisationsRef, err)
            log.Errors = append(log.Errors, thisError)
            thisXML, _ := xml.MarshalIndent(log, "", "")
            return thisXML
        }
        thisRef := fmt.Sprintf("%x", ref)
        organisationsRefs.Ref = append(organisationsRefs.Ref, thisRef)
    }
    thisXML, err := xml.MarshalIndent(organisationsRefs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationsList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
