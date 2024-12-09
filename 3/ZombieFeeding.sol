pragma solidity ^0.4.19;

import "./ZombieFactory.sol";

contract ZombieFeeding is ZombieFactory {

    function feedAndMultiply(uint _zombieId, uint _targetDna) public {
        require(zombieToOwner[_zombieId] == msg.sender);
        Zombie storage myZombie = zombies[_zombieId];
        _targetDna = _targetDna % dnaModulus;
        uint newDna = (myZombie.dna + _targetDna) / 2;
        newDna = newDna - (newDna % 100) + 99;  // 将最后两位设为99
        _createZombie("No-one", newDna);
    }

    function feedOnHuman(uint _zombieId, uint _humanDna) public {
        feedAndMultiply(_zombieId, _humanDna);
    }

    function _catchAHuman(uint _name) internal pure returns (uint) {
        uint rand = uint(keccak256(_name));
        return rand;
    }
}
