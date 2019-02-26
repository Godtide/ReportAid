pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationExpenditure.sol";
import "./IATIBudgets.sol";

contract IATIOrganisationExpenditure is OrganisationExpenditure {

  IATIBudgets budgets;

  constructor(address _budgets) public {
    require(_budgets != address(0x0));
    budgets = IATIBudgets(_budgets);
  }

  function setExpenditure(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef, Budgets.Budget memory _expenditure) public {
    budgets.setBudget(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef, _expenditure);
  }

  function getNumExpenditures(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    return budgets.getNumBudgets(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef);
  }

  function getExpenditureReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    return budgets.getBudgetReference(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _index);
  }

  function getExpenditure(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (Budgets.Budget memory) {
    return budgets.getBudget(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }

  function getExpenditureLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (bytes32) {
    return budgets.getBudgetLine(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }

  function getExpenditureValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (uint256) {
    return budgets.getBudgetValue(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }

  function getExpenditureStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (uint8) {
    return budgets.getBudgetStatus(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }

  function getExpenditureStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (bytes32) {
    return budgets.getBudgetStart(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }

  function getExpenditureEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) public view returns (bytes32) {
    return budgets.getBudgetEnd(uint8(Budgets.BudgetType.ORGEXPENDITURE), _organisationsRef, _organisationRef, _expenditureRef);
  }
}
