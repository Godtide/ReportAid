pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./IData.sol";
import "./INode.sol";
import "./Mapping.sol";
import "./Enums.sol";
import "./Storage.sol";

struct ActivitiesData {
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
}

contract ActivitiesNode is INode {

    ActivitiesData data;

    constructor (ActivitiesData memory _activities) public {
        require (
            _activities.version[0] != 0 &&
            _activities.generatedTime[0] != 0 &&
            _activities.linkedData[0] != 0
        );

        data.version = _activities.version;
        data.generatedTime = _activities.generatedTime;
        data.linkedData = _activities.linkedData;
    }

    function getter() override virtual public view returns (bytes4)
    {
        require(this.getter.address == address(this));
        require(this.get.address == address(this));

        return this.get.selector;
    }

    function get() public view returns (uint8, ActivitiesData memory) {
        return (uint8(IATIElement.ACTIVITIES), data);
    }
}

contract ActivitiesFactory is IData {

    Data store;
    using IterableData for Data;

    constructor() public {
      store.size = 0;
    }

    function setter() override virtual public view returns (bytes4)
    {
        require(this.setter.address == address(this));
        require(this.set.address == address(this));

        return this.set.selector;
    }

    function set(bytes32 _ref, ActivitiesData memory _activities) public {
        require (_ref[0] != 0);

        ActivitiesNode data = new ActivitiesNode(_activities);
        store.insert(_ref, address(data));
        emit Set(uint8(IATIElement.ACTIVITIES), _ref);
    }

    function get(bytes32 _ref) override virtual public view returns (address) {
        require (_ref[0] != 0);
        require (store.data[_ref].value != address(0x0));

        return store.data[_ref].value;
    }

    function getNum() override virtual public view returns (uint256)
    {
        return store.size;
    }

    function getReference(uint256 _index) override virtual public view returns (bytes32)
    {
        require (_index < store.size);

        return store.keys[_index].key;
    }
}
