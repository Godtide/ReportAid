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
    URLListActivities = "/activities-list"
    // URLTotalActivities - total activities XML URL
    URLTotalActivities = "/activities-total"

    // URLListActivity - activities XML URL
    URLListActivity = "/activity-list/{" + URLParamActivitiesRef + "}"
    // URLTotalActivity - total activity for an activities file
    URLTotalActivity = "/activity-total/{" + URLParamActivitiesRef + "}"

)
