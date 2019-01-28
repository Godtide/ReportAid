pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportCountryBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportCountryBudgets is OrganisationReportCountryBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private countryBudgetReferences;
  mapping(bytes32 => mapping(bytes32 => CountryBudget)) private countryBudgets;

  event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, CountryBudget _budget);

  function setCountryBudget(CountryBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.countryRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    countryBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getCountryBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      countryBudgetReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetCountryBudget(_budget.reportRef, _budget.budgetRef, _budget);
  }

  function getReportExists(bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences);
      exists = (index != reportReferences.length);
    }
    return exists;
  }

  function getCountryBudgetExists(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    bool exists = false;
    if ( !(countryBudgetReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_countryBudgetRef, countryBudgetReferences[_reportRef]);
      exists = (index != countryBudgetReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumCountryBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return countryBudgetReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < countryBudgetReferences[_reportRef].length);

    return countryBudgetReferences[_reportRef][_index];
  }

  function getCountryBudget(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (CountryBudget memory) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef];
  }


  function getCountryBudgetCountry(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].countryRef;
  }

  function getCountryBudgetLine(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].budgetLine;
  }

  function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].finance.value;
  }

  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].finance.status;
  }

  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].finance.start;
  }

  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryBudgetRef[0] != 0);

    return countryBudgets[_reportRef][_countryBudgetRef].finance.end;
  }
}
