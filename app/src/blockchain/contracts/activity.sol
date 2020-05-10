pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./ITree.sol";
import "./IData.sol";
import "./INode.sol";
import "./Types.sol";
import "./Mapping.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./activities.sol";

contract ActivityNode is INode {

    ActivityData data;

    constructor (ActivityData memory _activity) public {
        require (_activity.identifier[0] != 0 &&
                 _activity.reportingOrg.orgRef[0] != 0 &&
                 _activity.reportingOrg.orgType > 0 &&
                 bytes(_activity.title).length > 0 &&
                 _activity.lastUpdated[0] != 0 &&
                 _activity.lang[0]  != 0 &&
                 _activity.currency[0] != 0 &&
                 _activity.hierarchy > uint8(ActivityHierarchy.NONE) &&
                 _activity.hierarchy < uint8(ActivityHierarchy.MAX) &&
                 _activity.status > uint8(ActivityStatus.NONE) &&
                 _activity.status < uint8(ActivityStatus.MAX) &&
                 _activity.budgetNotProvided > uint8(BudgetNotProvided.NONE) &&
                 _activity.budgetNotProvided < uint8(BudgetNotProvided.MAX) &&
                 bytes(_activity.description).length > 0
        );

        data.identifier = _activity.identifier;
        data.reportingOrg.isSecondary = _activity.reportingOrg.isSecondary;
        data.reportingOrg.orgRef = _activity.reportingOrg.orgRef;
        data.reportingOrg.orgType = _activity.reportingOrg.orgType;
        data.title = _activity.title;
        data.lastUpdated = _activity.lastUpdated;
        data.lang = _activity.lang;
        data.currency = _activity.currency;
        data.hierarchy = _activity.hierarchy;
        data.status = _activity.status;
        data.budgetNotProvided = _activity.budgetNotProvided;
        data.description = _activity.description;
    }

    function getter() override virtual public view returns (bytes4)
    {
        require(this.getter.address == address(this));
        require(this.get.address == address(this));

        return this.get.selector;
    }

    function get() public view returns (uint8, ActivityData memory) {
        return (uint8(IATIElement.ACTIVITY), data);
    }
}

contract Activity is IData, ITree {

    Activities activitiesCon;
    Data store;
    Data factory;
    using IterableData for Data;

    constructor(address _activitiesCon) public {
        //require(_activitiesCon != address(0x0));
        activitiesCon = Activities(_activitiesCon);
        store.size = 0;
        factory.size = 0;
    }

    function setter() override virtual public view returns (bytes4)
    {
        require(this.setter.address == address(this));
        require(this.set.address == address(this));

        return this.set.selector;
    }

    function set(bytes32 _parentRef, bytes32 _thisRef,  ActivityData memory _activity) public {
        require (_parentRef[0] != 0);
        require (_thisRef[0] != 0);

        ActivityNode data = new ActivityNode(_activity);
        store.insert(_thisRef, address(data));
        activitiesCon.addChild(_parentRef, address(data), uint8(IATIElement.ACTIVITY));

        Factory warehouse = new Factory();
        factory.insert(_thisRef, address(warehouse));

        emit Set(uint8(IATIElement.ACTIVITY), _thisRef);
    }

    function addChild(bytes32 _ref, address _child, uint8 _childType) override virtual public {
        require (_ref[0] != 0);
        require (_child != address(0x0));
        require (factory.data[_ref].value != address(0x0));

        Factory warehouse = Factory(factory.data[_ref].value);
        warehouse.add(_child);

        emit AddChild(_ref, _child, _childType);
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

    function getFactory(bytes32 _ref) override virtual public view returns (address) {
        require (_ref[0] != 0);
        require (factory.data[_ref].value != address(0x0));

        return factory.data[_ref].value;
    }
}
