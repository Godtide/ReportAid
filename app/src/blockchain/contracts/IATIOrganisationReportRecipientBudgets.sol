pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportRecipientBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportRecipientBudgets is OrganisationReportRecipientBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private recipientOrgReferences;
  mapping(bytes32 => mapping(bytes32 => RecipientBudget)) private recipientOrgBudgets;

  event SetRecipientOrgBudget(bytes32 _reportRef, bytes32 _budgetRef, RecipientBudget _budget);

  function setRecipientOrgBudget(RecipientBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.orgRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    recipientOrgBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getRecipientOrgBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      recipientOrgReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetRecipientOrgBudget(_budget.reportRef, _budget.budgetRef, _budget);
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

  function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(recipientOrgReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, recipientOrgReferences[_reportRef]);
      exists = (index != recipientOrgReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumRecipientOrgBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return recipientOrgReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < recipientOrgReferences[_reportRef].length);

    return recipientOrgReferences[_reportRef][_index];
  }

  function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (RecipientBudget memory) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef];
  }

  function getRecipientOrgBudgetLine(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef].budgetLine;
  }

  function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef].finance.value;
  }

  function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef].finance.status;
  }

  function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef].finance.start;
  }

  function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _recipientOrgRef[0] != 0);

    return recipientOrgBudgets[_reportRef][_recipientOrgRef].finance.end;
  }
}
