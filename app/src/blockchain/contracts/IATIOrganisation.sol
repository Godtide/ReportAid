pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Organisation.sol";
import "./Strings.sol";

contract IATIOrganisation is Organisation {

  bytes32[] organisationsRefs;
  mapping(bytes32 => bytes32[]) private orgRefs;
  mapping(bytes32 =>  mapping(bytes32 => Org)) private organisations;

  event SetOrg(bytes32 _organisationsRef, Org _org);

  function setOrg(bytes32 _organisationsRef, Org memory _org) public {
    require (_organisationsRef[0] != 0 &&
             _org.orgRef[0] != 0 &&
             _org.reportingOrg.orgRef[0] != 0 &&
             _org.reportingOrg.orgType > 0 &&
             _org.lang[0] != 0 &&
             _org.currency[0] != 0 &&
             _org.lastUpdatedTime[0] != 0 );

    organisations[_organisationsRef][_org.orgRef] = _org;

    if (!Strings.getExists(_organisationsRef, organisationsRefs)) {
        organisationsRefs.push(_organisationsRef);
    }

    emit SetOrg(_organisationsRef, _org);
  }

  function getNumOrganisations() public view returns (uint256) {
    return organisationsRefs.length;
  }

  function getNumOrg(bytes32 _organisationsRef) public view returns (uint256) {
    require (_organisationsRef[0] != 0);

    return orgRefs[_organisationsRef].length;
  }

  function getOrganisationsReference(uint256 _index) public view returns (bytes32) {
    require (_index < organisationsRefs.length);

    return organisationsRefs[_index];
  }

  function getOrgReference(bytes32 _organisationsRef, uint256 _index) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _index < orgRefs[_organisationsRef].length);

    return orgRefs[_organisationsRef][_index];
  }

  function getOrg(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (Org memory) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef];
  }

  function getLang(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].lang;
  }

  function getCurrency(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].currency;
  }

  function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].lastUpdatedTime;
  }

  function getReportingOrg(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bytes32) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].reportingOrg.orgRef;
  }

  function getReportingOrgType(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint8) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].reportingOrg.orgType;
  }

  function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (bool) {
    require (_organisationsRef[0] != 0 && _orgRef[0] != 0);

    return organisations[_organisationsRef][_orgRef].reportingOrg.isSecondary;
  }
}
