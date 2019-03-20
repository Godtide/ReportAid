pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract ActivityTerritories {

  struct Territory {
    uint8 percentage;
    bytes32 territory;
  }

  event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, Territory _territory);

  function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _territoryRef, Territory memory _territory) public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (Territory memory);

  function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (uint8);
  function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) public view returns (bytes32);
}
