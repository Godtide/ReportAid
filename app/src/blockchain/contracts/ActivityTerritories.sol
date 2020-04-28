pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ActivityTerritories {

  struct Territory {
    uint8 percentage;
    bytes32 territory;
  }

  event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, Territory _territory);

  function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _territoryRef, Territory memory _territory) virtual public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) virtual public view returns (Territory memory);

  function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) virtual public view returns (uint8);
  function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) virtual public view returns (bytes32);
}
