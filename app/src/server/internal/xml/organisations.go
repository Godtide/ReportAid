package xml

import (
    "fmt"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/contracts/organisations"
	"github.com/ReportAid/app/src/server/internal/configs"
	"github.com/ReportAid/app/src/server/internal/utils"
)

// OrganisationsList - a list if all organisations
type organisationsList struct {
	XMLName   	xml.Name `xml:"iati-organisations-ids"`
	ID			[]string `xml:"id"`
}

// Organisations - the XML for the organisations file
type iatiOrganisations struct {
	XMLName   	xml.Name `xml:"iati-organisations"`
	Version 	string `xml:"version,attr"`
	Time 		string `xml:"generated-datetime,attr"`
    Organisation    []*IATIOrganisation
}

// TotalOrganisations - get the total number of organisations
type organisationsTotal struct {
	XMLName   	xml.Name `xml:"iati-organisations-total"`
	Total 		int64 `xml:"organisations-total"`
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
    totalXML := &organisationsTotal{Total: num}
    thisXML, err := xml.MarshalIndent(totalXML, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationsNum + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        thisXML, _ := xml.MarshalIndent(log, "", "")
        return thisXML
    }
    return thisXML
}

// OrganisationsID - get specific activitie
func OrganisationsID (contracts *contracts.Contracts, organisationsRef [32]byte, organisationRef [32]byte) ([]byte) {

    log := LogInit()
    organisations, err := contracts.OrganisationsContract.GetOrganisations(&bind.CallOpts{}, organisationsRef)
    if err != nil {
        thisError := Log(configs.ErrorOrganisations, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    version := utils.GetString(organisations.Version)
    time := utils.GetString(organisations.GeneratedTime)
    var organisation, errString = OrganisationID(contracts, organisationsRef, organisationRef)
    if organisation == nil {
        thisError := Log(configs.ErrorOrganisation + " - " + errString, nil)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }

    organisationsXML := &iatiOrganisations{
                                        Version: string(version),
                                        Time: string(time),
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
	//fmt.Printf("Organisations Reference: %x\n", ref)
}


// OrganisationsList - list all organisations
func OrganisationsList (contracts *contracts.Contracts) ([]byte) {

    log := LogInit()
    num := organisationsNum(contracts.OrganisationsContract)
    var i int64
    organisationsIDs := &organisationsList{}
    for ; i < num; i++ {
        ref, err := contracts.OrganisationsContract.GetOrganisationsReference(&bind.CallOpts{}, big.NewInt(i))
        if err != nil {
            thisError := Log(configs.ErrorOrganisationsID, err)
            log.Errors = append(log.Errors, thisError)
            thisXML, _ := xml.MarshalIndent(log, "", "")
            return thisXML
        }
        thisRef := fmt.Sprintf("%x", ref)
        organisationsIDs.ID = append(organisationsIDs.ID, thisRef)
    }
    thisXML, err := xml.MarshalIndent(organisationsIDs, "", "")
    if err != nil {
        thisError := Log(configs.ErrorOrganisationsList + " - " + configs.ErrorUnMarshall, err)
        log.Errors = append(log.Errors, thisError)
        error, _ := xml.MarshalIndent(log, "", "")
        return error
    }
    return thisXML
}
