pragma solidity ^0.4.24;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

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
    string title;
    string countryRef;
    string desc;
    string category;
    string lang;
    string date;
  }

  string[] reportReferences;
  mapping(string => Report) private reports;
  mapping(string => Organisation) private organisations;
  mapping(string => ReportingOrganisation) private reportingOrgs;
  mapping(string => Document) private docs;

  event SetReport(string _reference, string _version);
  event SetOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency);
  event SetReportingOrganisation(string _reference, string _reportingOrgRef, string _orgType, bool _isSecondary);
  event SetDocument(string _reference, string _title, string _countryRef, string _desc, string _category, string _lang, string _date);

  constructor() public {
  }

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

  function setDocument(string _reference, string _title, string _countryRef, string _desc, string _category, string _lang, string _date) public {
    require((bytes(_reference).length > 0) &&
            (bytes(_title).length > 0) &&
            (bytes(_desc).length > 0) &&
            (bytes(_category).length > 0)  &&
            (bytes(_lang).length > 0)  &&
            (bytes(_date).length > 0));

    docs[_reference].title = _title;
    docs[_reference].countryRef = _countryRef;
    docs[_reference].desc = _desc;
    docs[_reference].category = _category;
    docs[_reference].lang = _lang;
    docs[_reference].date = _date;

    emit SetDocument(_reference, _title, _countryRef, _desc, _category, _lang, _date);
  }

  function getReportExists(string _reference) public constant returns (bool) {
    require(bytes(_reference).length > 0);

    uint256 index = Strings.getIndex(_reference, reportReferences);
    return index != reportReferences.length;
  }

  function getNumReports() public constant returns (uint256) {
    return reportReferences.length;
  }

  function getReportReference(uint256 _index) public constant returns (string) {
    require(_index < reportReferences.length);
    return reportReferences[_index];
  }

  function getVersion(string _reference) public constant returns (string) {
    return reports[_reference].version;
  }

  function getOrganisation(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].orgRef;
  }

  function getOrganisationDefaultLang(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].defaultLang;
  }

  function getOrganisationDefaultCurrency(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return organisations[_reference].defaultCurrency;
  }

  function getReportingOrganisation(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].reportingOrgRef;
  }

  function getReportingOrganisationType(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].orgType;
  }

  function getReportingOrganisationIsSecondary(string _reference) public constant returns (bool) {
    require(bytes(_reference).length > 0);
    return reportingOrgs[_reference].isSecondary;
  }

  function getDocumentTitle(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].title;
  }
  function getDocumentCountry(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].countryRef;
  }
  function getDocumentDescription(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].desc;
  }
  function getDocumentCategory(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].category;
  }
  function getDocumentLang(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].lang;
  }
  function getDocumentDate(string _reference) public constant returns (string) {
    require(bytes(_reference).length > 0);
    return docs[_reference].date;
  }
}
