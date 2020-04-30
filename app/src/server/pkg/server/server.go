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

func homePage(w http.ResponseWriter){

    xml.Header(w)
    content := xml.Home()
    w.Write(content)
}

func listActivities (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.ListActivities(contracts)
    w.Write(content)
}

func singleActivities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	activityRef := params[configs.URLParamActivityRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	convertedActivityRef := common.HexToHash(activityRef)
	content := xml.ActivitiesID(contracts, convertedActivitiesRef, convertedActivityRef)
    w.Write(content)
}

func numActivities (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.NumActivites(contracts)
    w.Write(content)
}

func listActivity (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	content := xml.ListActivity(contracts, convertedActivitiesRef)
    w.Write(content)
}

func numActivity (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
    params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
    content := xml.NumActivity(contracts, convertedActivitiesRef)
    w.Write(content)
}

func handleRequests(contracts *contracts.Contracts) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w)
	})

    // Activities
	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        singleActivities(contracts, w, r)
	})

	router.HandleFunc(configs.URLListActivities, func(w http.ResponseWriter, r *http.Request) {
    	listActivities(contracts, w)
	})

	router.HandleFunc(configs.URLTotalActivities, func(w http.ResponseWriter, r *http.Request) {
    	numActivities(contracts, w)
	})

    router.HandleFunc(configs.URLListActivity, func(w http.ResponseWriter, r *http.Request) {
    	listActivity(contracts, w, r)
	})

    router.HandleFunc(configs.URLTotalActivity, func(w http.ResponseWriter, r *http.Request) {
    	numActivity(contracts, w, r)
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
