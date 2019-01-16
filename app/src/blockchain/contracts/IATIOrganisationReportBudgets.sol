pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportBudgets is OrganisationReportBudgets {

  bytes32[] reportReferences;

  mapping(bytes32 => bytes32[]) private budgetReferences;
  mapping(bytes32 => mapping(bytes32 => Budget)) private budgets;

  mapping(bytes32 => bytes32[]) private expenditureReferences;
  mapping(bytes32 => mapping(bytes32 => Expenditure)) private expenditure;

  mapping(bytes32 => bytes32[]) private recipientOrgReferences;
  mapping(bytes32 => mapping(bytes32 => RecipientBudget)) private recipientOrgBudgets;

  mapping(bytes32 => bytes32[]) private regionReferences;
  mapping(bytes32 => mapping(bytes32 => RegionBudget)) private regionBudgets;

  mapping(bytes32 => bytes32[]) private countryReferences;
  mapping(bytes32 => mapping(bytes32 => CountryBudget)) private countryBudgets;

  event SetTotalBudget(bytes32 _reportRef, bytes32 _budgetRef, Budget _budget);
  event SetExpenditure(bytes32 _reportRef, bytes32 _expenditureRef, Expenditure _expenditure);
  event SetRecipientOrgBudget(bytes32 _reportRef, bytes32 _budgetRef, RecipientBudget _budget);
  event SetRegionBudget(bytes32 _reportRef, bytes32 _budgetRef, RegionBudget _budget);
  event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, CountryBudget _budget);

  function setTotalBudget(Budget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    budgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getTotalBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      budgetReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetTotalBudget(_budget.reportRef, _budget.budgetRef, _budget);
  }

  function setExpenditure(Expenditure memory _expenditure) public {
    require (_expenditure.reportRef[0] != 0 &&
             _expenditure.expenditureRef[0] != 0 &&
             _expenditure.expenseLine[0] != 0 &&
             _expenditure.finance.status[0] != 0 &&
             _expenditure.finance.start[0] != 0 &&
             _expenditure.finance.end[0] != 0 );

    expenditure[_expenditure.reportRef][_expenditure.expenditureRef] = _expenditure;

    if (!getReportExists(_expenditure.reportRef)) {
      reportReferences.push(_expenditure.reportRef);
    }

    if (!getExpenditureExists(_expenditure.reportRef, _expenditure.expenditureRef)) {
    expenditureReferences[_expenditure.reportRef].push(_expenditure.expenditureRef);
    }

    emit SetExpenditure(_expenditure.reportRef, _expenditure.expenditureRef, _expenditure);
  }

  function setRecipientOrgBudget(RecipientBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.orgRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    recipientOrgBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getRecipientOrgBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      recipientOrgReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetRecipientOrgBudget(_budget.reportRef, _budget.budgetRef, _budget);
  }

  function setRegionBudget(RegionBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.regionRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    regionBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getRegionBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      regionReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetRegionBudget(_budget.reportRef, _budget.budgetRef, _budget);
  }

  function setCountryBudget(CountryBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.countryRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status[0] != 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    countryBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getRegionBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      countryReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetCountryBudget(_budget.reportRef, _budget.budgetRef, _budget);
  }

  function getReportExists(bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences);
      exists = (index != reportReferences.length);
    }
    return exists;
  }

  function getTotalBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(budgetReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, budgetReferences[_reportRef]);
      exists = (index != budgetReferences[_reportRef].length);
    }
    return exists;
  }

  function getExpenditureExists(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    bool exists = false;
    if ( !(expenditureReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_expenditureRef, expenditureReferences[_reportRef]);
      exists = (index != expenditureReferences[_reportRef].length);
    }
    return exists;
  }

  function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(recipientOrgReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, recipientOrgReferences[_reportRef]);
      exists = (index != recipientOrgReferences[_reportRef].length);
    }
    return exists;
  }

  function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(regionReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, regionReferences[_reportRef]);
      exists = (index != regionReferences[_reportRef].length);
    }
    return exists;
  }

  function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(countryReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, countryReferences[_reportRef]);
      exists = (index != countryReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumTotalBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return budgetReferences[_reportRef].length;
  }

  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return expenditureReferences[_reportRef].length;
  }

  function getNumRecipientOrgBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return recipientOrgReferences[_reportRef].length;
  }

  function getNumRegionBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return regionReferences[_reportRef].length;
  }

  function getNumCountryBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return countryReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getTotalBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < budgetReferences[_reportRef].length);

    return budgetReferences[_reportRef][_index];
  }

  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < expenditureReferences[_reportRef].length);

    return expenditureReferences[_reportRef][_index];
  }

  function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < recipientOrgReferences[_reportRef].length);

    return recipientOrgReferences[_reportRef][_index];
  }

  function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < regionReferences[_reportRef].length);

    return regionReferences[_reportRef][_index];
  }

  function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < countryReferences[_reportRef].length);

    return countryReferences[_reportRef][_index];
  }

  function getTotalBudget(bytes32 _reportRef, bytes32 _budgetRef) public view returns (Budget memory) {}
  function getTotalBudgetValue(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStatus(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStart(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetEnd(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}

  function getExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (Expenditure memory) {}
  function getExpenditureValue(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}

  function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (RecipientBudget memory) {}
  function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}

  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) public view returns (RegionBudget memory) {}
  function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}

  function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) public view returns (CountryBudget memory) {}
  function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
}
