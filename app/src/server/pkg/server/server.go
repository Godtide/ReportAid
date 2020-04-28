package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ReportAid/app/src/server/internal/xml"
	"github.com/ReportAid/app/src/server/internal/contracts/activities"
	"github.com/ReportAid/app/src/server/internal/configs"
)

func homePage(w http.ResponseWriter, r *http.Request){

    xml.First(w, r)
    xml.Home(w, r)
}

func listActivities (c *ethclient.Client, contract *activities.IATIActivities, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.ListActivities(c, contract, w, r)
}

func singleActivities (c *ethclient.Client, contract *activities.IATIActivities, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.ActivitiesID(c, contract, convertedValue, w, r)
}


func numActivities (c *ethclient.Client, contract *activities.IATIActivities, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
    xml.NumActivites(c, contract, w, r)
}

func handleRequests(c *ethclient.Client) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w, r)
	})

	contract := xml.ActivitesContract(c)
	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        singleActivities(c, contract, w, r)
	})
    //configs.URLActivities
    // "/activities/{ref}"

	router.HandleFunc(configs.URLListActivities, func(w http.ResponseWriter, r *http.Request) {
    	listActivities(c, contract, w, r)
	})

	router.HandleFunc(configs.URLTotalActivities, func(w http.ResponseWriter, r *http.Request) {
    	numActivities(c, contract, w, r)
	})

    log.Fatal(http.ListenAndServe(configs.ServerPort, router))
}

// Start - fire up the server
func Start() {

	conn, err := ethclient.Dial(string(configs.AddrRinkeby))
	if err != nil {
		log.Fatalf("%s: %v", configs.ErrorRinkeby, err)
	}
    handleRequests(conn)
}
