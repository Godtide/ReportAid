pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Activities.sol";
import "./Strings.sol";

contract IATIActivities is Activities {

  bytes32[] activitiesRefs;
  mapping(bytes32 =>  OrgActivities) private activities;

  function setActivities(bytes32 _activitiesRef, OrgActivities memory _activities) override virtual public {
    require (_activitiesRef[0] != 0  &&
             _activities.version[0] != 0 &&
             _activities.generatedTime[0] != 0);

    activities[_activitiesRef] = _activities;

    if (!Strings.getExists(_activitiesRef, activitiesRefs)) {
      activitiesRefs.push(_activitiesRef);
    }

    emit SetActivities(_activitiesRef, _activities);
  }

  function getNumActivities() override virtual public view returns (uint256) {
    return activitiesRefs.length;
  }

  function getReference(uint256 _index) override virtual public view returns (bytes32) {
    require (_index < activitiesRefs.length);

    return activitiesRefs[_index];
  }

  function getActivities(bytes32 _activitiesRef) override virtual public view returns (OrgActivities memory) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef];
  }

  function getVersion(bytes32 _activitiesRef) override virtual public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].version;
  }

  function getGeneratedTime(bytes32 _activitiesRef) override virtual public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].generatedTime;
  }

  function getLinkedData(bytes32 _activitiesRef) override virtual public view returns (bytes32) {
  	require (_activitiesRef[0] != 0);

  	return activities[_activitiesRef].linkedData;
  }
}
