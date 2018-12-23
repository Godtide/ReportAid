pragma solidity ^0.5.0;

contract Organisations {

  function setOrganisation(string memory _reference, string memory _name, string memory _identifier) public;

  function getOrganisationExists(string memory _reference) public view returns (bool);
  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (string memory);
  function getOrganisationName(string memory _reference) public view returns (string memory);
  function getOrganisationIdentifier(string memory _reference) public view returns (string memory);
}
