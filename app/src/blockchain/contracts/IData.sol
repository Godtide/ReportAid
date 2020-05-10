pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract IData {

    event Set(uint8 _elementType, bytes32 _ref);

    function setter() virtual public view returns (bytes4);
    function get(bytes32 _ref) virtual public view returns (address);

    function getNum() virtual public view returns (uint256);
    function getReference(uint256 _index) virtual public view returns (bytes32);
}
