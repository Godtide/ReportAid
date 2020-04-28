pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Organisations {

  struct Orgs {
    bytes32 version;
    bytes32 generatedTime;
  }

  event SetOrganisations(bytes32 _organisationsRef, Orgs _organisations);

  function setOrganisations(bytes32 _organisationsRef, Orgs memory _organisations) virtual public;

  function getNumOrganisations() virtual public view returns (uint256);
  function getOrganisationsReference(uint256 _index) virtual public view returns (bytes32);

  function getOrganisations(bytes32 _organisationsRef) virtual public view returns (Orgs memory);

  function getVersion(bytes32 _organisationsRef) virtual public view returns (bytes32);
  function getGeneratedTime(bytes32 _organisationsRef) virtual public view returns (bytes32);
}
