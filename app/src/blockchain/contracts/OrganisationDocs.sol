pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract OrganisationDocs {

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

  event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, Document _document);

  function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, Document memory _document) virtual public;

  function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (Document memory);

  function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (string memory);
  function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (string memory);
  function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (string memory);
  function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (bytes32);
  function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (bytes32);
  function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (string memory);
  function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (bytes32);
  function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) virtual public view returns (bytes32);
}
