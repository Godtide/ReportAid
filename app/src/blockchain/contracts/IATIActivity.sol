pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Activity.sol";
import "./Strings.sol";

contract IATIActivity is Activity {

  mapping(bytes32 => bytes32[]) private activityRefs;
  mapping(bytes32 =>  mapping(bytes32 => OrgActivity)) private activities;

  event SetActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity _activity);

  function setActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity memory _activity) override virtual public  {
    require (_activitiesRef[0] != 0 &&
             activityRef[0] != 0 &&
             _activity.identifier[0] != 0 &&
             _activity.reportingOrg.orgRef[0] != 0 &&
             _activity.reportingOrg.orgType > 0 &&
             bytes(_activity.title).length > 0 &&
             _activity.lastUpdated[0] != 0 &&
             _activity.lang[0]  != 0 &&
             _activity.currency[0] != 0 &&
             _activity.hierarchy > uint8(Hierarchy.NONE) &&
             _activity.hierarchy < uint8(Hierarchy.MAX) &&
             _activity.status > uint8(Status.NONE) &&
             _activity.status < uint8(Status.MAX) &&
             _activity.budgetNotProvided > uint8(BudgetNotProvided.NONE) &&
             _activity.budgetNotProvided < uint8(BudgetNotProvided.MAX) &&
             bytes(_activity.description).length > 0);

    activities[_activitiesRef][activityRef] = _activity;

    if(!Strings.getExists(activityRef, activityRefs[_activitiesRef])) {
      activityRefs[_activitiesRef].push(activityRef);
    }

    emit SetActivity(_activitiesRef, activityRef, _activity);
  }

  function getNumActivities(bytes32 _activitiesRef) override virtual public  view returns (uint256) {
    require (_activitiesRef[0] != 0);

    return activityRefs[_activitiesRef].length;
  }

  function getReference(bytes32 _activitiesRef, uint256 _index) override virtual public  view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _index < activityRefs[_activitiesRef].length);

    return activityRefs[_activitiesRef][_index];
  }

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (OrgActivity memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

    return activities[_activitiesRef][_activityRef];
  }

  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (string memory) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].title;
  }

  function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].identifier;
  }

  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (ReportingOrg memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].reportingOrg;
  }

  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lastUpdated;
  }

  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].lang;
  }

  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].currency;
  }

  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bool) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].humanitarian;
  }

  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].hierarchy;
  }

  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].status;
  }

  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].budgetNotProvided;
  }

  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (bytes32) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].linkedData;
  }


  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public  view returns (string memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 );

  	return activities[_activitiesRef][_activityRef].description;
  }

}
