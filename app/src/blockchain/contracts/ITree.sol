pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

abstract contract ITree {

    event AddChild(bytes32 _ref, address _child, uint8 _childType);

    function addChild(bytes32 _ref, address _child, uint8 _childType) virtual public;
    function getFactory(bytes32 _ref) virtual public view returns (address);
}
