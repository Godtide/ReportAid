pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationBudgets {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct Budget {
    bytes32 budgetLine;
    Finance finance;
  }

  function setBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, Budget memory _budget) public;

  function getNumBudgets(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256);
  function getBudgetReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32);

  function getBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (Budget memory);
  function getBudgetLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getBudgetValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256);
  function getBudgetStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint8);
  function getBudgetStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getBudgetEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
}
