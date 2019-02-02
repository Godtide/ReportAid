pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesRefs;
  mapping(bytes32 =>  Activities) private activities;

  event SetActivities(OrgActivities memory _activities);
  event SetReportingOrg(bytes32 _activitiesRef, ReportingOrg memory _reportingOrg);
  event SetVersion(bytes32 _activitiesRef, bytes32 _version);
  event SetGeneratedTime(bytes32 _activitiesRef, bytes32 _generatedTime);
  event SetLinkedData(bytes32 _activitiesRef, bytes32 _linkedData);

  function setActivities(Activities memory _activities) public {
    require (_activities.activitiesRef[0] != 0 &&
             _activities.reportingOrg.orgRef[0] != 0 &&
             _activities.reportingOrg.orgType > 0 &&
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

  function setReportingOrg(bytes32 _activitiesRef, ReportingOrg memory _reportingOrg) public {
    require (_activitiesRef[0] != 0 &&
             _reportingOrg.orgRef[0] != 0 &&
             _reportingOrg.orgType > 0);

    activities[_activitiesRef].reportingOrg = _reportingOrg;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    emit SetReportingOrg(_activitiesRef, _activity);
  }

  function setVersion(bytes32 _activitiesRef, bytes32 _version) public {
    require (_activitiesRef[0] != 0 &&
             _version[0] != 0);

    activities[_activitiesRef].version = _version;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    emit SetVersion(_activitiesRef, _version);
  }

  function setGeneratedTime(bytes32 _activitiesRef, bytes32 _generatedTime) public {
    require (_activitiesRef[0] != 0 &&
             _generatedTime[0] != 0);

    activities[_activitiesRef].generatedTime = _generatedTime;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    emit SetGeneratedTime(_activitiesRef, _generatedTime);
  }

  function setLinkedData(bytes32 _activitiesRef, bytes32 _linkedData) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _linkedData[0] != 0);

    activities[_activitiesRef].linkedData = _linkedData;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    emit SetLinkedData(_activitiesRef, _linkedData);
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
