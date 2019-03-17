pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;


import "./ActivitySectors.sol";
import "./Strings.sol";

contract IATIActivitySectors is ActivitySectors {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private sectorRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Sector))) private sectors;

  event SetSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, Sector _sector);

  function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, Sector memory _sector) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _sectorRef[0] != 0 &&
             _sector.dacCode >= ActivitySectors.minDACCode &&
             _sector.dacCode <= ActivitySectors.maxDACCode &&
             _sector.percentage <= 100 );

    sectors[_activitiesRef][_activityRef][_sectorRef] = _sector;

    if(!Strings.getExists(_sectorRef, sectorRefs[_activitiesRef][_activityRef])) {
      sectorRefs[_activitiesRef][_activityRef].push(_sectorRef);
    }

    emit SetSector(_activitiesRef, _activityRef, _sectorRef, _sector);
  }

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return sectorRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32)  {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < sectorRefs[_activitiesRef][_activityRef].length);

    return sectorRefs[_activitiesRef][_activityRef][_index];
  }

  function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) public view returns (Sector memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _sectorRef[0] != 0);

    return sectors[_activitiesRef][_activityRef][_sectorRef];
  }

  function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _sectorRef[0] != 0);

    return sectors[_activitiesRef][_activityRef][_sectorRef].dacCode;
  }

  function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _sectorRef[0] != 0);

    return sectors[_activitiesRef][_activityRef][_sectorRef].percentage;
  }
}
