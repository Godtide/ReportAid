pragma solidity ^0.5.0;

contract OrganisationReports {

  function setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime) public;
  function setDefaults(bytes32 _reference, bytes32 _orgRef, bytes32 _defaultLang, bytes32 _defaultCurrency) public;
  function setReportingOrgType(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, uint8 _type, bool _isSecondary) public;
  function setAssociatedDocument(bytes32 _reference, bytes32 _docRef, bytes32[] memory _attributes) public;
  function setDocumentDescription(bytes32 _reference, bytes32 _docRef, string memory _description) public;

  function getReportExists(bytes32 _reference) public view returns (bool);
  function getReportOrgExists(bytes32 _reference, bytes32 _orgRef) public view returns (bool);
  function getReportDocExists(bytes32 _reference, bytes32 _docRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumReportOrgs(bytes32 _reference) public view returns (uint256);
  function getNumReportDocs(bytes32 _reference) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getReportOrgReference(bytes32 _reference, uint256 _index) public view returns (bytes32);
  function getReportDocReference(bytes32 _reference, uint256 _index) public view returns (bytes32);

  function getReportingOrg(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32);
  function getLang(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32);
  function getCurrency(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getVersion(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32);

  function getReportingOrgType(bytes32 _reference, bytes32 _orgRef) public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _reference, bytes32 _orgRef) public view returns (bool);

  function getDocumentTitle(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentFormat(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentURL(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCategory(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentCountry(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDescription(bytes32 _reference, bytes32 _docRef) public view returns (string memory);
  function getDocumentLang(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
  function getDocumentDate(bytes32 _reference, bytes32 _docRef) public view returns (bytes32);
}
