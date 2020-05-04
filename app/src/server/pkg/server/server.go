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

func notFound(w http.ResponseWriter, r *http.Request) {
    xml.Header(w)
    content := xml.NotFound()
    w.Write(content)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    xml.Header(w)
    content := xml.PostDisallowed()
    w.Write(content)
}

func home(w http.ResponseWriter){

    xml.Header(w)
    content := xml.Home()
    w.Write(content)
}

func help(w http.ResponseWriter){

    xml.Header(w)
    content := xml.Help()
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
	content := xml.Activities(contracts, convertedActivitiesRef, convertedActivityRef)
    w.Write(content)
}

func activitiesAll (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	content := xml.ActivitiesAll(contracts, convertedActivitiesRef)
    w.Write(content)
}

func activitiesNum (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.ActivitiesNum(contracts)
    w.Write(content)
}

func activityList (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamActivitiesRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	content := xml.ActivityListXML(contracts, convertedActivitiesRef)
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

func organisationsList (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.OrganisationsList(contracts)
    w.Write(content)
}

func organisations (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    organisationsRef := params[configs.URLParamOrganisationsRef]
	organisationRef := params[configs.URLParamOrganisationRef]
	convertedOrganisationsRef := common.HexToHash(organisationsRef)
	convertedOrganisationRef := common.HexToHash(organisationRef)
	content := xml.Organisations(contracts, convertedOrganisationsRef, convertedOrganisationRef)
    w.Write(content)
}

func organisationsAll (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    activitiesRef := params[configs.URLParamOrganisationsRef]
	convertedActivitiesRef := common.HexToHash(activitiesRef)
	content := xml.OrganisationsAll(contracts, convertedActivitiesRef)
    w.Write(content)
}

func organisationsNum (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.OrganisationsNum(contracts)
    w.Write(content)
}

func organisationList (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    organisationsRef := params[configs.URLParamOrganisationsRef]
	convertedOrganisationsRef := common.HexToHash(organisationsRef)
	content := xml.OrganisationListXML(contracts, convertedOrganisationsRef)
    w.Write(content)
}

func organisationNum (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
    params := mux.Vars(r)
    organisationsRef := params[configs.URLParamOrganisationsRef]
	convertedOrganisationsRef := common.HexToHash(organisationsRef)
    content := xml.OrganisationNum(contracts, convertedOrganisationsRef)
    w.Write(content)
}

func orgsList (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.OrgsList(contracts)
    w.Write(content)
}

func orgs (contracts *contracts.Contracts, w http.ResponseWriter, r *http.Request) {

    xml.Header(w)
	params := mux.Vars(r)
    orgsRef := params[configs.URLParamOrgsRef]
	activityRef := params[configs.URLParamActivityRef]
	convertedOrgsRef := common.HexToHash(orgsRef)
	convertedActivityRef := common.HexToHash(activityRef)
	content := xml.OrgsID(contracts, convertedOrgsRef, convertedActivityRef)
    w.Write(content)
}

func orgsNum (contracts *contracts.Contracts, w http.ResponseWriter) {

    xml.Header(w)
    content := xml.OrgsNum(contracts)
    w.Write(content)
}

func handleRequests(contracts *contracts.Contracts) {

	r := mux.NewRouter().StrictSlash(true)

    // route all POST requests to defaultHandler:
    r.PathPrefix("/").HandlerFunc(postHandler).Methods(
                                                        "POST",
                                                        "PUT",
                                                        "PATCH",
                                                        "DELETE",
                                                        "COPY",
                                                        "HEAD",
                                                        "OPTIONS",
                                                        "LINK",
                                                        "UNLINK",
                                                        "PURGE",
                                                        "LOCK",
                                                        "UNLOCK",
                                                        "PROPFIND",
                                                        "VIEW",
                                                    )

    r.NotFoundHandler = http.HandlerFunc(notFound)

    r.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	home(w)
	}).Methods("GET")

    r.HandleFunc(configs.URLHelp, func(w http.ResponseWriter, r *http.Request) {
    	help(w)
	}).Methods("GET")

    // Activities
	r.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        activities(contracts, w, r)
	}).Methods("GET")

    r.HandleFunc(configs.URLActivitiesAll, func(w http.ResponseWriter, r *http.Request) {
        activitiesAll(contracts, w, r)
	}).Methods("GET")

	r.HandleFunc(configs.URLActivitiesList, func(w http.ResponseWriter, r *http.Request) {
    	activitiesList(contracts, w)
	}).Methods("GET")

	r.HandleFunc(configs.URLActivitiesTotal, func(w http.ResponseWriter, r *http.Request) {
    	activitiesNum(contracts, w)
	}).Methods("GET")

    // Activity
    r.HandleFunc(configs.URLActivityList, func(w http.ResponseWriter, r *http.Request) {
    	activityList(contracts, w, r)
	}).Methods("GET")

    r.HandleFunc(configs.URLActivityTotal, func(w http.ResponseWriter, r *http.Request) {
    	activityNum(contracts, w, r)
	}).Methods("GET")

    // Organisations
    r.HandleFunc(configs.URLOrganisations, func(w http.ResponseWriter, r *http.Request) {
        organisations(contracts, w, r)
    }).Methods("GET")

    r.HandleFunc(configs.URLOrganisationsAll, func(w http.ResponseWriter, r *http.Request) {
        organisationsAll(contracts, w, r)
	}).Methods("GET")

    r.HandleFunc(configs.URLOrganisationsList, func(w http.ResponseWriter, r *http.Request) {
        organisationsList(contracts, w)
    }).Methods("GET")

    r.HandleFunc(configs.URLOrganisationsTotal, func(w http.ResponseWriter, r *http.Request) {
        organisationsNum(contracts, w)
    }).Methods("GET")

    // Organisation
    r.HandleFunc(configs.URLOrganisationList, func(w http.ResponseWriter, r *http.Request) {
        organisationList(contracts, w, r)
    }).Methods("GET")

    r.HandleFunc(configs.URLOrganisationTotal, func(w http.ResponseWriter, r *http.Request) {
        organisationNum(contracts, w, r)
    }).Methods("GET")

    // Orgs
    r.HandleFunc(configs.URLOrgs, func(w http.ResponseWriter, r *http.Request) {
        orgs(contracts, w, r)
    }).Methods("GET")

    r.HandleFunc(configs.URLOrgsList, func(w http.ResponseWriter, r *http.Request) {
        orgsList(contracts, w)
    }).Methods("GET")

    r.HandleFunc(configs.URLOrgsTotal, func(w http.ResponseWriter, r *http.Request) {
        orgsNum(contracts, w)
    }).Methods("GET")

    srv := &http.Server{
		Handler:          r,
        Addr:             "127.0.0.1" + configs.ServerPort,
		WriteTimeout:     configs.WriteTimeout,
		ReadTimeout:      configs.ReadTimeout,
	}

    log.Fatal(srv.ListenAndServe())
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
