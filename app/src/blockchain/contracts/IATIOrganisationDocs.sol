pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationDocs.sol";
import "./Strings.sol";

contract IATIOrganisationDocs is OrganisationDocs {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private docRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Document))) private docs;

  function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, Document memory _document) public {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0 &&
             bytes(_document.title).length > 0 &&
             bytes(_document.format).length > 0 &&
             bytes(_document.url).length > 0 &&
             _document.category[0] != 0 &&
             _document.countryRef[0] != 0 &&
             bytes(_document.desc).length > 0 &&
             _document.lang[0] != 0 &&
             _document.date[0] != 0);

    docs[_organisationsRef][_organisationRef][_docRef] = _document;

    if(!Strings.getExists(_docRef, docRefs[_organisationsRef][_organisationRef])) {
      docRefs[_organisationsRef][_organisationRef].push(_docRef);
    }

    emit SetDocument(_organisationsRef, _organisationRef, _docRef, _document);
  }

  function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return docRefs[_organisationsRef][_organisationRef].length;
  }

  function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _index < docRefs[_organisationsRef][_organisationRef].length);

    return docRefs[_organisationsRef][_organisationRef][_index];
  }

  function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (Document memory) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef];
  }

  function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].title;
  }

  function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].format;
  }

  function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory){
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].url;
  }

  function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].category;
  }

  function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].countryRef;
  }

  function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].desc;
  }

  function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].lang;
  }

  function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_organisationRef][_docRef].date;
  }
}
