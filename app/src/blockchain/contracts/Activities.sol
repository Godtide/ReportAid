pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Activities {

  struct OrgActivities {
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
  }

  event SetActivities(bytes32 _activitiesRef, OrgActivities _activities);

  function setActivities(bytes32 _activitiesRef, OrgActivities memory _activities) virtual public;

  function getNumActivities() virtual public view returns (uint256);
  function getReference(uint256 _index) virtual public view returns (bytes32);

  function getActivities(bytes32 _activitiesRef) virtual public view returns (OrgActivities memory);
  function getVersion(bytes32 _activitiesRef) virtual public view returns (bytes32);
  function getGeneratedTime(bytes32 _activitiesRef) virtual public view returns (bytes32);
  function getLinkedData(bytes32 _activitiesRef) virtual public view returns (bytes32);
}
