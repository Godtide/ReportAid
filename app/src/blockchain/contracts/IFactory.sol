pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract IFactory {

    function add(address _node) virtual public;
    function getNum() virtual public view returns (uint256);
    function getNode(uint256 _index) virtual public view returns (address);
}
