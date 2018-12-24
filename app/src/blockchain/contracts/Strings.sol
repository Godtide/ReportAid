pragma solidity ^0.5.0;

library Strings {

  function compare(string memory _a, string memory _b) public pure returns (bool) {

    bytes memory a = bytes(_a);
    bytes memory b = bytes(_b);
    if(bytes(a).length != bytes(b).length) {
        return false;
    } else {
        return keccak256(a) == keccak256(b);
    }
  }

  function bytes32ToStr(bytes32 _bytes32) public constant returns (string memory) {

    bytes memory bytesArray = new bytes(32);
    for (uint256 i; i < 32; i++) {
      bytesArray[i] = _bytes32[i];
    }
    return string(bytesArray);
  }

  // S.Huckle - extra hack to find the index of a string in a string storage array
  function getIndex(string memory _a, bytes32[] memory _store) public view returns (uint256) {

    uint256 index = _store.length;
    for (uint256 i = 0; i < _store.length; i++)
    {
      string memory _b = bytes32ToStr(_store[i]);
      if (compare(_a, _b)) {
        index = i;
        break;
      }
    }
    return index;
  }
}
