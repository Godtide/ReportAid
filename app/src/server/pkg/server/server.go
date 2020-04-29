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

func listActivities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.ListActivities(contracts, w, r)
}

func singleActivities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.ActivitiesID(contracts, convertedValue, w, r)
}

func numActivities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.NumActivites(contracts, w, r)
}

func singleActivity (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	activitiesRef := params["activitiesRef"]
	activityRef := params["activityRef"]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	convertedActivityRef := common.HexToHash(activityRef)
	xml.ActivityID(contracts, convertedActivitiesRef, convertedActivityRef, w, r)
}

func listActivity (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.ListActivity(contracts, convertedValue, w, r)
}

func handleRequests(contracts *contracts.Contracts) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w, r)
	})

    // Activities
	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        singleActivities(contracts, w, r)
	})

	router.HandleFunc(configs.URLListActivities, func(w http.ResponseWriter, r *http.Request) {
    	listActivities(contracts, w, r)
	})

	router.HandleFunc(configs.URLTotalActivities, func(w http.ResponseWriter, r *http.Request) {
    	numActivities(contracts, w, r)
	})

    // Activity
    router.HandleFunc(configs.URLActivity, func(w http.ResponseWriter, r *http.Request) {
    	singleActivity(contracts, w, r)
	})

    router.HandleFunc(configs.URLListActivity, func(w http.ResponseWriter, r *http.Request) {
    	listActivity(contracts, w, r)
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
    handleRequests(contracts)
}
