pragma solidity ^0.5.0;

// IATI Organisations
// Steve Huckle

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  struct Organisation {
    string name;
    string identifier;
  }

  string[] orgReferences;
  mapping(string => Organisation) private organisations;

  event SetOrganisation(string _reference, string _name, string _identifier);

  constructor() public {
  }

  function setOrganisation(string memory _reference, string memory _name, string memory _identifier) public {
    require((bytes(_reference).length > 0) && (bytes(_name).length > 0)  && (bytes(_identifier).length > 0));

    organisations[_reference].name = _name;
    organisations[_reference].identifier = _identifier;
    if(!getOrganisationExists(_reference)) {
      orgReferences.push(_reference);
    }

    emit SetOrganisation(_reference, _name, _identifier);
  }

  function getOrganisationExists(string memory _reference) public view returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_reference, orgReferences);
    return index != orgReferences.length;
  }

  function getNumOrganisations() public view returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public view returns (string memory) {
    require(_index < orgReferences.length);

    return orgReferences[_index];
  }

  function getOrganisationName(string memory _reference) public view returns (string memory) {
    require(bytes(_reference).length > 0);

    return organisations[_reference].name;
  }

  function getOrganisationIdentifier(string memory _reference) public view returns (string memory) {
    require(bytes(_reference).length > 0);

    return organisations[_reference].identifier;
  }
}
