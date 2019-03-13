pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./ActivityParticipatingOrgs.sol";
import "./Strings.sol";

contract IATIActivityParticipatingOrgs is ActivityParticipatingOrgs {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private participatingOrgRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => ParticipatingOrg))) private participatingOrgs;

  function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef,  bytes32 _participatingOrgRef, ParticipatingOrg memory _participatingOrg) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _participatingOrgRef[0] != 0 &&
             _participatingOrg.orgRef[0] != 0 &&
             _participatingOrg.orgType >= ActivityParticipatingOrgs.minOrgType &&
             _participatingOrg.role > uint8(Role.NONE) &&
             _participatingOrg.role < uint8(Role.MAX) &&
             _participatingOrg.crsChannelCode >= ActivityParticipatingOrgs.minCrsChannelCode &&
             bytes(_participatingOrg.narrative).length > 0 &&
             _participatingOrg.lang[0] != 0 );

    participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef] = _participatingOrg;

    if(!Strings.getExists(_participatingOrgRef, participatingOrgRefs[_activitiesRef][_activityRef])) {
      participatingOrgRefs[_activitiesRef][_activityRef].push(_participatingOrgRef);
    }

    emit SetParticipatingOrg(_activitiesRef, _activityRef, _participatingOrgRef, _participatingOrg);
  }

  function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return participatingOrgRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < participatingOrgRefs[_activitiesRef][_activityRef].length);

    return participatingOrgRefs[_activitiesRef][_activityRef][_index];
  }

  function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (ParticipatingOrg memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef];
  }

  function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].orgRef;
  }

  function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].orgType;
  }

  function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].role;
  }

  function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].crsChannelCode;
  }

  function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (string memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].narrative;
  }

  function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _participatingOrgRef[0] != 0);

    return participatingOrgs[_activitiesRef][_activityRef][_participatingOrgRef].lang;
  }
}
