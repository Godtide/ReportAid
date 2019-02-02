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
}
