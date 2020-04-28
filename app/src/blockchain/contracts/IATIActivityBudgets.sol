pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

// IATI Activity Reports
// Steve Huckle

import "./ActivityBudgets.sol";
import "./IATIBudgets.sol";

contract IATIActivityBudgets is ActivityBudgets {

  IATIBudgets budgets;

  constructor(address _budgets) override virtual public {
    require(_budgets != address(0x0));
    budgets = IATIBudgets(_budgets);
  }

  function setBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef, Budgets.Budget memory _budget) override virtual public {
    budgets.setBudget(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef, _budget);
  }

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) override virtual public view returns (uint256) {
    return budgets.getNumBudgets(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef);
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) override virtual public view returns (bytes32) {
    return budgets.getBudgetReference(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _index);
  }

  function getBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (Budgets.Budget memory) {
    return budgets.getBudget(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }

  function getBudgetType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (uint8) {
    return budgets.getBudgetType(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }

  function getBudgetValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (uint256) {
    return budgets.getBudgetValue(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }

  function getBudgetStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (uint8) {
    return budgets.getBudgetStatus(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }

  function getBudgetStart(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (bytes32) {
    return budgets.getBudgetStart(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }

  function getBudgetEnd(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) override virtual public view returns (bytes32) {
    return budgets.getBudgetEnd(uint8(Budgets.Owner.ACTIVITY), _activitiesRef, _activityRef, _budgetRef);
  }
}
