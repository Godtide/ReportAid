pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activity {

  struct Activity {
    bytes32 activityRef;
    string identifier;
    string title;
    string description;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    boolean humanitarian;
    uint8 hierarchy;
    bytes32 linkedData;
    uint8 budgetProvided;
    unint8 status;
    bytes32 date;
    uint8 scope;
  }

  function setActivity(bytes32 _activitiesRef, Activity memory _activity) public;

  function getNumActivities(bytes32 _activitiesRef) public view returns (uint256);

  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32);

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (Activity memory);
  function getActivityIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (boolean);
  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getBudgetProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
}
