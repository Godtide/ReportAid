pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./types.sol";

struct ActivitiesMap {
    mapping(bytes32 => Activities) data;
    KeyFlag[] keys;
    uint size;
}

struct ActivityMap {
    mapping(bytes32 => KeyFlag[]) private keys;
    mapping(bytes32 =>  mapping(bytes32 => Activity)) private data;
    uint size;
}
