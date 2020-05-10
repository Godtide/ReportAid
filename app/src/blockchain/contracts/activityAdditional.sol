pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./ITree.sol";
import "./IData.sol";
import "./INode.sol";
import "./Types.sol";
import "./Mapping.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./activity.sol";

contract ActivityAdditionalNode is INode {

    AdditionalData data;

    constructor (AdditionalData memory _additional) public {
        require (_additional.scope > uint8(Scope.NONE) &&
                 _additional.scope < uint8(Scope.MAX) &&
                 _additional.capitalSpend <= 100 &&
                 _additional.collaborationType > uint8(CollaborationType.NONE) &&
                 _additional.collaborationType < uint8(CollaborationType.MAX) &&
                 _additional.defaultAidType[0] != 0 &&
                 _additional.defaultTiedStatus > uint8(TiedStatus.NONE) &&
                 _additional.defaultTiedStatus < uint8(TiedStatus.MAX));

        data.scope = _additional.scope;
        data.capitalSpend = _additional.capitalSpend;
        data.collaborationType = _additional.collaborationType;
        data.defaultTiedStatus = _additional.defaultTiedStatus;
        data.defaultFlowType = _additional.defaultFlowType;
        data.defaultAidType = _additional.defaultAidType;
        data.defaultFinanceType = _additional.defaultFinanceType;
    }

    function getter() override virtual public view returns (bytes4)
    {
        require(this.getter.address == address(this));
        require(this.get.address == address(this));

        return this.get.selector;
    }

    function get() public view returns (uint8, AdditionalData memory) {
        return (uint8(IATIElement.ACTIVITYADDITIONAL), data);
    }
}

contract ActivityAdditional is IData {

    Activity activityCon;
    Data store;
    using IterableData for Data;

    constructor(address _activityCon) public {
        assert(_activityCon != address(0x0));

        activityCon = Activity(_activityCon);
        store.size = 0;
    }

    function setter() override virtual public view returns (bytes4)
    {
        require(this.setter.address == address(this));
        require(this.set.address == address(this));

        return this.set.selector;
    }

    function set(bytes32 _parentRef, bytes32 _thisRef,  AdditionalData memory _additional) public {
        require (_parentRef[0] != 0);
        require (_thisRef[0] != 0);

        ActivityAdditionalNode data = new ActivityAdditionalNode(_additional);
        store.insert(_thisRef, address(data));
        activityCon.addChild(_parentRef, address(data), uint8(IATIElement.ACTIVITYADDITIONAL));

        emit Set(uint8(IATIElement.ACTIVITYADDITIONAL), _thisRef);
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
