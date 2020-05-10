pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

// Orgs

struct Org {
    string name;
    string identifier;
}

// Activities

struct ActivitiesData {
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
}

struct ReportingOrgData {
    uint8 orgType;
    bool isSecondary;
    bytes32 orgRef;
}

struct ActivityData {
    bool humanitarian;
    uint8 hierarchy;
    uint8 status;
    uint8 budgetNotProvided;
    ReportingOrgData reportingOrg;
    bytes32 lastUpdated;
    bytes32 lang;
    bytes32 currency;
    bytes32 linkedData;
    bytes32 identifier;
    string title;
    string description;
}

struct Additional {
    uint8 scope;
    uint8 capitalSpend;
    uint8 collaborationType;
    uint8 defaultFlowType;
    uint8 defaultTiedStatus;
    bytes32 defaultAidType;
    uint256 defaultFinanceType;
}

struct Date {
    uint8 dateType;
    bytes32 date;
    string narrative;
}

struct ParticipatingOrg {
    uint8 orgType;
    uint8 role;
    uint256 crsChannelCode;
    bytes32 lang;
    bytes32 orgRef;
    string narrative;
}

struct RelatedActivity {
    uint8 relationType;
    bytes32 activityRef;
}

struct Sector {
    uint8 percentage;
    uint256 dacCode;
}

struct Territory {
    uint8 percentage;
    bytes32 territory;
}

struct Value {
    uint256 value;
    bytes32 date;
    bytes32 currency;
}

struct TransactionOrg {
    uint8 orgType;
    bytes32 orgRef;
    bytes32 activityIdentifier;
}

struct Transaction {
    uint8 transactionType;
    uint8 disbursementChannel;
    uint8 flowType;
    uint8 tiedStatus;
    TransactionOrg providerOrg;
    TransactionOrg receiverOrg;
    Value value;
    uint256 financeType;
    bytes32 aidType;
    bytes32 date;
    uint256 sectorDacCode;
    bytes32 territory;
    string description;
}

struct Finance {
    uint8 status;
    uint256 value;
    bytes32 start;
    bytes32 end;
}

struct Budget {
    uint8 budgetType;
    bytes32 budgetLine;
    bytes32 otherRef;
    Finance finance;
}

// Organisations

struct Organisations {
    bytes32 version;
    bytes32 generatedTime;
}

struct Organisation {
    ReportingOrgData reportingOrg;
    bytes32 orgRef;
    bytes32 lang;
    bytes32 currency;
    bytes32 lastUpdatedTime;
}

struct Document {
    bytes32 category;
    bytes32 countryRef;
    bytes32 lang;
    bytes32 date;
    string desc;
    string title;
    string format;
    string url;
}
