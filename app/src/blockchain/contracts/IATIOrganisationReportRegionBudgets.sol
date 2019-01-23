pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportRegionBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportRegionBudgets is OrganisationReportRegionBudgets {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private regionReferences;
  mapping(bytes32 => mapping(bytes32 => RegionBudget)) private regionBudgets;

  event SetRegionBudget(bytes32 _reportRef, bytes32 _budgetRef, RegionBudget _budget);

  function setRegionBudget(RegionBudget memory _budget) public {
    require (_budget.reportRef[0] != 0 &&
             _budget.budgetRef[0] != 0 &&
             _budget.regionRef[0] != 0 &&
             _budget.budgetLine[0] != 0 &&
             _budget.finance.status > 0 &&
             _budget.finance.start[0] != 0 &&
             _budget.finance.end[0] != 0 );

    regionBudgets[_budget.reportRef][_budget.budgetRef] = _budget;

    if (!getReportExists(_budget.reportRef)) {
      reportReferences.push(_budget.reportRef);
    }

    if (!getRegionBudgetExists(_budget.reportRef, _budget.budgetRef)) {
      regionReferences[_budget.reportRef].push(_budget.budgetRef);
    }

    emit SetRegionBudget(_budget.reportRef, _budget.budgetRef, _budget);
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

  function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _budgetRef[0] != 0);

    bool exists = false;
    if ( !(regionReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_budgetRef, regionReferences[_reportRef]);
      exists = (index != regionReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumRegionBudgets(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return regionReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < regionReferences[_reportRef].length);

    return regionReferences[_reportRef][_index];
  }

  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) public view returns (RegionBudget memory) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef];
  }
  function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _regionRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef].budgetLine;
  }

  function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef].finance.value;
  }

  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef].finance.status;
  }

  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef].finance.start;
  }

  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _regionRef[0] != 0);

    return regionBudgets[_reportRef][_regionRef].finance.end;
  }
}
