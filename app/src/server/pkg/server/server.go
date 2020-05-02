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
	content := xml.Activities(contracts, convertedActivitiesRef, convertedActivityRef)
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
	content := xml.OrganisationsID(contracts, convertedOrganisationsRef, convertedOrganisationRef)
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
	content := xml.OrganisationList(contracts, convertedOrganisationsRef)
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

	router := mux.NewRouter().StrictSlash(true)

    router.NotFoundHandler = http.HandlerFunc(notFound)

    router.HandleFunc(configs.URLHome, func(w http.ResponseWriter, r *http.Request) {
    	homePage(w)
	})

    // Activities
	router.HandleFunc(configs.URLActivities, func(w http.ResponseWriter, r *http.Request) {
        activities(contracts, w, r)
	})

	router.HandleFunc(configs.URLActivitiesList, func(w http.ResponseWriter, r *http.Request) {
    	activitiesList(contracts, w)
	})

	router.HandleFunc(configs.URLActivitiesTotal, func(w http.ResponseWriter, r *http.Request) {
    	activitiesNum(contracts, w)
	})

    router.HandleFunc(configs.URLActivityList, func(w http.ResponseWriter, r *http.Request) {
    	activityList(contracts, w, r)
	})

    router.HandleFunc(configs.URLActivityTotal, func(w http.ResponseWriter, r *http.Request) {
    	activityNum(contracts, w, r)
	})

    router.HandleFunc(configs.URLOrganisations, func(w http.ResponseWriter, r *http.Request) {
        organisations(contracts, w, r)
    })

    router.HandleFunc(configs.URLOrganisationsList, func(w http.ResponseWriter, r *http.Request) {
        organisationsList(contracts, w)
    })

    router.HandleFunc(configs.URLOrganisationsTotal, func(w http.ResponseWriter, r *http.Request) {
        organisationsNum(contracts, w)
    })

    router.HandleFunc(configs.URLOrganisationList, func(w http.ResponseWriter, r *http.Request) {
        organisationList(contracts, w, r)
    })

    router.HandleFunc(configs.URLOrganisationTotal, func(w http.ResponseWriter, r *http.Request) {
        organisationNum(contracts, w, r)
    })

    // Orgs
    router.HandleFunc(configs.URLOrgs, func(w http.ResponseWriter, r *http.Request) {
        orgs(contracts, w, r)
    })

    router.HandleFunc(configs.URLOrgsList, func(w http.ResponseWriter, r *http.Request) {
        orgsList(contracts, w)
    })

    router.HandleFunc(configs.URLOrgsTotal, func(w http.ResponseWriter, r *http.Request) {
        orgsNum(contracts, w)
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
