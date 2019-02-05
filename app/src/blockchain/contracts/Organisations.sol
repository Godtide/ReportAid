pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Organisations {

  struct Orgs {
    bytes32 _organisationsRef;
    bytes32 version;
    bytes32 generatedTime;
  }

  function setOrganisations(Orgs memory _organisations) public;

  function getNumOrganisations() public view returns (uint256);
  function getOrganisationsReference(uint256 _index) public view returns (bytes32);

  function getOrganisations(bytes32 _organisationsRef) public view returns (Orgs memory);

  function getVersion(bytes32 _organisationsRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _organisationsRef) public view returns (bytes32);
}
