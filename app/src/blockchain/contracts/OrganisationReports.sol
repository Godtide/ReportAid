pragma solidity ^0.4.24;

contract OrganisationReports {

  function setReport(string _reference, string _version) public;
  function setOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public;
  function setReportingOrganisation(string _reference, string _reportingOrgRef, string _type, bool _isSecondary) public;
  function setDocument(string _reference, string _title, string _countryRef, string _desc, string _category, string _lang, string _date) public;

  function getReportExists(string _reference) public constant returns (bool);
  function getNumReports() public constant returns (uint256);
  function getReportReference(uint256 _index) public constant returns (string);
  function getVersion(string _reference) public constant returns (string);

  function getOrganisation(string _reference) public constant returns (string);
  function getOrganisationDefaultLang(string _reference) public constant returns (string);
  function getOrganisationDefaultCurrency(string _reference) public constant returns (string);

  function getReportingOrganisation(string _reference) public constant returns (string);
  function getReportingOrganisationType(string _reference) public constant returns (string);
  function getReportingOrganisationIsSecondary(string _reference) public constant returns (bool);

  function getDocumentTitle(string _reference) public constant returns (string);
  function getDocumentCountry(string _reference) public constant returns (string);
  function getDocumentDescription(string _reference) public constant returns (string);
  function getDocumentCategory(string _reference) public constant returns (string);
  function getDocumentLang(string _reference) public constant returns (string);
  function getDocumentDate(string _reference) public constant returns (string);
}
