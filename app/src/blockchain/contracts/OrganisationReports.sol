pragma solidity ^0.4.24;

contract OrganisationReports {

  function setVersion(string _version) public;
  function setOrganisation(string _orgRef, string _name, string _defaultLang, string _defaultCurrency);
  function setReportingOrganisation(string _orgRef, string _reportingOrgRef, string _type, bool _isSecondary);
  function setTotalBudget(string _orgRef, uint256 _value, string _status, string _start, string _end);
  function setRecipientOrgBudget(string _orgRef, string _recipientOrgRef, uint256 _value, string _status, string _start, string _end);
  function setRecipientRegionBudget(string _orgRef, string _regionRef, uint256 _value, string _status, string _start, string _end);
  function setRecipientCountryBudget(string _orgRef, string _countryRef, uint256 _value, string _status, string _start, string _end);
  function setTotalExpenditure(string _orgRef, uint256 _value, string _status, string _start, string _end);
  function setDocument(string _orgRef, string _title, string _countryRef, string _desc, string _category, _string _lang, string _date);

  function getVersion() public constant returns (string);

  function getNumOrganisations() public constant returns (uint256);
  function getOrganisationReference(uint256 _index) public constant returns (string);
  function getOrganisationExists(string _reference) public constant returns (bool);
  function getOrganisationName(string _reference) public constant returns (string);
  function getOrganisationDefaultLang(string _reference) public constant returns (string);
  function getOrganisationDefaultCurrency(string _reference) public constant returns (string);

  function  getReportingOrganisation(string _reference) public constant returns (string);
  function  getReportingOrganisationType(string _reference) public constant returns (string);
  function  getIsSecondary(string _reference) public constant returns (bool);

  function getTotalBudget(string _reference) public constant returns (uint256);
  function getTotalBudgetStatus(string _reference) public constant returns (uint256);
  function getTotalBudgetStart(string _reference) public constant returns (uint256);
  function getTotalBudgetEnd(string _reference) public constant returns (uint256);

  function getNumRecipientOrgs(string _reference) public constant returns (uint256);
  function getRecipientOrgReference(uint256 _index) public constant returns (string);
  function getRecipientOrgBudget(string _reference) public constant returns (uint256);
  function getRecipientOrgBudgetStatus(string _reference) public constant returns (uint256);
  function getRecipientOrgBudgetStart(string _reference) public constant returns (uint256);
  function getRecipientOrgBudgetEnd(string _reference) public constant returns (uint256);

  function getNumRecipientRegions(string _reference) public constant returns (uint256);
  function getRecipientRegionsReference(uint256 _index) public constant returns (string);
  function getRecipientRegionsBudget(string _reference) public constant returns (uint256);
  function getRecipientRegionsBudgetStatus(string _reference) public constant returns (uint256);
  function getRecipientRegionsBudgetStart(string _reference) public constant returns (uint256);
  function getRecipientRegionsBudgetEnd(string _reference) public constant returns (uint256);

  function getNumRecipientCountries(string _reference) public constant returns (uint256);
  function getRecipientCountryReference(uint256 _index) public constant returns (string);
  function getRecipientCountryBudget(string _reference) public constant returns (uint256);
  function getRecipientCountryBudgetStatus(string _reference) public constant returns (uint256);
  function getRecipientCountryBudgetStart(string _reference) public constant returns (uint256);
  function getRecipientCountryBudgetEnd(string _reference) public constant returns (uint256);

  function getTotalExpenditure(string _reference) public constant returns (uint256);
  function getTotalExpenditureStatus(string _reference) public constant returns (uint256);
  function getTotalExpenditureStart(string _reference) public constant returns (uint256);
  function getTotalExpenditureEnd(string _reference) public constant returns (uint256);

  function getDocumentTitle(string _reference) public constant returns (string);
  function getDocumentCountry(string _reference) public constant returns (string);
  function getDocumentDescription(string _reference) public constant returns (string);
  function getDocumentCategory(string _reference) public constant returns (string);
  function getDocumentLang(string _reference) public constant returns (string);
  function getDocumentDate(string _reference) public constant returns (string);
}
