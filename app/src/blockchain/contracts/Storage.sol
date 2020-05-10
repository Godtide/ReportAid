pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";
import "./Types.sol";
import "./IFactory.sol";

struct IndexedValue {
    uint keyIndex;
    address value;
}

struct KeyFlag {
    bytes32 key;
    bool deleted;
}

struct Data {
    mapping(bytes32 => IndexedValue) data;
    KeyFlag[] keys;
    uint size;
}

contract Factory is IFactory {

    address[] nodes;

    function add(address _node) override virtual public {
        require(_node != address(0x0));

        nodes.push(_node);
    }

    function getNum() override virtual public view returns (uint256) {
        return nodes.length;
    }

    function getNode(uint256 _index) override virtual public view returns (address) {
        require (_index < nodes.length);

        return nodes[_index];
    }
}
