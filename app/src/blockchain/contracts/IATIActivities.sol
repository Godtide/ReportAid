pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesReferences;
  mapping(bytes32 => bytes32[]) private activityReferences;
  mapping(bytes32 =>  mapping(bytes32 => Activity)) private activities

  event SetActivities(OrgActivities memory _activities);
  event setActivity(bytes32 _activitiesRef, Activity memory _activity);

  function setActivities(OrgActivities memory _activities) public {
   return null;
  }

  function setActivity(bytes32 _activitiesRef, Activity memory _activity) public {
   return null;
  }


  function getActivitiesExists(bytes32 _activitiesRef) public view returns (bool) {
   return null;
  }

  function getActivityExists(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool) {
   return null;
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
