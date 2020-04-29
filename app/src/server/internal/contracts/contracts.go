package contracts

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// Contracts - keeps pointers to all ethereum contracts
type Contracts struct {
	ActivitiesContract    	*activities.IATIActivities
}

// ActivitesContract - get activities contract address
func ActivitesContract (c *ethclient.Client) (*activities.IATIActivities) {

	contract, err := activities.NewIATIActivities(common.HexToAddress(string(configs.AddrActivitiesContract)), c)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorActivitiesContract, err)
		return nil
	}
	return contract
}

//  AllContracts - get instances of aqll the ethereum contracts
func AllContracts (c *ethclient.Client) (*Contracts) {

	var contracts = new(Contracts)
	var activitiesContract = ActivitesContract(c)
    contracts.ActivitiesContract = activitiesContract
	return contracts
}
