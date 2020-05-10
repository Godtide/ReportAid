pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ITree {

    event AddChild(uint8 _elementType, bytes32 _ref);

    function addChild(bytes32 _ref, address _child) virtual public;
    function getFactory(bytes32 _ref) virtual public view returns (address);
}
