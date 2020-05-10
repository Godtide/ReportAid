pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract INode {

    function getter() virtual public view returns (bytes4);
}
