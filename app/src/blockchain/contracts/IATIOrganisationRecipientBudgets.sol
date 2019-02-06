pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationRecipientBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationRecipientBudgets is OrganisationRecipientBudgets {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private budgetRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => RecipientBudget))) private budgets;

  event SetRecipientBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, RecipientBudget _budget);

  function setRecipientBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, RecipientBudget memory _budget) public {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.recipientOrgRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_organisationsRef][_orgRef][_budgetRef] = _budget;

    if(!Strings.getExists(_budgetRef, budgetRefs[_organisationsRef][_orgRef])) {
      budgetRefs[_organisationsRef][_orgRef].push(_budgetRef);
    }

    emit SetRecipientBudget(_organisationsRef, _orgRef, _budgetRef, _budget);
  }

  function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return budgetRefs[_organisationsRef][_orgRef].length;
  }

  function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _index < budgetRefs[_organisationsRef][_orgRef].length);

    return budgetRefs[_organisationsRef][_orgRef][_index];
  }

  function getRecipientBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (RecipientBudget memory) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef];
  }

  function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].recipientOrgRef;
  }

  function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].budgetLine;
  }

  function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.value;
  }

  function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.status;
  }

  function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.start;
  }

  function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.end;
  }
}
