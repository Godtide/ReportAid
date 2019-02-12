pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationDocs {

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
    string title;
    string format;
    string url;
    bytes32 category;
    bytes32 countryRef;
    string desc;
    bytes32 lang;
    bytes32 date;
  }

  function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, Document memory _document) public;

  function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256);
  function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32);

  function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (Document memory);

  function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory);
  function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32);
}
