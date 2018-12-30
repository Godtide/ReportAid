pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Organisations {

  struct Organisation {
    bytes32 orgRef;
    string name;
    string identifier;
  }

  function setOrganisation(Organisation memory _org) public;

  function getOrganisationExists(bytes32 _orgRef) public view returns (bool);
  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (bytes32);
  function getOrganisation(bytes32 _orgRef) public view returns (Organisation memory);
  function getOrganisationName(bytes32 _orgRef) public view returns (string memory);
  function getOrganisationIdentifier(bytes32 _orgRef) public view returns (string memory);
}
