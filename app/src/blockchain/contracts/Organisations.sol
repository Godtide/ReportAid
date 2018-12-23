pragma solidity ^0.5.0;

contract Organisations {

  function setOrganisation(bytes32 _reference, bytes32 _name, bytes32 _identifier) public;

  function getOrganisationExists(bytes32 _reference) public view returns (bool);
  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (bytes32);
  function getOrganisationName(bytes32 _reference) public view returns (bytes32);
  function getOrganisationIdentifier(bytes32 _reference) public view returns (bytes32);
}
