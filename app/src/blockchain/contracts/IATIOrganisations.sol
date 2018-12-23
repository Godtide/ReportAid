pragma solidity ^0.5.0;

// IATI Organisations
// Steve Huckle

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  struct Organisation {
    bytes32 name;
    bytes32 identifier;
  }

  bytes32[] orgReferences;
  mapping(bytes32 => Organisation) private organisations;

  event SetOrganisation(bytes32 _reference, bytes32 _name, bytes32 _identifier);

  constructor() public {
  }

  function setOrganisation(bytes32 _reference, bytes32 _name, bytes32 _identifier) public {
    require(_reference[0] != 0 && _name[0] != 0 && _identifier[0] != 0);

    organisations[_reference].name = _name;
    organisations[_reference].identifier = _identifier;
    if(!getOrganisationExists(_reference)) {
      orgReferences.push(_reference);
    }

    emit SetOrganisation(_reference, _name, _identifier);
  }

  function getOrganisationExists(bytes32 _reference) public view returns (bool) {
    require (_reference[0] != 0);

    uint256 index = Strings.getIndex(_reference, orgReferences);
    return index != orgReferences.length;
  }

  function getNumOrganisations() public view returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public view returns (bytes32) {
    require (_index < orgReferences.length);

    return orgReferences[_index];
  }

  function getOrganisationName(bytes32 _reference) public view returns (bytes32) {
    require (_reference[0] != 0);

    return organisations[_reference].name;
  }

  function getOrganisationIdentifier(bytes32 _reference) public view returns (bytes32) {
    require (_reference[0] != 0);

    return organisations[_reference].identifier;
  }
}
