pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportExpenditure.sol";
import "./Strings.sol";

contract IATIOrganisationReportExpenditure is OrganisationReportExpenditure {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private expenditureReferences;
  mapping(bytes32 => mapping(bytes32 => Expenditure)) private expenditures;

  event SetExpenditure(bytes32 _reportRef, bytes32 _expenditureRef, Expenditure _expenditure);

  function setExpenditure(Expenditure memory _expenditure) public {
    require (_expenditure.reportRef[0] != 0 &&
             _expenditure.expenditureRef[0] != 0 &&
             _expenditure.expenditureLine[0] != 0 &&
             _expenditure.finance.status > 0 &&
             _expenditure.finance.start[0] != 0 &&
             _expenditure.finance.end[0] != 0 );

    expenditures[_expenditure.reportRef][_expenditure.expenditureRef] = _expenditure;

    if (!getReportExists(_expenditure.reportRef)) {
      reportReferences.push(_expenditure.reportRef);
    }

    if (!getExpenditureExists(_expenditure.reportRef, _expenditure.expenditureRef)) {
    expenditureReferences[_expenditure.reportRef].push(_expenditure.expenditureRef);
    }

    emit SetExpenditure(_expenditure.reportRef, _expenditure.expenditureRef, _expenditure);
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

  function getExpenditureExists(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bool) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    bool exists = false;
    if ( !(expenditureReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_expenditureRef, expenditureReferences[_reportRef]);
      exists = (index != expenditureReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return expenditureReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _index < expenditureReferences[_reportRef].length);

    return expenditureReferences[_reportRef][_index];
  }

  function getExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (Expenditure memory) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef];
  }

  function getExpenditureLine(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].expenditureLine;
  }

  function getExpenditureValue(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.value;
  }

  function getExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint8) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.status;
  }

  function getExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.start;
  }

  function getExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_reportRef][_expenditureRef].finance.end;
  }
}
