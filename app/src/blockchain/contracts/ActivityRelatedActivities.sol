pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ActivityRelatedActivities {

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

  function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _relatedActivityRef, RelatedActivity memory _relatedActivity) virtual public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) virtual public view returns (RelatedActivity memory);

  function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) virtual public view returns (bytes32);
  function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) virtual public view returns (uint8);
}
