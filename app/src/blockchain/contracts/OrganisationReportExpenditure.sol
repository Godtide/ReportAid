pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReportExpenditure {

  struct Finance {
    uint256 value;
    bytes32 status;
    bytes32 start;
    bytes32 end;
  }

  struct Expenditure {
    bytes32 reportRef;
    bytes32 expenditureRef;
    bytes32 expenditureLine;
    Finance finance;
  }

  function setExpenditure(Expenditure memory _expenditure) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getExpenditureExists(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (Expenditure memory);
  function getExpenditureLine(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getExpenditureValue(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256);
  function getExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32);
  function getExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32);
  function getExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (bytes32);
}