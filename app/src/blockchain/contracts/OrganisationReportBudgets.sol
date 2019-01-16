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

  function setTotalBudget(Budget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getTotalBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumTotalBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getTotalBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getTotalBudget(bytes32 _reportRef, bytes32 _budgetRef) public view returns (Budget memory);
  function getTotalBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getTotalBudgetValue(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256);
  function getTotalBudgetStatus(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getTotalBudgetStart(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getTotalBudgetEnd(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
}
