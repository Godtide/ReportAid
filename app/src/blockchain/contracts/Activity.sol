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

  enum BudgetNotProvided {
    NONE,
    COMMERCIALRESTRICTIONS,
    LEGALRESTRICTIONS,
    RAPIDONSETEMERGENCY,
    MAX
  }

  enum Status {
    NONE,
    PIPELINEIDENTIFICATION,
    IMPLEMENTATION,
    FINALISATION,
    CLOSED,
    CANCELLED,
    SUSPENDED,
    MAX
  }

  enum Scope {
    NONE,
    GLOBAL,
    REGIONAL,
    MULTINATIONAL,
    NATIONAL,
    SUBNATIONALMULTIFIRSTLEVEL,
    SUBNATIONALSINGLEFIRSTLEVEL,
    SUBNATIONALSINGLESECONDLEVEL,
    MAX
  }

  struct ReportingOrg {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }


  struct OrgActivity {
    bytes32 activityRef;
    string identifier;
    ReportingOrg reportingOrg;
    string title;
    string description;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bool humanitarian;
    uint8 hierarchy;
    bytes32 linkedData;
    uint8 budgetNotProvided;
    uint8 status;
    bytes32 date;
    uint8 scope;
  }

  function setActivity(bytes32 _activitiesRef, OrgActivity memory _activity) public;

  function getNumActivities() public view returns (uint256);
  function getNumActivity(bytes32 _activitiesRef) public view returns (uint256);

  function getActivitiesReference(uint256 _index) public view returns (bytes32);
  function getActivityReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32);

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (OrgActivity memory);
  function getActivityIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (ReportingOrg memory);
  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool);
  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getDate(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getScope(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
}
