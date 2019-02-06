pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationExpenditure {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct Expenditure {
    bytes32 expenditureRef;
    bytes32 expenditureLine;
    Finance finance;
  }

  function setExpenditure(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef, Expenditure memory _expenditure) public;

  function getNumExpenditures(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256);
  function getExpenditureReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32);

  function getExpenditure(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (Expenditure memory);
  function getExpenditureLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32);
  function getExpenditureValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256);
  function getExpenditureStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint8);
  function getExpenditureStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32);
  function getExpenditureEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32);
}
