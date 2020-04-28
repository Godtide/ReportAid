pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Activity {

  enum Hierarchy {
    NONE,
    ACTIVITY,
    SUBACTIVITY,
    SUBSUBACTIVITY,
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

  enum BudgetNotProvided {
    NONE,
    COMMERCIALRESTRICTIONS,
    LEGALRESTRICTIONS,
    RAPIDONSETEMERGENCY,
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
    uint8 status;
    uint8 budgetNotProvided;
    ReportingOrg reportingOrg;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bytes32 linkedData;
    bytes32 identifier;
    string title;
    string description;
  }

  function setActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity memory _activity) virtual public;

  function getNumActivities(bytes32 _activitiesRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, uint256 _index) virtual public view returns (bytes32);

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (OrgActivity memory);

  function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bytes32);
  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (ReportingOrg memory);
  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (string memory);
  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bytes32);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bytes32);
  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bytes32);
  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bool);
  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint8);
  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint8);
  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint8);
  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (bytes32);
  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (string memory);
}
