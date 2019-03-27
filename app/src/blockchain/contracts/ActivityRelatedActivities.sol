pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract ActivityRelatedActivities {

  enum RelationType {
    NONE,
    PARENT,
    CHILD,
    SIBLING,
    COFUNDED,
    THIRDPARTY,
    MAX
  }

  struct RelatedActivity {
    uint8 relationType;
    bytes32 activityRef;
  }

  function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _relatedActivityRef, RelatedActivity memory _relatedActivity) public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) public view returns (RelatedActivity memory);

  function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) public view returns (bytes32);
  function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) public view returns (uint8);
}
