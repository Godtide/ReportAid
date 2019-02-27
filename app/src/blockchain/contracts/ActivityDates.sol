pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract ActivityDates {

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

  function setDate(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _dateRef, Date memory _date) public;

  function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (Date memory);

  function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (uint8);
  function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (bytes32);
  function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (string memory);
}
