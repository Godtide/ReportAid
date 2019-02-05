pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Orgs {

  struct Org {
    bytes32 orgRef;
    string name;
    string identifier;
  }

  function setOrg(Org memory _org) public;

  function getNumOrgs() public view returns (uint256);
  function getOrgReference(uint256 _index) public view returns (bytes32);

  function getOrg(bytes32 _orgRef) public view returns (Org memory);
  
  function getOrgName(bytes32 _orgRef) public view returns (string memory);
  function getOrgIdentifier(bytes32 _orgRef) public view returns (string memory);
}
