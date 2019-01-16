pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReportBudgets {

  struct Finance {
    uint256 value;
    bytes32 status;
    bytes32 start;
    bytes32 end;
  }

  struct Budget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 budgetLine;
    Finance finance;
  }

  struct Expenditure {
    bytes32 reportRef;
    bytes32 expenditureRef;
    bytes32 expenseLine;
    Finance finance;
  }

  struct RecipientBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 orgRef;
    bytes32 budgetLine;
    Finance finance;
  }

  struct RegionBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 regionRef;
    bytes32 budgetLine;
    Finance finance;
  }

  struct CountryBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 countryRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setTotalBudget(Budget memory _budget) public;
  function setExpenditure(Expenditure memory _expenditure) public;
  function setRecipientOrgBudget(RecipientBudget memory _budget) public;
  function setRegionBudget(RegionBudget memory _budget) public;
  function setCountryBudget(CountryBudget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getTotalBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);
  function getExpenditureExists(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bool);
  function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);
  function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);
  function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumTotalBudgets(bytes32 _reportRef) public view returns (uint256);
  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256);
  function getNumRecipientOrgBudgets(bytes32 _reportRef) public view returns (uint256);
  function getNumRegionBudgets(bytes32 _reportRef) public view returns (uint256);
  function getNumCountryBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getTotalBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);
  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);
  function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);
  function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);
  function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getTotalBudget(bytes32 _reportRef, bytes32 _budgetRef) public view returns (Budget memory);
  function getTotalBudgetValue(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256);
  function getTotalBudgetStatus(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256);
  function getTotalBudgetStart(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256);
  function getTotalBudgetEnd(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256);

  function getExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (Expenditure memory);
  function getExpenditureValue(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256);
  function getExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256);
  function getExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256);
  function getExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256);

  function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (RecipientBudget memory);
  function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256);

  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) public view returns (RegionBudget memory);
  function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256);
  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256);
  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256);
  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256);

  function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) public view returns (CountryBudget memory);
  function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256);
  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256);
  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256);
  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256);
}
