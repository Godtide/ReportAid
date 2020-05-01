package configs

const (
    // URLHome - homepage
    URLHome = "/"

    // URLParamActivitiesRef - activities reference
    URLParamActivitiesRef = "activitiesRef"
    // URLParamActivityRef - activity reference
    URLParamActivityRef = "activityRef"

    // URLParamOrganisationsRef - Organisations reference
    URLParamOrganisationsRef = "organisationsRef"
    // URLParamOrganisationRef - Organisation reference
    URLParamOrganisationRef = "organisationRef"
    // URLParamOrgsRef - orgs ref'
    URLParamOrgsRef = "orgsRef"

    // URLActivities - activities XML URL
    URLActivities = "/activities/{" + URLParamActivitiesRef + "}/{" + URLParamActivityRef + "}"
    // URLActivitiesList - list activities XML URL
    URLActivitiesList = "/activities-list"
    // URLActivitiesTotal - total activities XML URL
    URLActivitiesTotal = "/activities-total"

    // URLActivityList - activities XML URL
    URLActivityList = "/activity-list/{" + URLParamActivitiesRef + "}"
    // URLActivityTotal - total activity for an activities file
    URLActivityTotal = "/activity-total/{" + URLParamActivitiesRef + "}"

    // URLOrganisations - Organisations XML URL
    URLOrganisations = "/organisations/{" + URLParamOrganisationsRef + "}/{" + URLParamOrganisationRef + "}"
    // URLOrganisationsList - list Organisations XML URL
    URLOrganisationsList = "/organisations-list"
    // URLOrganisationsTotal - total Organisations XML URL
    URLOrganisationsTotal = "/organisations-total"

    // URLOrganisationList - list Organisations XML URL
    URLOrganisationList = "/organisation-list/{" + URLParamOrganisationRef + "}"
    // URLOrganisationTotal - total Organisation XML URL
    URLOrganisationTotal = "/organisation-total"

    // URLOrgs - orgs XML URL
    URLOrgs = "/orgs/{" + URLParamOrgsRef + "}"
    // URLOrgsList - Orgs XML URL
    URLOrgsList = "/orgs-list"
    // URLOrgsTotal - total Orgs
    URLOrgsTotal = "/orgs-total"

)
