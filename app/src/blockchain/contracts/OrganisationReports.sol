pragma solidity ^0.4.24;

contract OrganisationReports {

  function setReport(string _reference, string _version, string _generatedTime) public;
  function setOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public;
  function setReportingOrganisation(string _reference, string _reportingOrgRef, string _type, bool _isSecondary) public;
  function setDocument(string _reference, string _docRef, bytes32[] _attributes) public;

  function getReportExists(string _reference) public view returns (bool);
  function getReportDocExists(string _reference, string _docRef) public view returns (bool);

  function getVersion(string _reference) public view returns (string);

  function getNumReports() public view returns (uint256);
  function getNumReportDocs(string _reference) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (string);
  function getReportDocReference(string _reference, uint256 _index) public view returns (string);

  function getOrganisation(string _reference) public view returns (string);
  function getOrganisationDefaultLang(string _reference) public view returns (string);
  function getOrganisationDefaultCurrency(string _reference) public view returns (string);

  function getReportingOrganisation(string _reference) public view returns (string);
  function getReportingOrganisationType(string _reference) public view returns (string);
  function getReportingOrganisationIsSecondary(string _reference) public view returns (bool);

  function getDocumentTitle(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentFormat(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentURL(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentCategory(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentCountry(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentDescription(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentLang(string _reference, string _docRef) public view returns (bytes32);
  function getDocumentDate(string _reference, string _docRef) public view returns (bytes32);
}
