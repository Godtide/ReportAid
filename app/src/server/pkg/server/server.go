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

func activitiesList (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.ActivitiesList(contracts)
    w.Write(content)
}

func activities (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	activityRef := params[configs.URLParamActivityRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	convertedActivityRef := common.HexToHash(activityRef)
	content := xml.ActivitiesID(contracts, convertedActivitiesRef, convertedActivityRef)
    w.Write(content)
}

func activitiesNum (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.ActivitesNum(contracts)
    w.Write(content)
}

func activityList (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	content := xml.ActivityList(contracts, convertedActivitiesRef)
    w.Write(content)
}

func activityNum (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
    params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
    content := xml.ActivityNum(contracts, convertedActivitiesRef)
    w.Write(content)
}

func handleRequests(contracts *contracts.Contracts) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w)
	})

    // Activities
	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        activities(contracts, w, r)
	})

	router.HandleFunc(configs.URLListActivities, func(w http.ResponseWriter, r *http.Request) {
    	activitiesList(contracts, w)
	})

	router.HandleFunc(configs.URLTotalActivities, func(w http.ResponseWriter, r *http.Request) {
    	activitiesNum(contracts, w)
	})

    router.HandleFunc(configs.URLListActivity, func(w http.ResponseWriter, r *http.Request) {
    	activityList(contracts, w, r)
	})

    router.HandleFunc(configs.URLTotalActivity, func(w http.ResponseWriter, r *http.Request) {
    	activityNum(contracts, w, r)
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
