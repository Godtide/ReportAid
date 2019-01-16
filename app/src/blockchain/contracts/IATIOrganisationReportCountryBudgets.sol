pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportCountryBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportCountryBudgets is OrganisationReportCountryBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private countryReferences;
  mapping(bytes32 => mapping(bytes32 => CountryBudget)) private countryBudgets;

  event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, CountryBudget _budget);

  function setCountryBudget(CountryBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.countryRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    countryBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getCountryBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      countryReferences[_budget.reportRef].push(_budget.budgetRef);
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

  function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(countryReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, countryReferences[_reportRef]);
      exists = (index != countryReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumCountryBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return countryReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < countryReferences[_reportRef].length);

    return countryReferences[_reportRef][_index];
  }

  function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) public view returns (CountryBudget memory) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef];
  }

  function getCountryBudgetLine(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef].budgetLine;
  }

  function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef].finance.value;
  }

  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef].finance.status;
  }

  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef].finance.start;
  }

  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _countryRef[0] != 0);

    return countryBudgets[_reportRef][_countryRef].finance.end;
  }
}
