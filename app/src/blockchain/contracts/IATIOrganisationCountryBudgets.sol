pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationCountryBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationCountryBudgets is OrganisationCountryBudgets {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private budgetRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => CountryBudget))) private budgets;

  event SetCountryBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, CountryBudget _budget);

  function setCountryBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, CountryBudget memory _budget) public {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _budgetRef[0] != 0 &&
             _budget.countryRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_organisationsRef][_organisationRef][_budgetRef] = _budget;

    if(!Strings.getExists(_budgetRef, budgetRefs[_organisationsRef][_organisationRef])) {
      budgetRefs[_organisationsRef][_organisationRef].push(_budgetRef);
    }

    emit SetCountryBudget(_organisationsRef, _organisationRef, _budgetRef, _budget);
  }

  function getNumCountryBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return budgetRefs[_organisationsRef][_organisationRef].length;
  }

  function getCountryBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _index < budgetRefs[_organisationsRef][_organisationRef].length);

    return budgetRefs[_organisationsRef][_organisationRef][_index];
  }

  function getCountryBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (CountryBudget memory) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef];
  }

  function getCountryBudgetCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].countryRef;
  }

  function getCountryBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].budgetLine;
  }

  function getCountryBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.value;
  }

  function getCountryBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.status;
  }

  function getCountryBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.start;
  }

  function getCountryBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0 && _budgetRef[0] != 0);

    return budgets[_organisationsRef][_organisationRef][_budgetRef].finance.end;
  }
}
