pragma solidity ^0.4.24;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

  enum DocAttributes {
    TITLE,
    FORMAT,
    URL,
    CATEGORY,
    COUNTRYCODE,
    DESC,
    LANG,
    DATE
  }

  uint8 constant numDocAttributes = 8;

  struct Report {
    string version;
  }

  struct Organisation {
    string orgRef;
    string defaultLang;
    string defaultCurrency;
  }

  struct ReportingOrganisation {
    string reportingOrgRef;
    string orgType;
    bool isSecondary;
  }

  struct Document {
    bytes32 title;
    bytes32 format;
    bytes32 url;
    bytes32 category;
    bytes32 countryCode;
    bytes32 desc;
    bytes32 lang;
    bytes32 date;
  }

  string[] reportReferences;
  mapping(string => Report) private reports;
  mapping(string => Organisation) private organisations;
  mapping(string => ReportingOrganisation) private reportingOrgs;

  mapping(string => string[]) private reportDocReferences;
  mapping(string => mapping(string => Document)) private docs;

  event SetReport(string _reference, string _version);
  event SetOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency);
  event SetReportingOrganisation(string _reference, string _reportingOrgRef, string _orgType, bool _isSecondary);
  event SetDocument(string _reference, string _docRef);

  function setReport(string _reference, string _version) public {
    require((bytes(_reference).length > 0) && (bytes(_version).length > 0));

    reports[_reference].version = _version;
    if(!getReportExists(_reference)) {
      reportReferences.push(_reference);
    }
  }

  function setOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public {
    require((bytes(_reference).length > 0) && (bytes(_orgRef).length > 0) && (bytes(_defaultLang).length > 0) && (bytes(_defaultCurrency).length > 0));

    organisations[_reference].orgRef = _orgRef;
    organisations[_reference].defaultLang = _defaultLang;
    organisations[_reference].defaultCurrency = _defaultCurrency;

    emit SetOrganisation(_reference, _orgRef, _defaultLang, _defaultCurrency);
  }

  function setReportingOrganisation(string _reference, string _reportingOrgRef, string _orgType, bool _isSecondary) public {
    require((bytes(_reference).length > 0) && (bytes(_reportingOrgRef).length > 0) && (bytes(_orgType).length > 0));

    reportingOrgs[_reference].reportingOrgRef = _reportingOrgRef;
    reportingOrgs[_reference].orgType = _orgType;
    reportingOrgs[_reference].isSecondary = _isSecondary;

    emit SetReportingOrganisation(_reference, _reportingOrgRef, _orgType, _isSecondary);
  }

  function setDocument(string _reference, string _docRef, bytes32[] _attributes) public {
    require((bytes(_reference).length > 0) &&
            (bytes(_docRef).length > 0) &&
            (_attributes.length == numDocAttributes));

    docs[_reference][_docRef].title = _attributes[uint256(DocAttributes.TITLE)];
    docs[_reference][_docRef].format = _attributes[uint256(DocAttributes.FORMAT)];
    docs[_reference][_docRef].url = _attributes[uint256(DocAttributes.URL)];
    docs[_reference][_docRef].category = _attributes[uint256(DocAttributes.CATEGORY)];
    docs[_reference][_docRef].countryCode = _attributes[uint256(DocAttributes.COUNTRYCODE)];
    docs[_reference][_docRef].desc = _attributes[uint256(DocAttributes.DESC)];
    docs[_reference][_docRef].lang = _attributes[uint256(DocAttributes.LANG)];
    docs[_reference][_docRef].date = _attributes[uint256(DocAttributes.DATE)];

    if(!getReportDocExists(_reference, _docRef)) {
      reportDocReferences[_reference].push(_docRef);
    }

    emit SetDocument(_reference, _docRef);
  }

  function getReportExists(string _reference) public view returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_reference, reportReferences);
    return index != reportReferences.length;
  }

  function getReportDocExists(string _reference, string _docRef) public view returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_docRef, reportDocReferences[_reference]);
    return index != reportDocReferences[_reference].length;
  }

  function getVersion(string _reference) public view returns (string) {
    return reports[_reference].version;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumReportDocs(string _reference) public view returns (uint256) {
    return reportDocReferences[_reference].length;
  }

  function getReportReference(uint256 _index) public view returns (string) {
    require(_index < reportReferences.length);
    return reportReferences[_index];
  }

  function getReportDocReference(string _reference, uint256 _index) public view returns (string) {
    require(_index < reportReferences.length);
    return reportDocReferences[_reference][_index];
  }

  function getOrganisation(string _reference) public view returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].orgRef;
  }

  function getOrganisationDefaultLang(string _reference) public view returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].defaultLang;
  }

  function getOrganisationDefaultCurrency(string _reference) public view returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].defaultCurrency;
  }

  function getReportingOrganisation(string _reference) public view returns (string) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].reportingOrgRef;
  }

  function getReportingOrganisationType(string _reference) public view returns (string) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].orgType;
  }

  function getReportingOrganisationIsSecondary(string _reference) public view returns (bool) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].isSecondary;
  }

  function getDocumentTitle(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].title;
  }

  function getDocumentFormat(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].format;
  }

  function getDocumentURL(string _reference, string _docRef) public view returns (bytes32){
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].url;
  }

  function getDocumentCategory(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].category;
  }

  function getDocumentCountry(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].countryCode;
  }

  function getDocumentDescription(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].desc;
  }

  function getDocumentLang(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].lang;
  }

  function getDocumentDate(string _reference, string _docRef) public view returns (bytes32) {
    require(bytes(_reference).length > 0);
    return docs[_reference][_docRef].date;
  }
}
