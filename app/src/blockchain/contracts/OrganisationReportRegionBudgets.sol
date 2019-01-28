pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReportRegionBudgets {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct RegionBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    uint256 regionRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setRegionBudget(RegionBudget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumRegionBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (RegionBudget memory);
  function getRegionsBudgetRegion(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint256);
  function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32);
  function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint256);
  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint8);
  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32);
  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32);
}
