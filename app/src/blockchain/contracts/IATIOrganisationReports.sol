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
  string constant defaultCurrency = 'GBP';
  string constant defaultLang = 'en-gb';
  uint8 constant defaultOrgType = 10;
  string constant defaultIsSecondary = false;

  struct Report {
    string reportingOrgRef;
    mapping(string => ReportingOrganisation) reportingOrg;
    string version;
    string lang;
    string currency;
    string generatedTime;
    string lastUpdatedTime;
  }

  struct ReportingOrganisation {
    uint8 orgType;
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
  mapping(string => string[]) private orgReferences;
  mapping(string => mapping(string => Report)) private organisationReports;

  mapping(string => string[]) private reportDocReferences;
  mapping(string => mapping(string => Document)) private docs;

  event SetReport(string _reference, string _orgRef, string _reportingOrgRef, string _version, string _generatedTime);
  event SetDefaults(string _reference, string _defaultLang, string _defaultCurrency);
  event SetReportingOrgType(string _reference, string _orgRef, string _reportingOrgRef, string _type, bool _isSecondary);
  event SetAssociatedDocument(string _reference, string _docRef);

  function setReport(string _reference, string _orgRef, string _reportingOrgRef, string _version, string _generatedTime) public {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0) &&
            (bytes(_reportingOrgRef).length > 0) &&
            (bytes(_version).length > 0) &&
            (bytes(_generatedTime).length > 0));

    organisationReports[reference][_orgRef].version = _version;
    organisationReports[reference][_orgRef].lang = defaultLang;
    organisationReports[reference][_orgRef].currency = defaultCurrency;
    organisationReports[reference][_orgRef].generatedTime = _generatedTime;
    organisationReports[reference][_orgRef].lastUpdatedTime = _generatedTime;
    organisationReports[reference][_orgRef].reportingOrgRef = _reportingOrgRef;
    organisationReports[reference][_orgRef][_reportingOrgRef].orgType = defaultOrgType;
    organisationReports[reference][_orgRef][_reportingOrgRef].isSecondary = defaultIsSecondary;

    if(!getReportExists(_reference)) {
      reportReferences.push(_reference);
    }
    if(!getReportOrgExists(_reference, _orgRef)) {
      orgReferences[_reference].push(_orgRef);
    }

    emit SetReport(_reference, _orgRef, _reportingOrgRef, _version, __generatedTime);
  }

  function setDefaults(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public {
    require((bytes(_reference).length > 0) && (bytes(_defaultLang).length > 0) && (bytes(_defaultCurrency).length > 0));

    organisationReports[_reference][_orgRef].lang = _defaultLang;
    organisationReports[_reference][_orgRef].currency = _defaultCurrency;

    emit SetDefaults(_reference, _orgRef, _defaultLang, _defaultCurrency);
  }

  function setReportingOrgType(string _reference, string _orgRef, string _reportingOrgRef, string _type, bool _isSecondary) public {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0) &&
            (bytes(_reportingOrgRef).length > 0) &&
            (organisationReports[_reference][_orgRef].reportingOrgRef == _reportingOrgRef));

    organisationReports[_reference][_orgRef][_reportingOrgRef].orgType = _type;
    organisationReports[_reference][_orgRef][_reportingOrgRef].isSecondary = _isSecondary;

    emit SetReportingOrgType(_reference, _orgRef, _reportingOrgRef, _type, _isSecondary);
  }

  function setAssociatedDocument(string _reference, string _docRef, bytes32[] _attributes) public {
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

    emit SetAssociatedDocument(_reference, _docRef);
  }

  function getReportExists(string _reference) public view returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_reference, reportReferences);
    return index != reportReferences.length;
  }

  function getReportOrgExists(string _reference, string _orgRef) public view returns (bool) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));

    uint256 index = Strings.getIndex(_orgRef, orgReferences[_reference]);
    return index != orgReferences[_reference].length;
  }

  function getReportDocExists(string _reference, string _docRef) public view returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_docRef, reportDocReferences[_reference]);
    return index != reportDocReferences[_reference].length;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumReportOrgs(string _reference) public view returns (uint256) {
    require(bytes(_reference).length > 0);
    return orgReferences[_reference].length;
  }

  function getNumReportDocs(string _reference) public view returns (uint256) {
    return reportDocReferences[_reference].length;
  }

  function getReportReference(uint256 _index) public view returns (string) {
    require(_index < reportReferences.length);
    return reportReferences[_index];
  }


  function getReportOrgReference(string _reference, uint256 _index) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (_index < orgReferences[_reference].length));
    return orgReferences[_reference][_index];
  }

  function getReportDocReference(string _reference, uint256 _index) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (_index < reportReferences.length));
    return reportDocReferences[_reference][_index];
  }

  function getReportingOrg(string _reference, string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].reportingOrgRef;
  }

  function getLang(string _reference, string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].defaultLang;
  }

  function getCurrency(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].defaultCurrency;
  }

  function getVersion(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].version;
  }

  function getGeneratedTime(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].generatedTime;
  }

  function getLastUpdatedTime(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    return organisationReports[_reference][_orgRef].lastUpdatedTime;
  }

  function getReportingOrgType(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    string memory reportingOrgRef = organisationReports[_reference][_orgRef].reportingOrgRef;
    return organisationReports[_reference][_orgRef][reportingOrgRef].orgType;
  }

  function getReportingOrganisationIsSecondary(string _reference,  string _orgRef) public view returns (string) {
    require((bytes(_reference).length > 0) &&
            (bytes(_orgRef).length > 0));
    string memory reportingOrgRef = organisationReports[_reference][_orgRef].reportingOrgRef;
    return organisationReports[_reference][_orgRef][reportingOrgRef].isSecondary;
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
