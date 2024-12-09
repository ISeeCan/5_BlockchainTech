pragma solidity ^0.4.19;

import "./ZombieFactory-3.sol";

contract ZombieFeeding is ZombieFactory {

    function feedAndMultiply(uint _zombieId, uint _targetDna) internal {
        require(zombieToOwner[_zombieId] == msg.sender);
        Zombie storage myZombie = zombies[_zombieId];
        require(isReady(myZombie));
        _targetDna = _targetDna % dnaModulus;
        uint newDna = (myZombie.dna + _targetDna) / 2;
        newDna = newDna - (newDna % 100) + 99;  // 将最后两位改为99
        _createZombie("NoName", newDna);
        _triggerCooldown(myZombie);
    }

    function feedOnHuman(uint _zombieId, uint _humanDna) public {
        feedAndMultiply(_zombieId, _humanDna);
    }

    function _triggerCooldown(Zombie storage _zombie) internal {
        _zombie.readyTime = uint32(now + cooldownTime);
    }

    function isReady(Zombie storage _zombie) internal view returns (bool) {
        return (_zombie.readyTime <= now);
    }
}
