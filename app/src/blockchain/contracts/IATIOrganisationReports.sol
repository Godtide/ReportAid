pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

  uint8 constant numDocAttributes = 8;
  bytes32 constant defaultCurrency = 'GBP';
  bytes32 constant defaultLang = 'en-GB';
  uint8 constant defaultOrgType = 10;
  bool constant defaultIsSecondary = false;

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private orgReferences;
  mapping(bytes32 => mapping(bytes32 => Report)) private organisationReports;

  mapping(bytes32 => bytes32[]) private reportDocReferences;
  mapping(bytes32 => mapping(bytes32 => Document)) private docs;

  event SetReport(bytes32 _reference,  Report _report);
  event SetDocument(bytes32 _reportRef, bytes32 _docRef, Document _document);

  function setReport(Report _report) public {
    require (_report.reference[0] != 0 &&
             _report.reportingOrg.orgRef[0] != 0 &&
             _report.reportingOrg.orgType > 0 &&
             _report.version[0] != 0 &&
             _report.lang[0] != 0 &&
             _report.currency[0] != 0 &&
             _report.generatedTime[0] != 0 &&
             _report.lastUpdatedTime[0] != 0 );

    organisationReports[_report._reference][_report.reportingOrg.orgRef] = _report;

    if (!getReportExists(_report._reference)) {
      reportReferences.push(_report._reference);
    }
    if(!getReportOrgExists(_report._reference, _report.reportingOrg.orgRef)) {
      orgReferences[_report._reference].push(_report.reportingOrg.orgRef);
    }

    emit SetReport(_report.reference, _report);
  }

  function setDocument(Document _document) public {
    require (_document.reportRef[0] != 0 &&
             _document.docRef[0] != 0 &&
             bytes(_document.title).length > 0 &&
             _document.format[0] != 0 &&
             bytes(_document.url).length > 0 &&
             _document.category[0] != 0 &&
             _document.countryCode[0] != 0 &&
             bytes(_document.desc).length > 0 &&
             _document.lang[0] != 0 &&
             bytes(_document.date).length > 0);

    docs[_document.reportRef][_document.docRef] = _document;

    if(!getReportDocExists(_document.reportRef, _document.docRef)) {
      reportDocReferences[_document.reportRef].push(_document.docRef);
    }

    emit SetDocument(_document.reportRef, _document.docRef, _document);
  }

  function getReportExists(bytes32 _reference) public view returns (bool) {
    require (_reference[0] != 0);

    uint256 index = Strings.getIndex(_reference, reportReferences);
    return index != reportReferences.length;
  }

  function getReportOrgExists(bytes32 _reference, bytes32 _orgRef) public view returns (bool) {
    require (_reference[0] != 0 && _orgRef[0] != 0);

    uint256 index = Strings.getIndex(_orgRef, orgReferences[_reference]);
    return index != orgReferences[_reference].length;
  }

  function getReportDocExists(bytes32 _reference, bytes32 _docRef) public view returns (bool) {
    require (_reference[0] != 0);

    uint256 index = Strings.getIndex(_docRef, reportDocReferences[_reference]);
    return index != reportDocReferences[_reference].length;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumReportOrgs(bytes32 _reference) public view returns (uint256) {
    require (_reference[0] != 0);
    return orgReferences[_reference].length;
  }

  function getNumReportDocs(bytes32 _reference) public view returns (uint256) {
    require (_reference[0] != 0);
    return reportDocReferences[_reference].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);
    return reportReferences[_index];
  }

  function getReportOrgReference(bytes32 _reference, uint256 _index) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _index < orgReferences[_reference].length);
    return orgReferences[_reference][_index];
  }

  function getReportDocReference(bytes32 _reference, uint256 _index) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _index < reportReferences.length);
    return reportDocReferences[_reference][_index];
  }

  function getReportingOrg(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].reportingOrg.orgRef;
  }

  function getLang(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].lang;
  }

  function getCurrency(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].currency;
  }

  function getVersion(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].version;
  }

  function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].generatedTime;
  }

  function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].lastUpdatedTime;
  }

  function getReportingOrgType(bytes32 _reference,  bytes32 _orgRef) public view returns (uint8) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].reportingOrg.orgType;
  }

  function getReportingOrgIsSecondary(bytes32 _reference,  bytes32 _orgRef) public view returns (bool) {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0);
    return organisationReports[_reference][_orgRef].reportingOrg.isSecondary;
  }

  function getDocumentTitle(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].title;
  }

  function getDocumentFormat(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].format;
  }

  function getDocumentURL(bytes32 _reference, bytes32 _docRef) public view returns (bytes32){
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].url;
  }

  function getDocumentCategory(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].category;
  }

  function getDocumentCountry(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].countryCode;
  }

  function getDocumentDescription(bytes32 _reference, bytes32 _docRef) public view returns (string memory) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].desc;
  }

  function getDocumentLang(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].lang;
  }

  function getDocumentDate(bytes32 _reference, bytes32 _docRef) public view returns (bytes32) {
    require (_reference[0] != 0 &&
            _docRef[0] != 0);
    return docs[_reference][_docRef].date;
  }
}
