pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationDocs {

  struct Report {
    bytes32 orgRef;
    bytes32 reportRef;
  }

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

  struct Document {
    Report report;
    bytes32 docRef;
    string title;
    string format;
    string url;
    bytes32 category;
    bytes32 countryRef;
    string desc;
    bytes32 lang;
    bytes32 date;
  }

  function setDocument(Document memory _document) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getReportDocExists(bytes32 _reportRef, bytes32 _docRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumReportDocs(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getReportDocReference(bytes32 _docRef, uint256 _index) public view returns (bytes32);

  function getDocument(bytes32 _reportRef, bytes32 _docRef) public view returns (Document memory);
  function getDocumentReportingOrg(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentTitle(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentFormat(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentURL(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentCategory(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCountry(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDescription(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentLang(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDate(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
}
