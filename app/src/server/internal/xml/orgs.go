package xml

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ReportAid/app/src/server/internal/contracts"
)

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
