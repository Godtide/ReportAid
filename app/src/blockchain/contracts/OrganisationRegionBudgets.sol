pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract OrganisationRegionBudgets {

  function setRegionBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) virtual public;

  function getNumRegionBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getRegionBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getRegionsBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (Budgets.Budget memory);
  function getRegionsBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRegionsBudgetRegion(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRegionsBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getRegionsBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getRegionsBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRegionsBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
