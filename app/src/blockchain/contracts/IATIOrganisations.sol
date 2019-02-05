pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  bytes32[] organisationsRefs;
  mapping(bytes32 =>  Orgs) private organisations;

  event SetOrganisations(Orgs _organisations);

  function setOrganisations(Orgs memory _organisations) public {
    require (_organisations._organisationsRef[0] != 0 &&
             _organisations.version[0] != 0 &&
             _organisations.generatedTime[0] != 0);

    organisations[_organisations._organisationsRef] = _organisations;

    if (!Strings.getExists(_organisations._organisationsRef, organisationsRefs)) {
        organisationsRefs.push(_organisations._organisationsRef);
    }

    emit SetOrganisations(_organisations);
  }

  function getNumOrganisations() public view returns (uint256) {
    return organisationsRefs.length;
  }

  function getOrganisationsReference(uint256 _index) public view returns (bytes32) {
    require (_index < organisationsRefs.length);

    return organisationsRefs[_index];
  }

  function getOrganisations(bytes32 _organisationsRef) public view returns (Orgs memory) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef];
  }

  function getVersion(bytes32 _organisationsRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef].version;
  }

  function getGeneratedTime(bytes32 _organisationsRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef].generatedTime;
  }
}
