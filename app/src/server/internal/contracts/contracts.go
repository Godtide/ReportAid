package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/contracts/activity"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// Contracts - keeps pointers to all ethereum contracts
type Contracts struct {
	ActivitiesContract    	*activities.IATIActivities
	ActivityContract    	*activity.IATIActivity
}

// ActivitesContract - get activities contract
func ActivitesContract (c *ethclient.Client) (*activities.IATIActivities) {

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

//  AllContracts - get instances of aqll the ethereum contracts
func AllContracts (c *ethclient.Client) (*Contracts) {

	var contracts = new(Contracts)
	var activitiesContract = ActivitesContract(c)
	var activityContract = ActivityContract(c)

    contracts.ActivitiesContract = activitiesContract
    contracts.ActivityContract = activityContract
	
	return contracts
}
