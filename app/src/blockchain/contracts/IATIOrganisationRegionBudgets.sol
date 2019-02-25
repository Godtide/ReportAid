pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationRegionBudgets.sol";
import "./IATIBudgets.sol";

contract IATIOrganisationRegionBudgets is OrganisationRegionBudgets {

  IATIBudgets budgets;

  constructor(address _budgets) public {
    require(_budgets != address(0x0));
    budgets = IATIBudgets(_budgets);
  }

  function setRegionBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) public {
    budgets.setBudget(_organisationsRef, _organisationRef, _budgetRef, _budget);
  }

  function getNumRegionBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    return budgets.getNumBudgets(_organisationsRef, _organisationRef);
  }

  function getRegionBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    return budgets.getBudgetReference(_organisationsRef, _organisationRef, _index);
  }

  function getRegionsBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (Budgets.Budget memory) {
    return budgets.getBudget(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetLine(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetRegion(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getOtherRef(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256) {
    return budgets.getBudgetValue(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint8) {
    return budgets.getBudgetStatus(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetStart(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRegionsBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetEnd(_organisationsRef, _organisationRef, _budgetRef);
  }
}
