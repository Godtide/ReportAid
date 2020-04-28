pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract Organisation {

  struct ReportingOrg {
    uint8 orgType;
    bool isSecondary;
    bytes32 orgRef;
  }

  struct Org {
    bytes32 orgRef;
    ReportingOrg reportingOrg;
    bytes32 lang;
    bytes32 currency;
    bytes32 lastUpdatedTime;
  }

  event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, Org _org);

  function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, Org memory _org) virtual public;

  function getNumOrganisations(bytes32 _organisationsRef) virtual public view returns (uint256);
  function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) virtual public view returns (bytes32);

  function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (Org memory);

  function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bytes32);
  function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bytes32);
  function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint8);
  function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bool);
  function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bytes32);
  function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bytes32);
  function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (bytes32);
}
