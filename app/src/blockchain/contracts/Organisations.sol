pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Organisations {

  struct Organisation {
    bytes32 reference;
    string name;
    string identifier;
  }

  function setOrganisation(Organisation _org) public;

  function getOrganisationExists(bytes32 _reference) public view returns (bool);
  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (bytes32);
  function getOrganisation(bytes32 _reference) public view returns (Organisation);
  function getOrganisationName(bytes32 _reference) public view returns (string memory);
  function getOrganisationIdentifier(bytes32 _reference) public view returns (string memory);
}
