pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationRecipientBudgets {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct RecipientBudget {
    bytes32 recipientOrgRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, RecipientBudget memory _budget) public;

  function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256);
  function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32);

  function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (RecipientBudget memory);
  function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256);
  function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint8);
  function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32);
}
