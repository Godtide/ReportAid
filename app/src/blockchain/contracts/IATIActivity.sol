pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activity.sol";
import "./Strings.sol";

contract IATIActivity is Activity {

  bytes32[] activitiesRefs;
  mapping(bytes32 => bytes32[]) private activityRefs;
  mapping(bytes32 =>  mapping(bytes32 => Activity)) private activities;

  event SetActivity(bytes32 _activitiesRef, Activity memory _activity);

  function setActivity(bytes32 _activitiesRef, Activity memory _activity) public {
    require (_activitiesRef[0] != 0 &&
             _activity.activityRef[0] != 0 &&
             bytes(_activity.identifier).length > 0 &&
             _activity.reportingOrg.orgRef[0] != 0 &&
             _activity.reportingOrg.orgType > 0 &&
             bytes(_activity.title).length > 0 &&
             bytes(_activity.description).length > 0 &&
             _activity.lastUpdated[0] != 0 &&
             _activity.lang[0]  != 0 &&
             _activity.currency[0] != 0 &&
             _activity.hierarchy > Hierarchy.NONE &&
             _activity.hierarchy < Hierarchy.MAX &&
             _activity.budgetNotProvided >= budgetNotProvided.NONE &&
             _activity.budgetNotProvided < budgetNotProvided.MAX &&
             _activity.status > Status.NONE &&
             _activity.status < Status.MAX &&
             _activity.date[0] != 0 &&
             _activity.scope >= Scope.NONE &&
             _activity.scope < Scope.MAX);

    activities[_activitiesRef][_activity.activityRef].activity = _activity;

    if (!getExists(_activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activitiesRef);
    }

    if(!getExists(_activity.activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(_activity.activityRef);
    }

    emit SetActivity(_activitiesRef, _activity);
  }

  function getNumActivities(bytes32 _activitiesRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0);

    return activityRefs[_activitiesRef].length;
  }

  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _index < activityRefs[_activitiesRef].length);

    return activityRefs[_activitiesRef][_index];
  }

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (Activity memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

    return activities[_activitiesRef][_activityRef];
  }

  function getActivityIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].identifier;
  }

  function getReportingOrg(bytes32 _activitiesRef, , bytes32 _activityRef) public view returns (ReportingOrg memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].reportingOrg;
  }

  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].title;
  }

  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].description;
  }

  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lastUpdated;
  }

  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lang;
  }

  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].currency;
  }

  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (boolean) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].humanitarian;
  }

  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].hierarchy;
  }

  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].linkedData;
  }

  function getBudgetProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].budgetProvided;
  }

  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].status;
  }

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].date;
  }

  function getScope(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].scope;
  }
}
