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

  struct Dates {
    uint8 dateType;
    bytes32 date;
    string narrative;
  }

  function setActivitydates(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _dateRef, Dates memory _date) public;

  function getNumActivityDates(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getActivityDatesReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getActivityDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (Dates memory);
  function getActivityDateCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (uint8);
  function getActivityDateISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (bytes32);
  function getActivityDateNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) public view returns (string memory);
}
