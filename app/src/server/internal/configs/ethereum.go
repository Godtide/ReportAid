package configs

// EtherConfig - ethereum parameters
type EtherConfig string

const (
    //RinkebyAddr - Rinkeby address + infura key
    RinkebyAddr EtherConfig                = "https://rinkeby.infura.io/v3/f4851ba4a9e546e7a1f16bf8e14f0a92"
    //ActivitiesContractAddr - activities contract address + infura key
    ActivitiesContractAddr EtherConfig     = "0x1b61bE4D8f631afE6F24F595Df73D5275050B231"
)
