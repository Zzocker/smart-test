pragma solidity >0.4.17;

import "./erc20_interface.sol";
import "./SafeMath.sol";

contract ERC20 is ERC20Interface{
    using SafeMath for uint256;

    // OPTIONAL state variables
    // only for client to make sense of 
    string  public name;
    string public symbol;
    uint8 public decimals;

    uint256 private  _totalSupply;

    mapping(address => uint256) public balances;
    mapping(address => mapping(address => uint256)) public allowed;




    constructor(uint256 _initialToken,string memory _name,string memory _symbol,uint8 _decimals){
        balances[msg.sender] = _initialToken;
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        _totalSupply  = _initialToken;
    }

    function totalSupply()public override view returns(uint256){
        return _totalSupply;
    }  

    function balanceOf(address _owner) public override view returns(uint256){
        return balances[_owner];
    }

    function transfer(address _to,uint _value) public override returns(bool){
        require(_value <=balances[msg.sender]);
        require(_to!= address(0));
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function allowance(address _owner,address _spender) public override view returns(uint256){
        return allowed[_owner][_spender];
    }

    function approve(address _spender,uint256 _value) public override returns(bool){
        require(_spender != address(0));
        allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_value); 
        emit Approval(msg.sender, _spender, _value);
        return true;   
    }

    function transferFrom(address _from,address _to, uint256 _value) public override returns(bool){
        require(_to != address(0));
        require(_value <=balances[_from]);
        require(_value <= allowed[_from][msg.sender]);

        balances[_from] = balances[_from].sub(_value);
        balances[_to]  = balances[_to].add(_value);
        allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
        emit Transfer(_from, _to, _value);
        return true;
    }
}