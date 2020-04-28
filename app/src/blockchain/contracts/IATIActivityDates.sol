pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./ActivityDates.sol";
import "./Strings.sol";

contract IATIActivityDates is ActivityDates {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private dateRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Date))) private dates;

  function setDate(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _dateRef, Date memory _date) override virtual public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _dateRef[0] != 0 &&
             _date.dateType > uint8(DateCodes.NONE) &&
             _date.dateType < uint8(DateCodes.MAX) &&
             _date.date[0] != 0 &&
             bytes(_date.narrative).length > 0 );

    dates[_activitiesRef][_activityRef][_dateRef] = _date;

    if(!Strings.getExists(_dateRef, dateRefs[_activitiesRef][_activityRef])) {
      dateRefs[_activitiesRef][_activityRef].push(_dateRef);
    }

    emit SetDate(_activitiesRef, _activityRef, _dateRef, _date);
  }

  function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return dateRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) override virtual public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < dateRefs[_activitiesRef][_activityRef].length);

    return dateRefs[_activitiesRef][_activityRef][_index];
  }

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) override virtual public view returns (Date memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _dateRef[0] != 0);

    return dates[_activitiesRef][_activityRef][_dateRef];
  }

  function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) override virtual public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _dateRef[0] != 0);

    return dates[_activitiesRef][_activityRef][_dateRef].dateType;
  }

  function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) override virtual public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _dateRef[0] != 0);

    return dates[_activitiesRef][_activityRef][_dateRef].date;
  }

  function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) override virtual public view returns (string memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _dateRef[0] != 0);

    return dates[_activitiesRef][_activityRef][_dateRef].narrative;
  }
}
