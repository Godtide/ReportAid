pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

  bytes32[] orgReferences;
  mapping(bytes32 => bytes32[]) private reportReferences;
  mapping(bytes32 =>  mapping(bytes32 => Report)) private organisationReports;

  event SetReport(bytes32 _orgRef, bytes32 _reportRef, Report _report);

  function setReport(Report memory _report) public {
    require (_report.version[0] != 0 &&
             _report.orgRef[0] != 0 &&
             _report.reportRef[0] != 0 &&
             _report.reportingOrg.orgRef[0] != 0 &&
             _report.reportingOrg.orgType > 0 &&
             _report.lang[0] != 0 &&
             _report.currency[0] != 0 &&
             _report.lastUpdatedTime[0] != 0 );

    organisationReports[_report.orgRef][_report.reportRef] = _report;

    if (!getOrgExists(_report.orgRef)) {
      orgReferences.push(_report.orgRef);
    }

    if(!getReportExists(_report.orgRef, _report.reportRef)) {
      reportReferences[_report.orgRef].push(_report.reportRef);
    }

    emit SetReport(_report.orgRef, _report.reportRef, _report);
  }

  function getOrgExists(bytes32 _orgRef) public view returns (bool) {
    require (_orgRef[0] != 0);

    bool exists = false;
    if ( !(orgReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_orgRef, orgReferences);
      exists = (index != orgReferences.length);
    }
    return exists;
  }

  function getReportExists(bytes32 _orgRef, bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _orgRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences[_orgRef].length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences[_orgRef]);
      exists = (index != reportReferences[_orgRef].length);
    }
    return exists;
  }

  function getNumOrgs() public view returns (uint256) {
    return orgReferences.length;
  }

  function getNumReports(bytes32 _orgRef) public view returns (uint256) {
    require (_orgRef[0] != 0);

    return reportReferences[_orgRef].length;
  }

  function getOrganisationReference(uint256 _index) public view returns (bytes32) {
    require (_index < orgReferences.length);

    return orgReferences[_index];
  }

  function getReportReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _index < reportReferences[_orgRef].length);

    return reportReferences[_orgRef][_index];
  }

  function getReport(bytes32 _orgRef, bytes32 _reportRef) public view returns (Report memory) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef];
  }

  function getLang(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].lang;
  }

  function getCurrency(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].currency;
  }

  function getVersion(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].version;
  }

  function getLastUpdatedTime(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].lastUpdatedTime;
  }

  function getReportingOrg(bytes32 _orgRef, bytes32 _reportRef) public view returns (bytes32) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].reportingOrg.orgRef;
  }

  function getReportingOrgType(bytes32 _orgRef, bytes32 _reportRef) public view returns (uint8) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].reportingOrg.orgType;
  }

  function getReportingOrgIsSecondary(bytes32 _orgRef, bytes32 _reportRef) public view returns (bool) {
    require (_orgRef[0] != 0 && _reportRef[0] != 0);

    return organisationReports[_orgRef][_reportRef].reportingOrg.isSecondary;
  }
}
