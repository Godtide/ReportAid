package server

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"

	"github.com/ReportAid/app/src/server/internal/xml"
)

func homePage(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Fprintf(w, "ReportAid home")
}

func activities(w http.ResponseWriter, r *http.Request){
    xml.GetActivites(w,r)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/activities", activities)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// Start - fire up the server
func Start() {
    handleRequests()
}
