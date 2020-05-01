package configs

const (
    // URLHome - homepage
    URLHome = "/"

    // URLParamActivitiesRef - activities reference
    URLParamActivitiesRef = "activitiesRef"
    // URLParamActivityRef - activity reference
    URLParamActivityRef = "activityRef"

    // URLParamOrganisationsRef - Organisations reference
    URLParamOrganisationsRef = "OrganisationsRef"
    // URLParamOrganisationRef - Organisation reference
    URLParamOrganisationRef = "OrganisationRef"

    // URLActivities - activities XML URL
    URLActivities = "/activities/{" + URLParamActivitiesRef + "}/{" + URLParamActivityRef + "}"
    // URLListActivities - list activities XML URL
    URLListActivities = "/activities-list"
    // URLTotalActivities - total activities XML URL
    URLTotalActivities = "/activities-total"

    // URLListActivity - activities XML URL
    URLListActivity = "/activity-list/{" + URLParamActivitiesRef + "}"
    // URLTotalActivity - total activity for an activities file
    URLTotalActivity = "/activity-total/{" + URLParamActivitiesRef + "}"

    // URLOrganisations - Organisations XML URL
    URLOrganisations = "/organisations/{" + URLParamOrganisationsRef + "}/{" + URLParamOrganisationRef + "}"
    // URLListOrganisations - list Organisations XML URL
    URLListOrganisations = "/organisations-list"
    // URLTotalOrganisations - total Organisations XML URL
    URLTotalOrganisations = "/organisations-total"

    // URLListOrganisation - Organisations XML URL
    URLListOrganisation = "/organisation-list/{" + URLParamOrganisationsRef + "}"
    // URLTotalOrganisation - total Organisation for an Organisations file
    URLTotalOrganisation = "/organisation-total/{" + URLParamOrganisationsRef + "}"

)
