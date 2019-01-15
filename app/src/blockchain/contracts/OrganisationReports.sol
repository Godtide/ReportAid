pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReports {

  struct ReportingOrganisation {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct Report {
    bytes32 version;
    bytes32 orgRef;
    bytes32 reportRef;
    ReportingOrganisation reportingOrg;
    bytes32 lang;
    bytes32 currency;
    bytes32 lastUpdatedTime;
  }

  function setReport(Report memory _report) public;

  function getOrgExists(bytes32 _orgRef) public view returns (bool);
  function getReportExists(bytes32 _orgRef, bytes32 _reportRef) public view returns (bool);

  function getNumOrgs() public view returns (uint256);
  function getNumReports(bytes32 _orgRef) public view returns (uint256);

  function getOrganisationReference(uint256 _index) public view returns (bytes32);
  function getReportReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32);

  function getReport(bytes32 _orgRef, bytes32 _reportRef) public view returns (Report memory);

  function getLang(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32);
  function getCurrency(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32);
  function getVersion(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32);
  function getReportingOrg(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32);
  function getReportingOrgType(bytes32 _orgRef, bytes32 _reportRef) public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _orgRef, bytes32 _reportRef) public view returns (bool);
}
