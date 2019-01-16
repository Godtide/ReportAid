pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportBudgets is OrganisationReportBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private budgetReferences;
  mapping(bytes32 => mapping(bytes32 => Budget)) private budgets;

  event SetTotalBudget(bytes32 _reportRef, bytes32 _budgetRef, Budget _budget);

  function setTotalBudget(Budget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getTotalBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      budgetReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetTotalBudget(_budget.reportRef, _budget.budgetRef, _budget);
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

  function getTotalBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(budgetReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, budgetReferences[_reportRef]);
      exists = (index != budgetReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumTotalBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return budgetReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getTotalBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < budgetReferences[_reportRef].length);

    return budgetReferences[_reportRef][_index];
  }

  function getTotalBudget(bytes32 _reportRef, bytes32 _budgetRef) public view returns (Budget memory) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef];
  }

  function getTotalBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef].budgetLine;
  }

  function getTotalBudgetValue(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef].finance.value;
  }

  function getTotalBudgetStatus(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef].finance.status;
  }

  function getTotalBudgetStart(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef].finance.start;
  }
  function getTotalBudgetEnd(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_reportRef][_budgetRef].finance.end;
  }
}
