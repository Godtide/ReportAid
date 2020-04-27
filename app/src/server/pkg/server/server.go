package server

import (
    "fmt"
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
    //fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Fprintf(w, "ReportAid home")
}

func listActivities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {
    xml.First(w, r)
    xml.ListActivities(c, contract, w, r)
}

func singleActivities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

    xml.First(w, r)
	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.ActivitiesID(c, contract, convertedValue, w, r)
}


func numActivities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {
    xml.First(w, r)
    xml.NumActivites(c, contract, w, r)
}

func handleRequests(c *ethclient.Client) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	homePage(w, r)
	})

	contract := xml.ActivitesContract(c)
	router.HandleFunc("/activities/{ref}", func(w http.ResponseWriter, r *http.Request) {
        singleActivities(c, contract, w, r)
	})

	router.HandleFunc("/list-activities", func(w http.ResponseWriter, r *http.Request) {
    	listActivities(c, contract, w, r)
	})

	router.HandleFunc("/total-activities", func(w http.ResponseWriter, r *http.Request) {
    	numActivities(c, contract, w, r)
	})

    log.Fatal(http.ListenAndServe(":8080", router))
}

// Start - fire up the server
func Start() {

	conn, err := ethclient.Dial(string(configs.RinkebyAddr))
	if err != nil {
		log.Fatalf("Failed to connect to Rinkeby: %v", err)
	}
    handleRequests(conn)
}
