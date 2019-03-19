pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

contract ActivityBudgets {

  function setBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef, Budgets.Budget memory _budget) public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32);

  function getBudget(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) public view returns (Budgets.Budget memory);
  function getBudgetValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) public view returns (uint256);
  function getBudgetStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) public view returns (uint8);
  function getBudgetStart(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) public view returns (bytes32);
  function getBudgetEnd(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _budgetRef) public view returns (bytes32);
}
