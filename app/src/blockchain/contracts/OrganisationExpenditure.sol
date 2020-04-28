pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract OrganisationExpenditure {

  function setExpenditure(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef, Budgets.Budget memory _expenditure) virtual public;

  function getNumExpenditures(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getExpenditureReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getExpenditure(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (Budgets.Budget memory);
  function getExpenditureLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (bytes32);
  function getExpenditureValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (uint256);
  function getExpenditureStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (uint8);
  function getExpenditureStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (bytes32);
  function getExpenditureEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _expenditureRef) virtual public view returns (bytes32);
}
