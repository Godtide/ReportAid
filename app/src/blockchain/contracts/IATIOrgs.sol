pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Orgs.sol";
import "./Strings.sol";

contract IATIOrgs is Orgs {

  bytes32[] orgReferences;
  mapping(bytes32 => Org) private organisations;

  event SetOrg(Org _org);

  function setOrg(Org memory _org) public {
    require (_org.orgRef[0] != 0 && bytes(_org.name).length > 0 && bytes(_org.identifier).length > 0);

    organisations[_org.orgRef] = _org;
    if(!getOrgExists(_org.orgRef)) {
      orgReferences.push(_org.orgRef);
    }

    emit SetOrg(_org);
  }

  function getOrgExists(bytes32 _orgRef) public view returns (bool) {
    require (_orgRef[0] != 0);

    bool exists = false;
    if ( !(orgReferences.length == 0) ) {
      uint256 index = Strings.getIndex(_orgRef, orgReferences);
      exists = (index != orgReferences.length);
    }
    return exists;
  }

  function getNumOrgs() public view returns (uint256) {
    return orgReferences.length;
  }

  function getOrgReference(uint256 _index) public view returns (bytes32) {
    require (_index < orgReferences.length);

    return orgReferences[_index];
  }

  function getOrg(bytes32 _orgRef) public view returns (Org memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef];
  }

  function getOrgName(bytes32 _orgRef) public view returns (string memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef].name;
  }

  function getOrgIdentifier(bytes32 _orgRef) public view returns (string memory) {
    require (_orgRef[0] != 0);

    return organisations[_orgRef].identifier;
  }
}
