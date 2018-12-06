pragma solidity ^0.4.24;

// IATI Organisations
// Steve Huckle

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  struct Organisation {
    string name;
  }

  string[] orgReferences;
  mapping(string => Organisation) private organisations;

  event SetOrganisation(string _reference, string _name);

  constructor() public {
  }

  function setOrganisation(string _reference, string _name) public {
    require((bytes(_reference).length > 0) && (bytes(_name).length > 0));
    organisations[_reference].name = _name;
    orgReferences.push(_reference);
    emit SetOrganisation(_reference, _name);
  }

  function getOrganisationExists(string _reference) public constant returns (bool) {
    require(bytes(_reference).length > 0);
    uint256 index = Strings.getIndex(_reference, orgReferences);
    return index != orgReferences.length;
  }

  function getNumOrganisations() public constant returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public constant returns (string) {
    require(_index < orgReferences.length);
    return orgReferences[_index];
  }

  function getOrganisationName(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].name;
  }
}
