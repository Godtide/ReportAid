pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./ActivityTerritories.sol";
import "./Strings.sol";

contract IATIActivityTerritories is ActivityTerritories {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private territoryRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Territory))) private territories;

  event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, Territory _territory);

  function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _territoryRef, Territory memory _territory) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _territoryRef[0] != 0 &&
             _territory.territory[0] != 0 &&
             _territory.percentage <= 100 );

    territories[_activitiesRef][_activityRef][_territoryRef] = _territory;

    if(!Strings.getExists(_territoryRef, territoryRefs[_activitiesRef][_activityRef])) {
      territoryRefs[_activitiesRef][_activityRef].push(_territoryRef);
    }

    emit SetTerritory(_activitiesRef, _activityRef, _territoryRef, _territory);
  }

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return territoryRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < territoryRefs[_activitiesRef][_activityRef].length);

    return territoryRefs[_activitiesRef][_activityRef][_index];
  }

  function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (Territory memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _territoryRef[0] != 0);

    return territories[_activitiesRef][_activityRef][_territoryRef];
  }

  function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _territoryRef[0] != 0);

    return territories[_activitiesRef][_activityRef][_territoryRef].percentage;
  }

  function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _territoryRef[0] != 0);

    return territories[_activitiesRef][_activityRef][_territoryRef].territory;
  }
}
