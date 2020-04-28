pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ActivityDates {

  enum DateCodes {
    NONE,
    PLANNEDSTART,
    ACTUALSTART,
    PLANNEDEND,
    ACTUALEND,
    MAX
  }

  struct Date {
    uint8 dateType;
    bytes32 date;
    string narrative;
  }

  event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, Date _date);

  function setDate(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _dateRef, Date memory _date) virtual public;

  function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) virtual public view returns (Date memory);

  function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) virtual public view returns (uint8);
  function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) virtual public view returns (bytes32);
  function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) virtual public view returns (string memory);
}
