pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReportRecipientBudgets {

  struct Finance {
    uint256 value;
    bytes32 status;
    bytes32 start;
    bytes32 end;
  }

  struct RecipientBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 orgRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setRecipientOrgBudget(RecipientBudget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumRecipientOrgBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (RecipientBudget memory);
  function getRecipientOrgBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32);
  function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32);
  function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (bytes32);
}
