pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  bytes32[] organisationsRefs;
  mapping(bytes32 =>  Orgs) private organisations;

  function setOrganisations(bytes32 _organisationsRef, Orgs memory _organisations) override virtual public {
    require (_organisationsRef[0] != 0 &&
             _organisations.version[0] != 0 &&
             _organisations.generatedTime[0] != 0);

    organisations[_organisationsRef] = _organisations;

    if (!Strings.getExists(_organisationsRef, organisationsRefs)) {
        organisationsRefs.push(_organisationsRef);
    }

    emit SetOrganisations(_organisationsRef, _organisations);
  }

  function getNumOrganisations() override virtual public view returns (uint256) {
    return organisationsRefs.length;
  }

  function getOrganisationsReference(uint256 _index) override virtual public view returns (bytes32) {
    require (_index < organisationsRefs.length);

    return organisationsRefs[_index];
  }

  function getOrganisations(bytes32 _organisationsRef) override virtual public view returns (Orgs memory) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef];
  }

  function getVersion(bytes32 _organisationsRef) override virtual public view returns (bytes32) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef].version;
  }

  function getGeneratedTime(bytes32 _organisationsRef) override virtual public view returns (bytes32) {
    require (_organisationsRef[0] != 0);

    return organisations[_organisationsRef].generatedTime;
  }
}
