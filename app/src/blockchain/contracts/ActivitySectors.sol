pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ActivitySectors {

  uint256 constant minDACCode = 11110;
  uint256 constant maxDACCode = 99820;

  struct Sector {
    uint8 percentage;
    uint256 dacCode;
  }

  function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, Sector memory _sector) virtual public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) virtual public view returns (Sector memory);

  function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) virtual public view returns (uint256);
  function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) virtual public view returns (uint8);
}
