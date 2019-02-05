pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationRecipientBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationRecipientBudgets is OrganisationRecipientBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private recipientBudgetReferences;
  mapping(bytes32 => mapping(bytes32 => RecipientBudget)) private recipientBudgets;

  event SetRecipientBudget(RecipientBudget _budget);

  function setRecipientBudget(RecipientBudget memory _budget) public {
    require (_budget.report.reportRef[0] != 0 &&
             _budget.report.orgRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.recipientOrgRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    recipientBudgets[_budget.report.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.report.reportRef)) {
      reportReferences.push(_budget.report.reportRef);
    }

    if (!getRecipientBudgetExists(_budget.report.reportRef, _budget.budgetRef)) {
      recipientBudgetReferences[_budget.report.reportRef].push(_budget.budgetRef);
    }

    emit SetRecipientBudget(_budget);
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

  function getRecipientBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(recipientBudgetReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, recipientBudgetReferences[_reportRef]);
      exists = (index != recipientBudgetReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumRecipientBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return recipientBudgetReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getRecipientBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < recipientBudgetReferences[_reportRef].length);

    return recipientBudgetReferences[_reportRef][_index];
  }

  function getRecipientBudget(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (RecipientBudget memory) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef];
  }

  function getRecipientBudgetReportingOrg(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].report.orgRef;
  }

  function getRecipientBudgetOrg(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].recipientOrgRef;
  }

  function getRecipientBudgetLine(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].budgetLine;
  }

  function getRecipientBudgetValue(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].finance.value;
  }

  function getRecipientBudgetStatus(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].finance.status;
  }

  function getRecipientBudgetStart(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].finance.start;
  }

  function getRecipientBudgetEnd(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientBudgetRef[0] != 0);

    return recipientBudgets[_reportRef][_recipientBudgetRef].finance.end;
  }
}
