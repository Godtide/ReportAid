pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Activities {

  struct Activity {
    bytes32 activityRef;
    string identifier;
  }

  function setActivity(Activity memory _activity) public;

  function getActivityExists(bytes32 _activityRef) public view returns (bool);
  function getNumActivities() public view returns (uint256);
  function getActivityReference(uint256 _index) public view returns (bytes32);
  function getActivity(bytes32 _activityRef) public view returns (Activity memory);
  function getActivityIdentifier(bytes32 _activityRef) public view returns (string memory);
}
