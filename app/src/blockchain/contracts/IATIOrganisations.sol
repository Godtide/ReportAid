pragma solidity ^0.4.24;

// IATI Organisations
// Steve Huckle

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  struct Organisation {
    string name;
    string namespaceCode;
    string baseIdentifier;
  }

  string[] orgReferences;
  mapping(string => Organisation) private organisations;

  event SetOrganisation(string _reference, string _name, string _namespaceCode, string _baseIdentifier);

  constructor() public {
  }

  function setOrganisation(string _reference, string _name, string _namespaceCode, string _baseIdentifier) public {
    require((bytes(_reference).length > 0) && (bytes(_name).length > 0));
    organisations[_reference].name = _name;
    organisations[_reference].namespaceCode = _namespaceCode;
    organisations[_reference].baseIdentifier = _baseIdentifier;
    orgReferences.push(_reference);
    emit SetOrganisation(_reference, _name, _namespaceCode, _baseIdentifier);
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

  function getOrganisationNamespaceCode(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].namespaceCode;
  }

  function getOrganisationBaseIdentifier(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].baseIdentifier;
  }
}
