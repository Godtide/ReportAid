pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  bytes32[] orgReferences;
  mapping(bytes32 => Organisation) private organisations;

  event SetOrganisation(bytes32 _reference, Organisation _org);

  constructor() public {
  }

  function setOrganisation(Organisation _org) public {
    require (_org.reference[0] != 0 && bytes(_org.name).length > 0 && bytes(_org.identifier).length > 0);

    organisations[_org.reference].name = _org.name;
    organisations[_org.reference].identifier = _org.identifier;
    if(!getOrganisationExists(_org.reference)) {
      orgReferences.push(_org.reference);
    }

    emit SetOrganisation(_org.reference, _org);
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


  function getOrganisation(bytes32 _reference) public view returns (Organisation) {
    require (_reference[0] != 0);

    return organisations[_reference];
  }

  function getOrganisationName(bytes32 _reference) public view returns (string memory) {
    require (_reference[0] != 0);

    return organisations[_reference].name;
  }

  function getOrganisationIdentifier(bytes32 _reference) public view returns (string memory) {
    require (_reference[0] != 0);

    return organisations[_reference].identifier;
  }
}
