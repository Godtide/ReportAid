package configs

const (
    // URLHome - homepage
    URLHome = "/"

    // URLParamActivitiesRef - activities reference
    URLParamActivitiesRef = "activitiesRef"
    // URLParamActivityRef - activity reference
    URLParamActivityRef = "activityRef"

    // URLActivities - activities XML URL
    URLActivities = "/activities/{" + URLParamActivitiesRef + "}/{" + URLParamActivityRef + "}"
    // URLListActivities - list activities XML URL
    URLListActivities = "/list-activities"
    // URLTotalActivities - total activities XML URL
    URLTotalActivities = "/total-activities"

    // URLListActivity - activities XML URL
    URLListActivity = "/list-activity/{" + URLParamActivitiesRef + "}"
    // URLTotalActivity - total activity for an activities file
    URLTotalActivity = "/total-activity/{" + URLParamActivitiesRef + "}"

)
