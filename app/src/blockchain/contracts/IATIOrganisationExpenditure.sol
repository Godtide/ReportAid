pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationExpenditure.sol";
import "./Strings.sol";

contract IATIOrganisationExpenditure is OrganisationExpenditure {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private expenditureRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Expenditure))) private expenditures;

  event SetExpenditure(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef, Expenditure _expenditure);

  function setExpenditure(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef, Expenditure memory _expenditure) public {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _expenditureRef[0] != 0 &&
             _expenditure.expenditureLine[0] != 0 &&
             _expenditure.finance.status > 0 &&
             _expenditure.finance.start[0] != 0 &&
             _expenditure.finance.end[0] != 0 );

     expenditures[_organisationsRef][_orgRef][_expenditureRef] = _expenditure;

     if(!Strings.getExists(_expenditureRef, expenditureRefs[_organisationsRef][_orgRef])) {
       expenditureRefs[_organisationsRef][_orgRef].push(_expenditureRef);
     }

     emit SetExpenditure(_organisationsRef, _orgRef, _expenditureRef, _expenditure);
  }

  function getNumExpenditures(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return expenditureRefs[_organisationsRef][_orgRef].length;
  }

  function getExpenditureReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _index < expenditureRefs[_organisationsRef][_orgRef].length);

    return expenditureRefs[_organisationsRef][_orgRef][_index];
  }

  function getExpenditure(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (Expenditure memory) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef];
  }

  function getExpenditureLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef].expenditureLine;
  }

  function getExpenditureValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef].finance.value;
  }

  function getExpenditureStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef].finance.status;
  }

  function getExpenditureStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef].finance.start;
  }

  function getExpenditureEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _expenditureRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0 && _expenditureRef[0] != 0);

    return expenditures[_organisationsRef][_orgRef][_expenditureRef].finance.end;
  }
}