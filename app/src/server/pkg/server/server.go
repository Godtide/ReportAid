package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ReportAid/app/src/server/internal/xml"
	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/configs"
)

func homePage(w http.ResponseWriter, r *http.Request){

    xml.First(w, r)
    xml.Home(w, r)
}

func listActivities (c *ethclient.Client, contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.ListActivities(c, contracts, w, r)
}

func singleActivities (c *ethclient.Client, contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.ActivitiesID(c, contracts, convertedValue, w, r)
}


func numActivities (c *ethclient.Client, contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.NumActivites(c, contracts, w, r)
}

func handleRequests(c *ethclient.Client, contracts *contracts.Contracts) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w, r)
	})

	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        singleActivities(c, contracts, w, r)
	})
    //configs.URLActivities
    // "/activities/{ref}"

	router.HandleFunc(configs.URLListActivities, func(w http.ResponseWriter, r *http.Request) {
    	listActivities(c, contracts, w, r)
	})

	router.HandleFunc(configs.URLTotalActivities, func(w http.ResponseWriter, r *http.Request) {
    	numActivities(c, contracts, w, r)
	})

    log.Fatal(http.ListenAndServe(configs.ServerPort, router))
}

// Start - fire up the server
func Start() {

	conn, err := ethclient.Dial(string(configs.AddrRinkeby))
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorRinkeby, err)
	}
    contracts := contracts.AllContracts(conn)
    handleRequests(conn, contracts)
}
