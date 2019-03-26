pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./ActivityTransactions.sol";
import "./Strings.sol";

contract IATIActivityTransactions is ActivityTransactions {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private transactionRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Transaction))) private transactions;

  event SetTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, Transaction _transaction);

  function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, Transaction memory _transaction) public {
    require (_activitiesRef[0] != 0 &&
             _activityRef[0] != 0 &&
             _transactionRef[0] != 0 &&
             _transaction.transactionType > uint8(TransactionType.NONE) &&
             _transaction.transactionType < uint8(TransactionType.MAX) &&
             _transaction.disbursementChannel > uint8(DisbursementChannel.NONE) &&
             _transaction.disbursementChannel < uint8(DisbursementChannel.MAX) &&
             _transaction.date[0] != 0 &&
             _transaction.value.date[0] != 0 &&
             _transaction.value.currency[0] != 0 &&
             _transaction.providerOrg.orgType >= ActivityTransactions.minOrgType &&
             _transaction.providerOrg.orgRef[0] != 0 &&
             _transaction.receiverOrg.orgType >= ActivityTransactions.minOrgType &&
             _transaction.receiverOrg.orgRef[0] != 0 &&
             _transaction.sectorDacCode >= ActivityTransactions.minDACCode &&
             _transaction.sectorDacCode <= ActivityTransactions.maxDACCode &&
             _transaction.territory[0] != 0 &&
             bytes(_transaction.description).length > 0);

    transactions[_activitiesRef][_activityRef][_transactionRef] = _transaction;

    if(!Strings.getExists(_transactionRef, transactionRefs[_activitiesRef][_activityRef])) {
      transactionRefs[_activitiesRef][_activityRef].push(_transactionRef);
    }

    emit SetTransaction(_activitiesRef, _activityRef, _transactionRef, _transaction);
  }

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0);

    return transactionRefs[_activitiesRef][_activityRef].length;
  }

  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _index < transactionRefs[_activitiesRef][_activityRef].length);

    return transactionRefs[_activitiesRef][_activityRef][_index];
  }

  function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (Transaction memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef];
  }


  function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].transactionType;
  }

  function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].disbursementChannel;
  }

  function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].flowType;
  }

  function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].tiedStatus;
  }

  function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].financeType;
  }

  function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].aidType;
  }

  function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].date;
  }

  function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].value.value;
  }

  function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].value.date;
  }

  function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].value.currency;
  }

  function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].providerOrg.orgRef;
  }

  function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].providerOrg.orgType;
  }

  function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].providerOrg.activityRef;
  }

  function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].receiverOrg.orgRef;
  }

  function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint8) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].receiverOrg.orgType;
  }

  function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].receiverOrg.activityRef;
  }

  function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (uint256) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].sectorDacCode;
  }

  function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (bytes32) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].territory;
  }

  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) public view returns (string memory) {
    require (_activitiesRef[0] != 0 && _activityRef[0] != 0 && _transactionRef[0] != 0);

    return transactions[_activitiesRef][_activityRef][_transactionRef].description;
  }
}
