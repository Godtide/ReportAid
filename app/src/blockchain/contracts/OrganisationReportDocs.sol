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

  struct Document {
    bytes32 reportRef;
    bytes32 docRef;
    string title;
    string format;
    string url;
    bytes32 category;
    bytes32 countryCode;
    string desc;
    bytes32 lang;
    bytes32 date;
  }

  function setDocument(Document memory _document) public;

  function getReportDocExists(bytes32 _reportRef, bytes32 _docRef) public view returns (bool);

  function getNumReportDocs(bytes32 _reportRef) public view returns (uint256);

  function getReportDocReference(bytes32 _docRef, uint256 _index) public view returns (bytes32);

  function getDocument(bytes32 _reportRef, bytes32 _docRef) public view returns (Document memory);

  function getDocumentTitle(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentFormat(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentURL(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentCategory(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCountry(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDescription(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentLang(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDate(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32);
}
