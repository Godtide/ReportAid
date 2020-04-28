pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Budgets {

  enum Owner {
    NONE,
    ORGEXPENDITURE,
    ORG,
    ORGRECIPIENT,
    ORGREGION,
    ORGCOUNTRY,
    ACTIVITY,
    MAX
  }

  enum Status {
    NONE,
    INDICATIVE,
    COMMITTED,
    MAX
  }

  enum Type {
    NONE,
    ORIGINAL,
    REVISED,
    MAX
  }

  struct Finance {
    uint8 status;
    uint256 value;
    bytes32 start;
    bytes32 end;
  }

  struct Budget {
    uint8 budgetType;
    bytes32 budgetLine;
    bytes32 otherRef;
    Finance finance;
  }

  event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, Budget _budget);

  function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, Budget memory _budget) virtual public;

  function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) virtual public view returns (uint256);
  function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) virtual public view returns (bytes32);

  function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (Budget memory);
  function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
