// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

contract Store {
    address private _owner;

    struct mediaFile {
        bytes name;
        bytes32 hash;
        uint32 created;
    }

    mapping(bytes16 => mediaFile) private mediaFiles;

    uint private mapSize;

    constructor()  {
        _owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == _owner, "only the owner can call this function");
        _;
    }

    function changeOwner(address newOwner) public onlyOwner {
        _owner = newOwner;
    }

    function owner() public view returns (address) {
        return _owner;
    }

    function add(bytes16 key, bytes memory name, bytes32 hash, uint32 created) public onlyOwner {
        mediaFiles[key] = mediaFile(name, hash, created);
        mapSize++;
    }

    function get(bytes16 key) public view returns (bytes memory, bytes32, uint32) {
        mediaFile memory file = mediaFiles[key];
        return (file.name, file.hash, file.created);
    }

    function getSize() public view returns (uint) {
        return mapSize;
    }
}
