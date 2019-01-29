pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportRegionBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportRegionBudgets is OrganisationReportRegionBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private regionBudgetReferences;
  mapping(bytes32 => mapping(bytes32 => RegionBudget)) private regionBudgets;

  event SetRegionBudget(RegionBudget _budget);

  function setRegionBudget(RegionBudget memory _budget) public {
    require (_budget.report.reportRef[0] != 0 &&
             _budget.report.orgRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.regionRef > 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    regionBudgets[_budget.report.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.report.reportRef)) {
      reportReferences.push(_budget.report.reportRef);
    }

    if (!getRegionBudgetExists(_budget.report.reportRef, _budget.budgetRef)) {
      regionBudgetReferences[_budget.report.reportRef].push(_budget.budgetRef);
    }

    emit SetRegionBudget(_budget);
  }

  function getReportExists(bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences);
      exists = (index != reportReferences.length);
    }
    return exists;
  }

  function getRegionBudgetExists(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    bool exists = false;
    if ( !(regionBudgetReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_regionBudgetRef, regionBudgetReferences[_reportRef]);
      exists = (index != regionBudgetReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumRegionBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return regionBudgetReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < regionBudgetReferences[_reportRef].length);

    return regionBudgetReferences[_reportRef][_index];
  }

  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (RegionBudget memory) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef];
  }

  function getRegionsBudgetReportingOrg(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].report.orgRef;
  }

  function getRegionsBudgetRegion(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].regionRef;
  }

  function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].budgetLine;
  }

  function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].finance.value;
  }

  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].finance.status;
  }

  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].finance.start;
  }

  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionBudgetRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionBudgetRef[0] != 0);

    return regionBudgets[_reportRef][_regionBudgetRef].finance.end;
  }
}
