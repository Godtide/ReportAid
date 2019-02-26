pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./Budgets.sol";
import "./Strings.sol";

contract IATIBudgets is Budgets {

  mapping(uint8 => mapping(bytes32 => mapping(bytes32 => bytes32[]))) private budgetRefs;
  mapping(uint8 => mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Budget)))) private budgets;

  event SetBudget(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, Budget _budget);

  function setBudget(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, Budget memory _budget) public {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > uint8(Status.NONE) &&
             _budget.finance.status < uint8(Status.MAX) &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0);

    budgets[_budgetType][_firstRef][_secondRef][_budgetRef] = _budget;

    if(!Strings.getExists(_budgetRef, budgetRefs[_budgetType][_firstRef][_secondRef])) {
      budgetRefs[_budgetType][_firstRef][_secondRef].push(_budgetRef);
    }

    emit SetBudget(_budgetType, _firstRef, _secondRef, _budgetRef, _budget);
  }

  function getNumBudgets(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef) public view returns (uint256) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0);

    return budgetRefs[_budgetType][_firstRef][_secondRef].length;
  }

  function getBudgetReference(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) public view returns (bytes32) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _index < budgetRefs[_budgetType][_firstRef][_secondRef].length);

    return budgetRefs[_budgetType][_firstRef][_secondRef][_index];
  }

  function getBudget(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (Budget memory) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef];
  }

  function getBudgetLine(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].budgetLine;
  }

  function getOtherRef(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].otherRef;
  }

  function getBudgetValue(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].finance.value;
  }

  function getBudgetStatus(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].finance.status;
  }

  function getBudgetStart(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].finance.start;
  }

  function getBudgetEnd(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_budgetType > uint8(BudgetType.NONE) &&
             _budgetType < uint8(BudgetType.MAX) &&
             _firstRef[0] != 0 &&
             _secondRef[0] != 0 &&
             _budgetRef[0] != 0);

    return budgets[_budgetType][_firstRef][_secondRef][_budgetRef].finance.end;
  }
}
