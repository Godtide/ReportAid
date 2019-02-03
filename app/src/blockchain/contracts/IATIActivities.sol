pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesRefs;
  mapping(bytes32 =>  Activities) private activities;

  event SetActivities(OrgActivities memory _activities);

  function setActivities(Activities memory _activities) public {
    require (_activities.activitiesRef[0] != 0
             _activities.version[0] != 0 &&
             _activities.generatedTime[0] != 0);

    activities[_activities.activitiesRef] = _activities;

    if (!getExists(_activities.activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activities.activitiesRef);
    }

    if(!getExists(_activities.activity.activityRef, orgActivities[_activities.activitiesRef])) {
      activityRefs[_activities.activitiesRef].push(_activities.activity.activityRef);
    }

    emit SetActivities(_activities);
  }

  function getNumActivities() public view returns (uint256) {
    return activitiesRefs.length;
  }

  function getActivitiesReference(uint256 _index) public view returns (bytes32) {
    require (_index < activitiesRefs.length);

    return activitiesRefs[_index];
  }

  function getActivities(bytes32 _activitiesRef) public view returns (Activities memory) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef];
  }

  function getReportingOrg(bytes32 _activitiesRef) public view returns (ReportingOrg memory) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].reportingOrg;
  }

  function getVersion(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].version;
  }

  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].generatedTime;
  }

  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].linkedData;
  }
}
