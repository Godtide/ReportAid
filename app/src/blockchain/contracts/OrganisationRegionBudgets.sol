pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationRegionBudgets {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct RegionBudget {
    uint256 regionRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setRegionBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, RegionBudget memory _budget) public;

  function getNumRegionBudgets(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256);
  function getRegionBudgetReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32);

  function getRegionsBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (RegionBudget memory);
  function getRegionsBudgetRegion(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256);
  function getRegionsBudgetLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRegionsBudgetValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256);
  function getRegionsBudgetStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint8);
  function getRegionsBudgetStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getRegionsBudgetEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
}
