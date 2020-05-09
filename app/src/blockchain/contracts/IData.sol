pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract IData {

    event Set(string contractName, bytes32 _ref);

    function setter() virtual public view returns (bytes4);
    function getter() virtual public view returns (bytes4);

    function getNum() virtual public view returns (uint256);
    function getReference(uint256 _index) virtual public view returns (bytes32);
}
