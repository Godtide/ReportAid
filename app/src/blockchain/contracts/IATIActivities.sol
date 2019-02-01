pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivitys is Activities {

  bytes32[] activityReferences;
  mapping(bytes32 => Activity) private activities;

  event SetActivity(Activity _activity);

  function setActivity(Activity memory _activity) public {
    require (_activity.activityRef[0] != 0 && bytes(_activity.identifier).length > 0);

    activities[_activity.activityRef] = _activity;
    if(!getActivityExists(_activity.activityRef)) {
      activityReferences.push(_activity.activityRef);
    }

    emit SetActivity(_activity);
  }

  function getActivityExists(bytes32 _activityRef) public view returns (bool) {
    require (_activityRef[0] != 0);

    bool exists = false;
    if ( !(activityReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_activityRef, activityReferences);
      exists = (index != activityReferences.length);
    }
    return exists;
  }

  function getNumActivitys() public view returns (uint256) {
    return activityReferences.length;
  }

  function getActivityReference(uint256 _index) public view returns (bytes32) {
    require (_index < activityReferences.length);

    return activityReferences[_index];
  }

  function getActivity(bytes32 _activityRef) public view returns (Activity memory) {
    require (_activityRef[0] != 0);

    return activities[_activityRef];
  }

  function getActivityIdentifier(bytes32 _activityRef) public view returns (string memory) {
    require (_activityRef[0] != 0);

    return activities[_activityRef].identifier;
  }
}
