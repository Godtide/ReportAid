pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./ActivityRelatedActivities.sol";
import "./Strings.sol";

contract IATIActivityRelatedActivities is ActivityRelatedActivities {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private relatedActivityRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => RelatedActivity))) private relatedActivities;

  event SetRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _relatedActivityRef, RelatedActivity _relatedActivity);

  function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _relatedActivityRef, RelatedActivity memory _relatedActivity) override virtual public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _relatedActivityRef[0] != 0 &&
             _relatedActivity.activityRef[0] != 0 &&
             _relatedActivity.relationType > uint8(RelationType.NONE) &&
             _relatedActivity.relationType < uint8(RelationType.MAX) );

    relatedActivities[_activitiesRef][_activityRef][_relatedActivityRef] = _relatedActivity;

    if(!Strings.getExists(_relatedActivityRef, relatedActivityRefs[_activitiesRef][_activityRef])) {
      relatedActivityRefs[_activitiesRef][_activityRef].push(_relatedActivityRef);
    }

    emit SetRelatedActivity(_activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity);
  }

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return relatedActivityRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) override virtual public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < relatedActivityRefs[_activitiesRef][_activityRef].length);

    return relatedActivityRefs[_activitiesRef][_activityRef][_index];
  }

  function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) override virtual public view returns (RelatedActivity memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _relatedActivityRef[0] != 0);

    return relatedActivities[_activitiesRef][_activityRef][_relatedActivityRef];
  }

  function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) override virtual public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _relatedActivityRef[0] != 0);

    return relatedActivities[_activitiesRef][_activityRef][_relatedActivityRef].activityRef;
  }

  function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) override virtual public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _relatedActivityRef[0] != 0);

    return relatedActivities[_activitiesRef][_activityRef][_relatedActivityRef].relationType;
  }
}
