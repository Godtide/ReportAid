pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ActivityTransactions {

  uint8 constant minFlowType = 10;
  uint8 constant maxFlowType = 50;

  uint8 constant minFinanceType = 1;
  uint256 constant maxFinanceType = 1100;

  uint256 constant minDACCode = 11110;
  uint256 constant maxDACCode = 99820;

  uint8 constant minOrgType = 10;

  enum TransactionType {
    NONE,
    INCOMINGFUNDS,
    OUTGOINGCOMMITMENT,
    DISBURSEMENT,
    EXPENDITURE,
    INTERESTPAYMENT,
    LOANREPAYMENT,
    REIMBURSEMENT,
    PURCHASEOFEQUITY,
    SALEOFEQUITY,
    CREDITGUARANTEE,
    INCOMINGCOMMITMENT,
    OUTGOINGPLEDGE,
    INCOMINGPLEDGE,
    MAX
  }

  enum DisbursementChannel {
    NONE,
    TREASURY,
    BANKACCOUNT,
    NGO,
    DONORMANAGED,
    MAX
  }

  enum TiedStatus {
    UNSPECIFIED,
    UNSPECIFIEDONE,
    NONE,
    PARTIALLYTIED,
    TIED,
    UNTIED,
    MAX
  }

  struct Value {
    uint256 value;
    bytes32 date;
    bytes32 currency;
  }

  struct Org {
    uint8 orgType;
    bytes32 orgRef;
    bytes32 activityRef;
  }

  struct Transaction {
    uint8 transactionType;
    uint8 disbursementChannel;
    uint8 flowType;
    uint8 tiedStatus;
    uint256 financeType;
    bytes32 aidType;
    bytes32 date;
    Value value;
    Org providerOrg;
    Org receiverOrg;
    uint256 sectorDacCode;
    bytes32 territory;
    string description;
  }

  function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, Transaction memory _transaction) virtual public;

  function getNum(bytes32 _activitiesRef, bytes32 _activityRef) virtual public view returns (uint256);
  function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) virtual public view returns (bytes32);

  function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (Transaction memory);

  function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint256);
  function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint256);
  function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint8);
  function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (uint256);
  function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (bytes32);
  function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) virtual public view returns (string memory);
}
