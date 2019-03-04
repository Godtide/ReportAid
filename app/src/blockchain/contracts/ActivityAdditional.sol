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
    SINGLELOCATION,
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
    bytes32 defaultAidType;
    uint256 defaultFinanceType;
    uint8 budgetNotProvided;
    uint8 status;
    uint8 scope;
    uint8 capitalSpend;
    uint8 collaborationType;
    uint8 defaultFlowType;
    uint8 defaultTiedStatus;
  }

  function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, Additional memory _activity) public;

  function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (Additional memory);

  function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (bytes32);

  function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint256);

  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
  function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8);
}
