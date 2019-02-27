pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activities {

  struct OrgActivities {
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
  }

  event SetActivities(bytes32 _activitiesRef, OrgActivities _activities);

  function setActivities(bytes32 _activitiesRef, OrgActivities memory _activities) public;

  function getNumActivities() public view returns (uint256);
  function getReference(uint256 _index) public view returns (bytes32);

  function getActivities(bytes32 _activitiesRef) public view returns (OrgActivities memory);
  function getVersion(bytes32 _activitiesRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32);
  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32);
}
