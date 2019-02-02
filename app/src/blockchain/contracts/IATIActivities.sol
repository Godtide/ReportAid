pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesRefs;
  mapping(bytes32 => bytes32[]) private activityRefs;
  mapping(bytes32 =>  mapping(bytes32 => Activity)) private activities

  event SetActivities(OrgActivities memory _activities);
  event SetActivity(bytes32 _activitiesRef, Activity memory _activity);
  event SetReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef, ReportingOrg memory _reportingOrg);
  event SetVersion(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _version);
  event SetGeneratedTime(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _generatedTime);
  event SetLinkedData(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _linkedData);

  function setActivities(OrgActivities memory _activities) public {
    require (_activities.activity.activityRef[0] != 0 &&
             _activities.activity.identifier[0] != 0 &&
             bytes(_activities.activity.identifier).length > 0 &&
             _activities.activity.lastUpdated[0] != 0 &&
              _activities.activity.lang[0]  != 0 &&
             _activities.activity.currency[0] != 0 &&
             _activities.activity.hierarchy > 0 &&
             _activities.activity.hierarchy < 4 &&
             _activities.reportingOrg.orgRef[0] != 0 &&
             _activities.reportingOrg.orgType > 0 &&
             _activities.activitiesRef[0] != 0 &&
             _activities.version[0] != 0 &&
             _activities.generatedTime[0] != 0);

    activities[_activities.activitiesRef][_activities.activity.activityRef] = _activities;

    if (!getExists(_activities.activitiesRef, activitiesReferences)) {
      activitiesRefs.push(_activities.activitiesRef);
    }

    if(!getExists(_activities.activity.activityRef, activityReferences[activities.activitiesRef])) {
      activityRefs[_activities.activitiesRef].push(_activities.activity.activityRef);
    }

    emit SetActivities(_activities);
  }

  function setActivity(bytes32 _activitiesRef, Activity memory _activity) public {
    require (_activitiesRef[0] != 0 &&
             _activity.activityRef[0] != 0 &&
             _activity.identifier[0] != 0 &&
             bytes(_activity.identifier).length > 0 &&
             _activity.lastUpdated[0] != 0 &&
             _activity.lang[0]  != 0 &&
             _activity.currency[0] != 0 &&
             _activity.hierarchy > 0 &&
             _activity.hierarchy < 4);

    activities[_activitiesRef][_activity.activityRef].activity = _activity;

    if (!getExists(_activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activity.activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activity.activityRef);
    }

    emit SetActivity(_activitiesRef, _activity);
  }

  function setReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef, ReportingOrg memory _reportingOrg) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _reportingOrg.orgRef[0] != 0 &&
             _reportingOrg.orgType > 0);

    activities[_activitiesRef][_activityRef].reportingOrg = _reportingOrg;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activityRef);
    }

    emit SetReportingOrg(_activitiesRef, _activityRef, _activity);
  }

  function setVersion(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _version) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _version[0] != 0);

    activities[_activitiesRef][_activityRef].version = _version;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activityRef);
    }

    emit SetVersion(_activitiesRef, _activityRef, _version);
  }

  function setGeneratedTime(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _generatedTime) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _generatedTime[0] != 0);

    activities[_activitiesRef][_activityRef].generatedTime = _generatedTime;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activityRef);
    }

    emit SetGeneratedTime(_activitiesRef, _activityRef, _generatedTime);
  }

  function setLinkedData(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _linkedData) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _linkedData[0] != 0);

    activities[_activitiesRef][_activityRef].linkedData = _linkedData;

    if (!getExists(_activitiesRef, activitiesRefs)) {
       activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activityRef);
    }

    emit SetLinkedData(_activitiesRef, _activityRef, _linkedData);
  }

  function getNumActivities() public view returns (uint256) {
   return null;
  }

  function getNumActivity(bytes32 _activitiesRef) public view returns (uint256) {
   return null;
  }

  function getActivitiesReference(uint256 _index) public view returns (bytes32) {
   return null;
  }

  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32) {
   return null;
  }

  function getActivities(bytes32 _activitiesRef) public view returns (OrgActivities memory) {
   return null;
  }

  function getVersion(bytes32 _activitiesRef) public view returns (bytes32) {
   return null;
  }

  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32) {
   return null;
  }

  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32) {
   return null;
  }

  function getReportingOrg(bytes32 _activitiesRef) public view returns (bytes32) {
   return null;
  }

  function getReportingOrgType(bytes32 _activitiesRef) public view returns (uint8) {
   return null;
  }

  function getReportingOrgIsSecondary(bytes32 _activitiesRef) public view returns (bool) {
   return null;
  }

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (Activity memory) {
   return null;
  }

  function getActivityIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
   return null;
  }

  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
   return null;
  }

  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
   return null;
  }

  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
   return null;
  }

  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (boolean) {
   return null;
  }

  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
   return null;
  }

  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
   return null;
  }

  function getBudgetProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
   return null;
  }
}
