pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract ActivityBudgets {

  function setBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef, Budgets.Budget memory _budget) virtual public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (Budgets.Budget memory);

  function getBudgetType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getBudgetValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getBudgetStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getBudgetStart(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getBudgetEnd(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
