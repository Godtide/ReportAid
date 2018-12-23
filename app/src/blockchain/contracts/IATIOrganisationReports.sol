pragma solidity ^0.5.0;

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
  bytes32 constant defaultCurrency = 'GBP';
  bytes32 constant defaultLang = 'en-GB';
  uint8 constant defaultOrgType = 10;
  bool constant defaultIsSecondary = false;

  struct Report {
    ReportingOrganisation reportingOrg;
    bytes32 version;
    bytes32 lang;
    bytes32 currency;
    bytes32 generatedTime;
    bytes32 lastUpdatedTime;
  }

  struct ReportingOrganisation {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct Document {
    bytes32 title;
    bytes32 format;
    bytes32 url;
    bytes32 category;
    bytes32 countryCode;
    string desc;
    bytes32 lang;
    bytes32 date;
  }

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private orgReferences;
  mapping(bytes32 => mapping(bytes32 => Report)) private organisationReports;

  mapping(bytes32 => bytes32[]) private reportDocReferences;
  mapping(bytes32 => mapping(bytes32 => Document)) private docs;

  event SetReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime);
  event SetDefaults(bytes32 _reference, bytes32 _orgRef, bytes32 _defaultLang, bytes32 _defaultCurrency);
  event SetReportingOrgType(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, uint8 _type, bool _isSecondary);
  event SetAssociatedDocument(bytes32 _reference, bytes32 _docRef);
  event SetDocumentDescription(bytes32 _reference, bytes32 _docRef);

  function setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime) public {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0 &&
            _reportingOrgRef[0] != 0 &&
            _version[0] != 0 &&
            _generatedTime[0] != 0);

    organisationReports[_reference][_orgRef].version = _version;
    organisationReports[_reference][_orgRef].lang = defaultLang;
    organisationReports[_reference][_orgRef].currency = defaultCurrency;
    organisationReports[_reference][_orgRef].generatedTime = _generatedTime;
    organisationReports[_reference][_orgRef].lastUpdatedTime = _generatedTime;
    organisationReports[_reference][_orgRef].reportingOrg.orgRef = _reportingOrgRef;
    organisationReports[_reference][_orgRef].reportingOrg.orgType = defaultOrgType;
    organisationReports[_reference][_orgRef].reportingOrg.isSecondary = defaultIsSecondary;

    if (!getReportExists(_reference)) {
      reportReferences.push(_reference);
    }
    if(!getReportOrgExists(_reference, _orgRef)) {
      orgReferences[_reference].push(_orgRef);
    }

    emit SetReport(_reference, _orgRef, _reportingOrgRef, _version, _generatedTime);
  }

  function setDefaults(bytes32 _reference, bytes32 _orgRef, bytes32 _defaultLang, bytes32 _defaultCurrency) public {
    require (_reference[0] != 0 && _defaultLang[0] != 0 && _defaultCurrency[0] != 0);

    organisationReports[_reference][_orgRef].lang = _defaultLang;
    organisationReports[_reference][_orgRef].currency = _defaultCurrency;

    emit SetDefaults(_reference, _orgRef, _defaultLang, _defaultCurrency);
  }

  function setReportingOrgType(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, uint8 _type, bool _isSecondary) public {
    require (_reference[0] != 0 &&
            _orgRef[0] != 0 &&
            _reportingOrgRef[0] != 0 &&
            Strings.equal(organisationReports[_reference][_orgRef].reportingOrg.orgRef, _reportingOrgRef));

    organisationReports[_reference][_orgRef].reportingOrg.orgType = _type;
    organisationReports[_reference][_orgRef].reportingOrg.isSecondary = _isSecondary;

    emit SetReportingOrgType(_reference, _orgRef, _reportingOrgRef, _type, _isSecondary);
  }

  function setAssociatedDocument(bytes32 _reference, bytes32 _docRef, bytes32[] memory _attributes) public {
    require (_reference[0] != 0 && _docRef[0] != 0 && _attributes.length == numDocAttributes - 1);

    docs[_reference][_docRef].title = _attributes[uint256(DocAttributes.TITLE)];
    docs[_reference][_docRef].format = _attributes[uint256(DocAttributes.FORMAT)];
    docs[_reference][_docRef].url = _attributes[uint256(DocAttributes.URL)];
    docs[_reference][_docRef].category = _attributes[uint256(DocAttributes.CATEGORY)];
    docs[_reference][_docRef].countryCode = _attributes[uint256(DocAttributes.COUNTRYCODE)];
    docs[_reference][_docRef].lang = _attributes[uint256(DocAttributes.LANG)];
    docs[_reference][_docRef].date = _attributes[uint256(DocAttributes.DATE)];

    if(!getReportDocExists(_reference, _docRef)) {
      reportDocReferences[_reference].push(_docRef);
    }

    emit SetAssociatedDocument(_reference, _docRef);
  }


  function setDocumentDescription(bytes32 _reference, bytes32 _docRef, string memory _description) public  {
    require (_reference[0] != 0 && _docRef[0] != 0);

    docs[_reference][_docRef].desc = _description;

    if(!getReportDocExists(_reference, _docRef)) {
      reportDocReferences[_reference].push(_docRef);
    }

    emit SetDocumentDescription(_reference, _docRef);
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
