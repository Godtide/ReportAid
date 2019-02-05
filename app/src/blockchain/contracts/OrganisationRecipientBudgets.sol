pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationRecipientBudgets {

  struct Report {
    bytes32 orgRef;
    bytes32 reportRef;
  }

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct RecipientBudget {
    Report report;
    bytes32 budgetRef;
    bytes32 recipientOrgRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setRecipientBudget(RecipientBudget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getRecipientBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumRecipientBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getRecipientBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getRecipientBudget(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (RecipientBudget memory);
  function getRecipientBudgetReportingOrg(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32);
  function getRecipientBudgetOrg(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32);
  function getRecipientBudgetLine(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32);
  function getRecipientBudgetValue(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (uint256);
  function getRecipientBudgetStatus(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (uint8);
  function getRecipientBudgetStart(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32);
  function getRecipientBudgetEnd(bytes32 _reportRef, bytes32 _recipientBudgetRef) public view returns (bytes32);
}
