pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Orgs {

  struct Org {
    string name;
    string identifier;
  }

  event SetOrg(bytes32 _orgRef, Org _org);

  function setOrg(bytes32 _orgRef, Org memory _org) virtual public;

  function getNumOrgs() virtual public view returns (uint256);
  function getOrgReference(uint256 _index) virtual public view returns (bytes32);

  function getOrg(bytes32 _orgRef) virtual public view returns (Org memory);

  function getOrgName(bytes32 _orgRef) virtual public view returns (string memory);
  function getOrgIdentifier(bytes32 _orgRef) virtual public view returns (string memory);
}
