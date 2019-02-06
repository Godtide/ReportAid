pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationBudgets is OrganisationBudgets {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private budgetRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Budget))) private budgets;

  event SetBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, Budget _budget);

  function setBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, Budget memory _budget) public {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_organisationsRef][_orgRef][_budgetRef] = _budget;

    if(!Strings.getExists(_budgetRef, budgetRefs[_organisationsRef][_orgRef])) {
      budgetRefs[_organisationsRef][_orgRef].push(_budgetRef);
    }

    emit SetBudget(_organisationsRef, _orgRef, _budgetRef, _budget);
  }

  function getNumBudgets(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return budgetRefs[_organisationsRef][_orgRef].length;
  }

  function getBudgetReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _index < budgetRefs[_organisationsRef][_orgRef].length);

    return budgetRefs[_organisationsRef][_orgRef][_index];
  }

  function getBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (Budget memory) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef];
  }

  function getBudgetLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].budgetLine;
  }

  function getBudgetValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.value;
  }

  function getBudgetStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.status;
  }

  function getBudgetStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.start;
  }
  function getBudgetEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_orgRef][_budgetRef].finance.end;
  }
}
