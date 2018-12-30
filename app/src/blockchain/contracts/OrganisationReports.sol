pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReports {

  enum DocAttributes {
    TITLE,
    FORMAT,
    URL,
    CATEGORY,
    COUNTRYCODE,
    DESC,
    LANG,
    DATE
  }

  struct ReportingOrganisation {
    bytes32 orgRef;
    uint8 orgType;
    bool isSecondary;
  }

  struct Report {
    bytes32 reportRef;
    ReportingOrganisation reportingOrg;
    bytes32 issuingOrgRef;
    bytes32 version;
    bytes32 lang;
    bytes32 currency;
    bytes32 generatedTime;
    bytes32 lastUpdatedTime;
  }

  struct Document {
    bytes32 reportRef;
    bytes32 docRef;
    string title;
    bytes32 format;
    string url;
    bytes32 category;
    bytes32 countryCode;
    string desc;
    bytes32 lang;
    bytes32 date;
  }

  function setReport(Report memory _report) public;
  function setDocument(Document memory _document) public;

  function getReportExists(bytes32 _reference) public view returns (bool);
  function getIssuingOrgExists(bytes32 _reference, bytes32 _orgRef) public view returns (bool);
  function getReportDocExists(bytes32 _reference, bytes32 _docRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumReportOrgs(bytes32 _reference) public view returns (uint256);
  function getNumReportDocs(bytes32 _reference) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getReportOrgReference(bytes32 _reference, uint256 _index) public view returns (bytes32);
  function getReportDocReference(bytes32 _reference, uint256 _index) public view returns (bytes32);

  function getReport(bytes32 _reference, bytes32 _orgRef) public view returns (Report memory);

  function getReportingOrg(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32);
  function getLang(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32);
  function getCurrency(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getVersion(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);

  function getReportingOrgType(bytes32 _reference, bytes32 _reportingOrgRef) public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _reference, bytes32 _reportingOrgRef) public view returns (bool);

  function getDocument(bytes32 _reference, bytes32 _docRef) public view returns (Document memory);
  function getDocumentTitle(bytes32 _reference, bytes32 _docRef) public view returns (string memory);
  function getDocumentFormat(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentURL(bytes32 _reference, bytes32 _docRef) public view returns (string memory);
  function getDocumentCategory(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCountry(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDescription(bytes32 _reference, bytes32 _docRef) public view returns (string memory);
  function getDocumentLang(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDate(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
}
