pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

  uint8 constant numDocAttributes = 8;
  bytes32 constant defaultCurrency = 'GBP';
  bytes32 constant defaultLang = 'en-GB';
  uint8 constant defaultOrgType = 10;
  bool constant defaultIsSecondary = false;

  bytes32[] orgReferences;
  mapping(bytes32 => bytes32[]) private reportReferences;
  mapping(bytes32 =>  mapping(bytes32 => Report)) private organisationReports;

  mapping(bytes32 => bytes32[]) private reportDocReferences;
  mapping(bytes32 => mapping(bytes32 => Document)) private docs;

  event SetReport(bytes32 _orgRef, bytes32 _reportRef, Report _report);
  event SetDocument(bytes32 _reportRef, bytes32 _docRef, Document _document);

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

  function setDocument(Document memory _document) public {
    require (_document.reportRef[0] != 0 &&
             _document.docRef[0] != 0 &&
             bytes(_document.title).length > 0 &&
             _document.format[0] != 0 &&
             bytes(_document.url).length > 0 &&
             _document.category[0] != 0 &&
             _document.countryCode[0] != 0 &&
             bytes(_document.desc).length > 0 &&
             _document.lang[0] != 0 &&
             _document.date[0] != 0);

    docs[_document.reportRef][_document.docRef] = _document;

    if(!getReportDocExists(_document.reportRef, _document.docRef)) {
      reportDocReferences[_document.reportRef].push(_document.docRef);
    }

    emit SetDocument(_document.reportRef, _document.docRef, _document);
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

  function getReportDocExists(bytes32 _reportRef, bytes32 _docRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportDocReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_docRef, reportDocReferences[_reportRef]);
      exists = (index != reportDocReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumOrgs() public view returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public view returns (bytes32) {
    return orgReferences[_index];
  }

  function getNumReports(bytes32 _orgRef) public view returns (uint256) {
    require (_orgRef[0] != 0);

    return reportReferences[_orgRef].length;
  }

  function getNumReportDocs(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return reportDocReferences[_reportRef].length;
  }

  function getReportReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_orgRef[0] != 0);

    return reportReferences[_orgRef][_index];
  }

  function getReportDocReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0);

    return reportDocReferences[_reportRef][_index];
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

  function getDocument(bytes32 _reportRef, bytes32 _docRef) public view returns (Document memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef];
  }

  function getDocumentTitle(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].title;
  }

  function getDocumentFormat(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].format;
  }

  function getDocumentURL(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory){
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].url;
  }

  function getDocumentCategory(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].category;
  }

  function getDocumentCountry(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].countryCode;
  }

  function getDocumentDescription(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].desc;
  }

  function getDocumentLang(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].lang;
  }

  function getDocumentDate(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
  require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].date;
  }
}
