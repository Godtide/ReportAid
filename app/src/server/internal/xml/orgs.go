package xml

import (
    "log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// OrgName - get the org name for the reference
func OrgName (contract *contracts.Contracts, orgRef [32]byte) (string) {

    name, err := contract.OrgsContract.GetOrgName(&bind.CallOpts{}, orgRef)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorOrgsName, err)
		return ""
	}
	return name

}

// OrgID - get the IATI ref' for the org'
func OrgID (contract *contracts.Contracts, orgRef [32]byte) (string) {

    id, err := contract.OrgsContract.GetOrgIdentifier(&bind.CallOpts{}, orgRef)
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorOrgsName, err)
		return ""
	}
	return id

}
