pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./OrganisationDocs.sol";
import "./Strings.sol";

contract IATIOrganisationDocs is OrganisationDocs {

  bytes32[] reportReferences;
  mapping(bytes32 => bytes32[]) private reportDocReferences;
  mapping(bytes32 => mapping(bytes32 => Document)) private docs;

  event SetDocument(Document _document);

  function setDocument(Document memory _document) public {
    require (_document.report.reportRef[0] != 0 &&
             _document.report.orgRef[0] != 0 &&
             _document.docRef[0] != 0 &&
             bytes(_document.title).length > 0 &&
             bytes(_document.format).length > 0 &&
             bytes(_document.url).length > 0 &&
             _document.category[0] != 0 &&
             _document.countryRef[0] != 0 &&
             bytes(_document.desc).length > 0 &&
             _document.lang[0] != 0 &&
             _document.date[0] != 0);

    docs[_document.report.reportRef][_document.docRef] = _document;

    if (!getReportExists(_document.report.reportRef)) {
      reportReferences.push(_document.report.reportRef);
    }

    if(!getReportDocExists(_document.report.reportRef, _document.docRef)) {
      reportDocReferences[_document.report.reportRef].push(_document.docRef);
    }

    emit SetDocument(_document);
  }

  function getReportExists(bytes32 _reportRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_reportRef, reportReferences);
      exists = (index != reportReferences.length);
    }
    return exists;
  }

  function getReportDocExists(bytes32 _reportRef, bytes32 _docRef) public view returns (bool) {
    require (_reportRef[0] != 0);

    bool exists = false;
    if ( !(reportDocReferences[_reportRef].length == 0) ) {
      uint256 index = Strings.getIndex(_docRef, reportDocReferences[_reportRef]);
      exists = (index != reportDocReferences[_reportRef].length);
    }
    return exists;
  }

  function getNumReports() public view returns (uint256) {
    return reportReferences.length;
  }

  function getNumReportDocs(bytes32 _reportRef) public view returns (uint256) {
    require (_reportRef[0] != 0);

    return reportDocReferences[_reportRef].length;
  }

  function getReportReference(uint256 _index) public view returns (bytes32) {
    require (_index < reportReferences.length);

    return reportReferences[_index];
  }

  function getReportDocReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {
    require (_reportRef[0] != 0);

    return reportDocReferences[_reportRef][_index];
  }

  function getDocument(bytes32 _reportRef, bytes32 _docRef) public view returns (Document memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef];
  }

  function getDocumentReportingOrg(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].report.orgRef;
  }

  function getDocumentTitle(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].title;
  }

  function getDocumentFormat(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].format;
  }

  function getDocumentURL(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory){
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].url;
  }

  function getDocumentCategory(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].category;
  }

  function getDocumentCountry(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].countryRef;
  }

  function getDocumentDescription(bytes32 _reportRef, bytes32 _docRef) public view returns (string memory) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].desc;
  }

  function getDocumentLang(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
    require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].lang;
  }

  function getDocumentDate(bytes32 _reportRef, bytes32 _docRef) public view returns (bytes32) {
  require (_reportRef[0] != 0 && _docRef[0] != 0);

    return docs[_reportRef][_docRef].date;
  }
}
