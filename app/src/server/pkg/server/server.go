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

func activities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {
    xml.AllActivites(c, contract, w, r)
}

func singleActivitie (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	value := params["ref"]
	convertedValue := common.HexToHash(value)
	xml.AnActivitie(c, contract, convertedValue, w, r)
}


func numActivities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {
    xml.NumActivites(c, contract, w, r)
}

func handleRequests(c *ethclient.Client) {

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	homePage(w, r)
	})

	contract := xml.ActivitesContract(c)
	router.HandleFunc("/activities/{ref}", func(w http.ResponseWriter, r *http.Request) {
    	singleActivitie(c, contract, w, r)
	})

	router.HandleFunc("/activities", func(w http.ResponseWriter, r *http.Request) {
    	activities(c, contract, w, r)
	})

	router.HandleFunc("/num-activities", func(w http.ResponseWriter, r *http.Request) {
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
