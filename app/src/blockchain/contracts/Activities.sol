pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activities {

  struct ReportingOrg {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct Activities {
    bytes32 activitiesRef;
    ReportingOrg reportingOrg;
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
  }

  function setActivities(Activities memory _activities) public;
  function setReportingOrg(bytes32 _activitiesRef, ReportingOrg memory _reportingOrg) public;
  function setVersion(bytes32 _activitiesRef, bytes32 _version) public;
  function setGeneratedTime(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _generatedTime) public;
  function setLinkedData(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _linkedData) public;

  function getNumActivities() public view returns (uint256);
  function getActivitiesReference(uint256 _index) public view returns (bytes32);

  function getActivities(bytes32 _activitiesRef) public view returns (Activities memory);
  function getReportingOrg(bytes32 _activitiesRef) public view returns (ReportingOrg memory);
  function getVersion(bytes32 _activitiesRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32);
  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32);
}
