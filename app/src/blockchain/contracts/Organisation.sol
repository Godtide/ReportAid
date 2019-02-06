pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Organisation {

  struct ReportingOrg {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct Org {
    ReportingOrg reportingOrg;
    bytes32 lang;
    bytes32 currency;
    bytes32 lastUpdatedTime;
  }

  function setOrganisation(bytes32 _organisationsRef, bytes32 _orgRef, Org memory _org) public;

  function getNumOrganisations(bytes32 _organisationsRef) public view returns (uint256);
  function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) public view returns (bytes32);

  function getOrganisation(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (Org memory);

  function getLang(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32);
  function getCurrency(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32);
  function getReportingOrg(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32);
  function getReportingOrgType(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bool);
}