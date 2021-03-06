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

    // URLReportAidNamespace - define namespace for extra info
    URLReportAidNamespace = "http://reportaid.org/ns#"

    // URLHelp - help
    URLHelp = "/help"

    // URLActivities -  URL
    URLActivities = "/activities"
    // URLActivitiesActivity - activities XML URL
    URLActivitiesActivity = URLActivities + "/{" + URLParamActivitiesRef + "}/{" + URLParamActivityRef + "}"
    // URLActivitiesAll - all activities activity XML URL
    URLActivitiesAll = URLActivities + "/{" + URLParamActivitiesRef + "}"
    // URLActivitiesList - list activities XML URL
    URLActivitiesList = URLActivities + "-list"
    // URLActivitiesTotal - total activities XML URL
    URLActivitiesTotal = URLActivities + "-total"

    // URLActivityList - activities XML URL
    URLActivityList = "/activity-list/{" + URLParamActivitiesRef + "}"
    // URLActivityTotal - total activity for an activities file
    URLActivityTotal = "/activity-total/{" + URLParamActivitiesRef + "}"

    // URLOrganisations - Organisations XML URL
    URLOrganisations = "/organisations/{" + URLParamOrganisationsRef + "}/{" + URLParamOrganisationRef + "}"
    // URLOrganisationsAll - all organisations organisation XML URL
    URLOrganisationsAll = "/organisations/{" + URLParamOrganisationsRef + "}"
    // URLOrganisationsList - list Organisations XML URL
    URLOrganisationsList = "/organisations-list"
    // URLOrganisationsTotal - total Organisations XML URL
    URLOrganisationsTotal = "/organisations-total"

    // URLOrganisationList - list Organisations XML URL
    URLOrganisationList = "/organisation-list/{" + URLParamOrganisationsRef + "}"
    // URLOrganisationTotal - total Organisation XML URL
    URLOrganisationTotal = "/organisation-total"

    // URLOrgs - orgs XML URL
    URLOrgs = "/orgs/{" + URLParamOrgsRef + "}"
    // URLOrgsList - Orgs XML URL
    URLOrgsList = "/orgs-list"
    // URLOrgsTotal - total Orgs
    URLOrgsTotal = "/orgs-total"

)
