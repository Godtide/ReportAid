pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract Budgets {

  enum BudgetType {
    NONE,
    ORGEXPENDITURE,
    ORG,
    ORGRECIPIENT,
    ORGREGION,
    ORGCOUNTRY,
    MAX
  }

  enum Status {
    NONE,
    INDICATIVE,
    COMMITTED,
    MAX
  }

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct Budget {
    bytes32 budgetLine;
    bytes32 otherRef;
    Finance finance;
  }

  function setBudget(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, Budget memory _budget) public;

  function getNumBudgets(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef) public view returns (uint256);
  function getBudgetReference(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) public view returns (bytes32);

  function getBudget(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (Budget memory);
  function getBudgetLine(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32);
  function getOtherRef(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32);
  function getBudgetValue(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (uint256);
  function getBudgetStatus(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (uint8);
  function getBudgetStart(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32);
  function getBudgetEnd(uint8 _budgetType, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) public view returns (bytes32);
}
