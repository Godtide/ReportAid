pragma solidity ^0.4.24;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReports.sol";
import "./Strings.sol";

contract IATIOrganisationReports is OrganisationReports {

  string version;

  string[] orgReferences;


  function getVersion() public constant returns (string) {
    return version;
  }

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
  function getOrganisationExists(string _orgRef) public constant returns (bool);
  function getOrganisationName(string _orgRef) public constant returns (string);
  function getOrganisationDefaultLang(string _orgRef) public constant returns (string);
  function getOrganisationDefaultCurrency(string _orgRef) public constant returns (string);

  function getReportingOrganisation(string _orgRef) public constant returns (string);
  function getReportingOrganisationType(string _orgRef) public constant returns (string);
  function getIsSecondary(string _orgRef) public constant returns (bool);

  function getNumBudgets(string _orgRef) public constant returns (uint256);
  function getBudgetReference(string _orgRef, uint256 _index) public constant returns (string);
  function getTotalBudget(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetStatus(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetStart(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetEnd(string _orgRef, string _budgetRef) public constant returns (uint256);

  function getNumRecipientOrgs(string _orgRef) public constant returns (uint256);
  function getRecipientOrgReference(string _orgRef, uint256 _index) public constant returns (string);
  function getRecipientOrgBudget(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetStatus(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetStart(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetEnd(string _orgRef, string _recipientOrgRef) public constant returns (uint256);

  function getNumRecipientRegions(string _orgRef) public constant returns (uint256);
  function getRecipientRegionsReference(string _orgRef, uint256 _index) public constant returns (string);
  function getRecipientRegionsBudget(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRecipientRegionsBudgetStatus(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRecipientRegionsBudgetStart(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRecipientRegionsBudgetEnd(string _orgRef, string _regionRef) public constant returns (uint256);

  function getNumRecipientCountries(string _orgRef) public constant returns (uint256);
  function getRecipientCountryReference(string _orgRef, uint256 _index) public constant returns (string);
  function getRecipientCountryBudget(string _orgRef, string _countryRef) public constant returns (uint256);
  function getRecipientCountryBudgetStatus(string _orgRef, string _countryRef) public constant returns (uint256);
  function getRecipientCountryBudgetStart(string _orgRef, string _countryRef) public constant returns (uint256);
  function getRecipientCountryBudgetEnd(string _orgRef, string _countryRef) public constant returns (uint256);

  function getNumExpenditures(string _orgRef) public constant returns (uint256);
  function getExpenditureReference(string _orgRef, uint256 _index) public constant returns (string);
  function getTotalExpenditure(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureStatus(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureStart(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureEnd(string _orgRef, string expenditureRef) public constant returns (uint256);

  function getNumDocuments(string _orgRef) public constant returns (uint256);
  function getDocumentReference(string _orgRef, uint256 _index) public constant returns (string);
  function getDocumentTitle(string _orgRef, string _docRef) public constant returns (string);
  function getDocumentCountry(string _orgRef, string _docRef) public constant returns (string);
  function getDocumentDescription(string _orgRef, string _docRef) public constant returns (string);
  function getDocumentCategory(string _orgRef, string _docRef) public constant returns (string);
  function getDocumentLang(string _orgRef, string _docRef) public constant returns (string);
  function getDocumentDate(string _orgRef, string _docRef) public constant returns (string);
}
