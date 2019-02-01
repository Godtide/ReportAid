pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activities {

  struct Activity {
    bytes32 activityRef;
    string identifier;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    boolean humanitarian;
    uint8 hierarchy;
    bytes32 linkedData;
    uint8 budgetProvided;
  }

  struct ReportingOrganisation {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct OrgActivities {
    Activity activity;
    ReportingOrganisation reportingOrg;
    bytes32 activitiesRef;
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
  }

  function setActivities(OrgActivities memory _activities) public;
  function setActivity(bytes32 _activitiesRef, Activity memory _activity) public;

  function getActivitiesExists(bytes32 _activitiesRef) public view returns (bool);
  function getActivityExists(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool);

  function getNumActivities() public view returns (uint256);
  function getNumActivity(bytes32 _activitiesRef) public view returns (uint256);

  function getActivitiesReference(uint256 _index) public view returns (bytes32);
  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32);

  function getActivities(bytes32 _activitiesRef) public view returns (OrgActivities memory);
  function getVersion(bytes32 _activitiesRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _activitiesRef) public view returns (bytes32);
  function getLinkedData(bytes32 _activitiesRef) public view returns (bytes32);
  function getReportingOrg(bytes32 _activitiesRef) public view returns (bytes32);
  function getReportingOrgType(bytes32 _activitiesRef) public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _activitiesRef) public view returns (bool);

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
