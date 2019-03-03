pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./ActivityAdditional.sol";
import "./Strings.sol";

contract IATIActivityAdditional is ActivityAdditional {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private additionalRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Additional))) private additionals;

  event SetActivityAdditional(bytes32 _activitiesRef, bytes32 activityRef, bytes32 _additionalRef, Additional _additional);

  function setActivityAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, Additional memory _additional) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _additionalRef[0] != 0 &&
             _additional.budgetNotProvided >= uint8(BudgetNotProvided.NONE) &&
             _additional.budgetNotProvided < uint8(BudgetNotProvided.MAX) &&
             _additional.status > uint8(Status.NONE) &&
             _additional.status < uint8(Status.MAX) &&
             _additional.scope >= uint8(Scope.NONE) &&
             _additional.scope < uint8(Scope.MAX) &&
             _additional.capitalSpend <= 100 &&
             _additional.collaborationType > uint8(CollaborationType.NONE) &&
             _additional.collaborationType < uint8(CollaborationType.MAX) &&
             _additional.defaultAidType[0] != 0 &&
             _additional.defaultTiedStatus > uint8(TiedStatus.NONE) &&
             _additional.defaultTiedStatus < uint8(TiedStatus.MAX));

    additionals[_activitiesRef][_activityRef][_additionalRef] = _additional;

    if(!Strings.getExists(_additionalRef, additionalRefs[_activitiesRef][_activityRef])) {
      additionalRefs[_activitiesRef][_activityRef].push(_additionalRef);
    }

    emit SetActivityAdditional(_activitiesRef, _activityRef, _additionalRef, _additional);
  }

  function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return additionalRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < additionalRefs[_activitiesRef][_activityRef].length);

    return additionalRefs[_activitiesRef][_activityRef][_index];
  }

  function getActivityAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (Additional memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

    return additionals[_activitiesRef][_activityRef][_additionalRef];
  }

  function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].budgetNotProvided;
  }

  function getStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].status;
  }

  function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
  	require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].scope;
  }

  function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].capitalSpend;
  }

  function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].collaborationType;
  }

  function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].defaultFlowType;
  }

  function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].defaultFinanceType;
  }

  function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].defaultAidType;
  }

  function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _additionalRef[0] != 0);

  	return additionals[_activitiesRef][_activityRef][_additionalRef].defaultTiedStatus;
  }

}
