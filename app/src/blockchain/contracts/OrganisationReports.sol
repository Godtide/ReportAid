pragma solidity ^0.4.24;

contract OrganisationReports {

  function setReport(string _reference, string _orgRef, string _reportingOrgRef, string _version, string _generatedTime) public;
  function setDefaults(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public;
  function setReportingOrgType(string _reference, string _orgRef, string _reportingOrgRef, uint8 _type, bool _isSecondary) public;
  function setAssociatedDocument(string _reference, string _docRef, bytes32[] _attributes) public;

  function getReportExists(string _reference) public view returns (bool);
  function getReportOrgExists(string _reference, string _orgRef) public view returns (bool);
  function getReportDocExists(string _reference, string _docRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumReportOrgs(string _reference) public view returns (uint256);
  function getNumReportDocs(string _reference) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (string);
  function getReportOrgReference(string _reference, uint256 _index) public view returns (string);
  function getReportDocReference(string _reference, uint256 _index) public view returns (string);

  function getReportingOrg(string _reference, string _orgRef) public view returns (string);
  function getLang(string _reference, string _orgRef) public view returns (string);
  function getCurrency(string _reference,  string _orgRef) public view returns (string);
  function getVersion(string _reference,  string _orgRef) public view returns (string);
  function getGeneratedTime(string _reference,  string _orgRef) public view returns (string);
  function getLastUpdatedTime(string _reference,  string _orgRef) public view returns (string);

  function getReportingOrgType(string _reference, string _orgRef) public view returns (uint8);
  function getReportingOrgIsSecondary(string _reference, string _orgRef) public view returns (bool);

  function getDocumentTitle(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentFormat(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentURL(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentCategory(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentCountry(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentDescription(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentLang(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentDate(string _reference, string _docRef) public view returns (bytes32);
}
