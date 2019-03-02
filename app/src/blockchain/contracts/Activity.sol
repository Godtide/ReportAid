pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activity {

  enum Hierarchy {
    NONE,
    ACTIVITY,
    SUBACTIVITY,
    SUBSUBACTIVITY,
    MAX
  }

  struct ReportingOrg {
    uint8 orgType;
    bool isSecondary;
    bytes32 orgRef;
  }

  struct OrgActivity {
    bool humanitarian;
    uint8 hierarchy;
    ReportingOrg reportingOrg;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bytes32 linkedData;
    bytes32 identifier;
    bytes32 title;
  }

  function setActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity memory _activity) public;

  function getNumActivities(bytes32 _activitiesRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32);

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (OrgActivity memory);

  function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (ReportingOrg memory);
  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool);
  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
}
