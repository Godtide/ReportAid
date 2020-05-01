package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/contracts/activity"
	"github.com/ReportAid/app/src/server/internal/contracts/organisations"
	"github.com/ReportAid/app/src/server/internal/contracts/organisation"
	"github.com/ReportAid/app/src/server/internal/contracts/orgs"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// Contracts - keeps pointers to all ethereum contracts
type Contracts struct {

	ActivitiesContract    	*activities.IATIActivities
	ActivityContract    	*activity.IATIActivity
	OrganisationsContract    	*organisations.IATIOrganisations
	OrganisationContract    	*organisation.IATIOrganisation
	OrgsContract			*orgs.IATIOrgs

}

// ActivitiesContract - get activities contract
func ActivitiesContract (c *ethclient.Client) (*activities.IATIActivities) {

	contract, err := activities.NewIATIActivities(common.HexToAddress(string(configs.AddrActivitiesContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorActivitiesContract, err)
		return nil
	}
	return contract
}

// ActivityContract - get activity contract
func ActivityContract (c *ethclient.Client) (*activity.IATIActivity) {

	contract, err := activity.NewIATIActivity(common.HexToAddress(string(configs.AddrActivityContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorActivityContract, err)
		return nil
	}
	return contract
}

// OrganisationsContract - get organisations contract
func OrganisationsContract (c *ethclient.Client) (*organisations.IATIOrganisations) {

	contract, err := organisations.NewIATIOrganisations(common.HexToAddress(string(configs.AddrOrganisationsContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorOrganisationsContract, err)
		return nil
	}
	return contract
}

// OrganisationContract - get organisation contract
func OrganisationContract (c *ethclient.Client) (*organisation.IATIOrganisation) {

	contract, err := organisation.NewIATIOrganisation(common.HexToAddress(string(configs.AddrOrganisationContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorOrganisationContract, err)
		return nil
	}
	return contract
}

// OrgsContract (c *ethclient.Client) (*activity.IATIActivity) {
func OrgsContract (c *ethclient.Client) (*orgs.IATIOrgs) {
	contract, err := orgs.NewIATIOrgs(common.HexToAddress(string(configs.AddrOrgsContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorOrgsContract, err)
		return nil
	}
	return contract
}

//  AllContracts - get instances of aqll the ethereum contracts
func AllContracts (c *ethclient.Client) (*Contracts) {

	var contracts = new(Contracts)
	var activitiesContract = ActivitiesContract(c)
	var activityContract = ActivityContract(c)
	var organisationsContract = OrganisationsContract(c)
	var organisationContract = OrganisationContract(c)
	var orgsContract = OrgsContract(c)

    contracts.ActivitiesContract = activitiesContract
    contracts.ActivityContract = activityContract
	contracts.OrganisationsContract = organisationsContract
    contracts.OrganisationContract = organisationContract
    contracts.OrgsContract = orgsContract

	return contracts
}
