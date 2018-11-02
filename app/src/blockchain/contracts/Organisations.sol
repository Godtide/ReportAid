pragma solidity ^0.4.24;

contract Organisations {

  function setOrganisation(string _reference, string _name, string _type) public;

  function getOrganisationExists(string _reference) public constant returns (bool);
  function getNumOrganisations() public constant returns (uint256);
  function getOrganisationReference(uint256 _index) public constant returns (string);
  function getOrganisationName(string _reference) public constant returns (string);
  function getOrganisationType(string _reference) public constant returns (string);
}
