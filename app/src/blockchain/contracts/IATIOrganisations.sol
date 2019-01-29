pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Organisations.sol";
import "./Strings.sol";

contract IATIOrganisations is Organisations {

  bytes32[] orgReferences;
  mapping(bytes32 => Organisation) private organisations;

  event SetOrganisation(Organisation _org);

  function setOrganisation(Organisation memory _org) public {
    require (_org.orgRef[0] != 0 && bytes(_org.name).length > 0 && bytes(_org.identifier).length > 0);

    organisations[_org.orgRef] = _org;
    if(!getOrganisationExists(_org.orgRef)) {
      orgReferences.push(_org.orgRef);
    }

    emit SetOrganisation(_org);
  }

  function getOrganisationExists(bytes32 _orgRef) public view returns (bool) {
    require (_orgRef[0] != 0);

    bool exists = false;
    if ( !(orgReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_orgRef, orgReferences);
      exists = (index != orgReferences.length);
    }
    return exists;
  }

  function getNumOrganisations() public view returns (uint256) {
    return orgReferences.length;
  }

  function getOrganisationReference(uint256 _index) public view returns (bytes32) {
    require (_index < orgReferences.length);

    return orgReferences[_index];
  }

  function getOrganisation(bytes32 _orgRef) public view returns (Organisation memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef];
  }

  function getOrganisationName(bytes32 _orgRef) public view returns (string memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef].name;
  }

  function getOrganisationIdentifier(bytes32 _orgRef) public view returns (string memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef].identifier;
  }
}
