pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./IData.sol";
import "./Mapping.sol";
import "./Enums.sol";
import "./Storage.sol";

struct ActivitiesData {
    bytes32 version;
    bytes32 generatedTime;
    bytes32 linkedData;
}

contract ActivitiesStore {

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

    function get() public view returns (ActivitiesData memory) {
        return data;
    }
}

contract IATIActivities is IData {

    Data store;
    using IterableData for Data;

    constructor() public {
      store.size = 0;
    }

    function setter() override  virtual public view returns (bytes4)
    {
        assert(this.setter.address == address(this));
        assert(this.set.address == address(this));
        return this.set.selector;
    }

    function getter() override  virtual public view returns (bytes4)
    {
        assert(this.getter.address == address(this));
        assert(this.get.address == address(this));
        return this.get.selector;
    }

    function set(bytes32 _ref, ActivitiesData memory _activities) public {

        ActivitiesStore data = new ActivitiesStore(_activities);
        store.insert(_ref, address(data));
        emit Set("Activities", _ref);
    }

    function get(bytes32 _ref) public view returns (ActivitiesData memory) {
        require (_ref[0] != 0);
        require (store.data[_ref].value != address(0x0));

        ActivitiesStore data = ActivitiesStore(store.data[_ref].value);
        return data.get();
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
