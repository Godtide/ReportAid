pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationBudgets.sol";
import "./IATIBudgets.sol";

contract IATIOrganisationBudgets is OrganisationBudgets {

  IATIBudgets budgets;

  constructor(address _budgets) override virtual public {
    require(_budgets != address(0x0));
    budgets = IATIBudgets(_budgets);
  }

  function setBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) override virtual public {
    budgets.setBudget(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef, _budget);
  }

  function getNumBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) override virtual public view returns (uint256) {
    return budgets.getNumBudgets(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef);
  }

  function getBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) override virtual public view returns (bytes32) {
    return budgets.getBudgetReference(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _index);
  }

  function getBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (Budgets.Budget memory) {
    return budgets.getBudget(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }

  function getBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (bytes32) {
    return budgets.getBudgetLine(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }

  function getBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (uint256) {
    return budgets.getBudgetValue(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }

  function getBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (uint8) {
    return budgets.getBudgetStatus(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }

  function getBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (bytes32) {
    return budgets.getBudgetStart(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }

  function getBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) override virtual public view returns (bytes32) {
    return budgets.getBudgetEnd(uint8(Budgets.Owner.ORG), _organisationsRef, _organisationRef, _budgetRef);
  }
}
