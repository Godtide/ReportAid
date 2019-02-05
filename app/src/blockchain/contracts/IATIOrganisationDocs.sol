pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationDocs.sol";
import "./Strings.sol";

contract IATIOrganisationDocs is OrganisationDocs {

  mapping(bytes32 => mapping(bytes32 => bytes32[])) private docRefs;
  mapping(bytes32 => mapping(bytes32 => mapping(bytes32 => Document))) private docs;

  event SetDocument(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef, Document _document);

  function setDocument(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef, Document memory _document) public {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0 &&
             bytes(_document.title).length > 0 &&
             bytes(_document.format).length > 0 &&
             bytes(_document.url).length > 0 &&
             _document.category[0] != 0 &&
             _document.countryRef[0] != 0 &&
             bytes(_document.desc).length > 0 &&
             _document.lang[0] != 0 &&
             _document.date[0] != 0);

    docs[_organisationsRef][_orgRef][_docRef] = _document;

    if(!Strings.getExists(_docRef, docRefs[_organisationsRef][_orgRef])) {
      docRefs[_organisationsRef][_orgRef].push(_docRef);
    }

    emit SetDocument(_organisationsRef, _orgRef, _docRef, _document);
  }

  function getNumDocs(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return docRefs[_organisationsRef][_orgRef].length;
  }

  function getDocReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _index < docRefs[_organisationsRef][_orgRef].length);

    return docRefs[_organisationsRef][_orgRef][_index];
  }

  function getDocument(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (Document memory) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef];
  }

  function getDocumentTitle(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].title;
  }

  function getDocumentFormat(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].format;
  }

  function getDocumentURL(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (string memory){
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].url;
  }

  function getDocumentCategory(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].category;
  }

  function getDocumentCountry(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].countryRef;
  }

  function getDocumentDescription(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (string memory) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].desc;
  }

  function getDocumentLang(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].lang;
  }

  function getDocumentDate(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _docRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 &&
             _orgRef[0] != 0 &&
             _docRef[0] != 0);

    return docs[_organisationsRef][_orgRef][_docRef].date;
  }
}
