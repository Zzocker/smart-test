pragma solidity >0.4.17;

// ERC20Interface : contract interface which needs be implemented 
// a ERC20 token
interface ERC20Interface{

    // totalSupply : retun total supply of token 
    function totalSupply()external view returns(uint256 _totalSupply);
    
    // balanceOf : return account balance of another account with address `_owner`
    function balanceOf(address _owner)external view returns (uint256 _balance);

    // transfer : transfer `_value` amount of token from `msg.sender`'s account to account with address `_to`
    // Will also emit `Transfer` event
    // `throw` if the message caller's account balance does not have enough token to spend.
    // _value == 0 , should also be treadted as a NORAML transfer
    function transfer(address _to,uint _value)external returns(bool success);

    // transferFrom : will transfer `_value` amount of token from a account with address `_from` to account with address `_to`
    // Will also emit `Transfer` event
    // `throw` if account with address `_from` haven't authorized to message sender to tranfer token on his behalf.
    function transferFrom(address _from,address _to,uint256 _value)external returns(bool success);

    // approve : allows `_spender` to withdraw from msg.sender account multiple times, up to the _`value` amount.
    // Will also emit `Approval` event on success
    function approve(address _spender,uint256 _value)external returns (  bool success);

    // allowance : returns amount to token `_spender` is allowed to spend from _owner's account
    function allowance(address _owner, address _spender) external view returns(uint256 remaining);

    // Events
    
    event Transfer(address indexed _from,address indexed _to,uint256 _value);
    event Approval(address indexed _owner,address indexed _to,uint256 _value);
}
