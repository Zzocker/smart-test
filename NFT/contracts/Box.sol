pragma solidity >0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Box is Ownable{
    uint256 private value;

    event ValueChange(uint256 newValue);



    function store(uint256 _value)public onlyOwner{
        value = _value;
        emit ValueChange(_value);
    }

    function get()public view returns(uint256){
        return value;
    }
}


