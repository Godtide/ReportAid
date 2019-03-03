pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract ActivityAdditional {

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

  struct Additional {
    uint8 budgetNotProvided;
    uint8 status;
    uint8 scope;
    uint8 capitalSpend;
    uint8 collaborationType;
    uint8 defaultFlowType;
    uint8 defaultTiedStatus;
    uint256 defaultFinanceType;
    bytes32 defaultAidType;
  }

  function setActivityAdditional(bytes32 _activitiesRef, bytes32 activityRef, bytes32 _additionalRef, Additional memory _activity) public;

  function getNumAdditional(bytes32 _activitiesRef, bytes32 activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 activityRef, uint256 _index) public view returns (bytes32);

  function getActivityAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (Additional memory);

  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint256);
  function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (bytes32);
  function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
}
