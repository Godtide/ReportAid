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

  enum CollaborationType {
    NONE,
    BILATERAL,
    MULTILATERALINFLOWS,
    BILATERALPRIVATE,
    MULTILATERALOUTFLOWS,
    PRIVATEOUTFLOWS,
    BILATERALCOREFUNDED,
    TRIANGULAR,
    MAX
  }

  enum TiedStatus {
    UNSPECIFIED,
    UNSPECIFIEDONE,
    NONE,
    PARTIALLYTIED,
    TIED,
    UNTIED,
    MAX
  }

  struct ReportingOrg {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct OrgActivity {
    string identifier;
    ReportingOrg reportingOrg;
    string title;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bool humanitarian;
    uint8 hierarchy;
    bytes32 linkedData;
    uint8 budgetNotProvided;
    uint8 status;
    uint8 scope;
    uint8 capitalSpend;
    uint8 collaborationType;
    uint8 defaultFlowType;
    uint256 defaultFinanceType;
    bytes32 defaultAidType;
    uint8 defaultTiedStatus;
  }

  event SetActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity _activity);

  function setActivity(bytes32 _activitiesRef, bytes32 activityRef, OrgActivity memory _activity) public;

  function getNumActivities(bytes32 _activitiesRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, uint256 _index) public view returns (bytes32);

  function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (OrgActivity memory);

  function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (ReportingOrg memory);
  function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (string memory);
  function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bool);
  function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getScope(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
  function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (bytes32);
  function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint8);
}
