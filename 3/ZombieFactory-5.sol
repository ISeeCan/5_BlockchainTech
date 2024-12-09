pragma solidity ^0.4.19;

import "./Ownable-3.sol";

contract ZombieFactory is Ownable {
    struct Zombie {
        string name;
        uint dna;
        uint32 level;
        uint32 readyTime;
    }

    Zombie[] public zombies;
    mapping (uint => address) public zombieToOwner;
    mapping (address => uint) ownerZombieCount;

    uint public levelUpFee = 0.001 ether;

    function levelUp(uint _zombieId) external payable {
        require(msg.value == levelUpFee, "Incorrect Ether amount");
        zombies[_zombieId].level++;
    }

    function withdraw() external onlyOwner {
        owner.transfer(address(this).balance);
    }

    function setLevelUpFee(uint _fee) external onlyOwner {
        levelUpFee = _fee;
    }

    modifier aboveLevel(uint _level, uint _zombieId) {
        require(zombies[_zombieId].level >= _level, "Zombie level too low");
        _;
    }

    function changeName(uint _zombieId, string _newName) external aboveLevel(2, _zombieId) {
        require(msg.sender == zombieToOwner[_zombieId], "Only the owner can change the name");
        zombies[_zombieId].name = _newName;
    }

    function getZombiesByOwner(address _owner) external view returns (uint[]) {
        uint[] memory result = new uint[](ownerZombieCount[_owner]);
        uint counter = 0;
        for (uint i = 0; i < zombies.length; i++) {
            if (zombieToOwner[i] == _owner) {
                result[counter] = i;
                counter++;
            }
        }
        return result;
    }
}