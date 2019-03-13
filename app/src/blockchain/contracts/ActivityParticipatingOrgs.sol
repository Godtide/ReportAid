pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract ActivityParticipatingOrgs {

  uint8 constant minOrgType = 10;
  uint256 constant minCrsChannelCode = 10000;

  enum Role {
    NONE,
    FUNDING,
    ACCOUNTABLE,
    EXTENDING,
    IMPLEMENTING,
    MAX
  }

  struct ParticipatingOrg {
    uint8 orgType;
    uint8 role;
    uint256 crsChannelCode;
    bytes32 lang;
    bytes32 orgRef;
    string narrative;
  }

  event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _particpatingOrgRef, ParticipatingOrg _participatingOrg);

  function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _particpatingOrgRef, ParticipatingOrg memory _participatingOrg) public;

  function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (ParticipatingOrg memory);

  function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (bytes32);
  function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (uint8);
  function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (uint8);
  function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (uint256);
  function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (string memory);
  function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) public view returns (bytes32);
}
