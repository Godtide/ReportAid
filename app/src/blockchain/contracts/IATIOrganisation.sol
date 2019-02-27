pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Organisation.sol";
import "./Strings.sol";

contract IATIOrganisation is Organisation {

  mapping(bytes32 => bytes32[]) private orgRefs;
  mapping(bytes32 =>  mapping(bytes32 => Org)) private organisations;

  function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, Org memory _org) public {
    require (_organisationsRef[0] != 0 &&
             _organisationRef[0] != 0 &&
             _org.orgRef[0] != 0 &&
             _org.reportingOrg.orgRef[0] != 0 &&
             _org.reportingOrg.orgType > 0 &&
             _org.lang[0] != 0 &&
             _org.currency[0] != 0 &&
             _org.lastUpdatedTime[0] != 0 );

    organisations[_organisationsRef][_organisationRef] = _org;

    if(!Strings.getExists(_organisationRef, orgRefs[_organisationsRef])) {
      orgRefs[_organisationsRef].push(_organisationRef);
    }

    emit SetOrganisation(_organisationsRef, _organisationRef, _org);
  }

  function getNumOrganisations(bytes32 _organisationsRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0);

    return orgRefs[_organisationsRef].length;
  }

  function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _index < orgRefs[_organisationsRef].length);

    return orgRefs[_organisationsRef][_index];
  }

  function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (Org memory) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef];
  }

  function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].orgRef;
  }

  function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].reportingOrg.orgRef;
  }

  function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].reportingOrg.orgType;
  }

  function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bool) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].reportingOrg.isSecondary;
  }

  function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].lang;
  }

  function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].currency;
  }

  function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _organisationRef[0] != 0);

    return organisations[_organisationsRef][_organisationRef].lastUpdatedTime;
  }
}
