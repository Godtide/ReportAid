pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./Budgets.sol";
import "./Strings.sol";

contract IATIBudgets is Budgets {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private budgetRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Budget))) private budgets;

  event SetBudget(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef, Budget _budget);

  function setBudget(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef, Budget memory _budget) public {
    require (_topLevelRef[0] != 0 &&
             _secondLevelRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > uint8(Status.NONE) &&
             _budget.finance.status < uint8(Status.MAX) &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_topLevelRef][_secondLevelRef][_budgetRef] = _budget;

    if(!Strings.getExists(_budgetRef, budgetRefs[_topLevelRef][_secondLevelRef])) {
      budgetRefs[_topLevelRef][_secondLevelRef].push(_budgetRef);
    }

    emit SetBudget(_topLevelRef, _secondLevelRef, _budgetRef, _budget);
  }

  function getNumBudgets(bytes32 _topLevelRef, bytes32 _secondLevelRef) public view returns (uint256) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0);

    return budgetRefs[_topLevelRef][_secondLevelRef].length;
  }

  function getBudgetReference(bytes32 _topLevelRef, bytes32 _secondLevelRef, uint256 _index) public view returns (bytes32) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _index < budgetRefs[_topLevelRef][_secondLevelRef].length);

    return budgetRefs[_topLevelRef][_secondLevelRef][_index];
  }

  function getBudget(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (Budget memory) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef];
  }

  function getBudgetLine(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef].budgetLine;
  }

  function getBudgetValue(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef].finance.value;
  }

  function getBudgetStatus(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef].finance.status;
  }

  function getBudgetStart(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef].finance.start;
  }
  function getBudgetEnd(bytes32 _topLevelRef, bytes32 _secondLevelRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_topLevelRef[0] != 0 && _secondLevelRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_topLevelRef][_secondLevelRef][_budgetRef].finance.end;
  }
}
