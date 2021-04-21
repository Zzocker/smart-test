pragma solidity >0.7.0;

/// @notice ZombieFactory creates new zombies with unquie dna and stores then
/// @dev Read learning about solidity start with this simple contract
contract ZombieFactory{
    /// @dev dnaDigits The number of digit use for representing a dna
    uint private constant dnaDigits = 16;
    /// @dev heights dnaDigits sized number
    uint private constant dnaModulus = (10 ** (dnaDigits+1)) -1;

    /// @dev Zombie represents a Zombie
    struct Zombie{
        string name;
        uint dna;
    }
    constructor(){}
    Zombie[] public zombies;
    /// @dev NewZombiew emitted whenever a new zombie is created
    event NewZombie(uint _id,string _name,uint _dna);

    /// @dev _generateRandomDna Generates a random uint number from a string
    function _generateRandomDna(string memory _str)private pure returns(uint){
        uint rand = uint(keccak256(abi.encodePacked(_str)));
        return rand % dnaModulus;
    }

    function _createZombie(string memory _name,uint _dna)private{
        uint id = zombies.length;
        zombies.push(Zombie({name:_name,dna:_dna}));
        emit NewZombie(id,_name,_dna);
    }

    /// @notice createRandomZombie create a new zombie and emit a new zombie created event
    function createRandomZombie(string memory _name) public{
        uint dna = _generateRandomDna(_name);
        _createZombie(_name, dna);
    }
}


