pragma solidity ^0.4.24;

contract Organisations {

  function setOrganisation(string _reference, string _name, string _identifier) public;

  function getOrganisationExists(string _reference) public view returns (bool);
  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (string);
  function getOrganisationName(string _reference) public view returns (string);
  function getOrganisationIdentifier(string _reference) public view returns (string);
}
