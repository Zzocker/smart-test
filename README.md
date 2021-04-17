# Smart Test

Smart Test is a collection of ethereum smart contract written in solidity with uint test written in golang.

## Contracts
- [Ballot](#Ballot)
- [ Open Auction](#Open-Auction)
- [ERC20](#ERC20)

## [Ballot](https://docs.soliditylang.org/en/v0.8.3/solidity-by-example.html#voting)

- [Contract](Ballot/Ballot.sol)
- [Uint Test File](Ballot/ballot_test.go)

## [ Open Auction](https://docs.soliditylang.org/en/v0.8.3/solidity-by-example.html#simple-open-auction)

- [Contract](OpenAuction/OpenAuction.sol)
- [Uint Test File](OpenAuction/OpenAuction_test.go)

## [ERC20](https://eips.ethereum.org/EIPS/eip-20)

- Function
    - totalSupply() public view returns (uint256)
    - balanceOf(address tokenOwner) public view returns (uint)
    - allowance(address tokenOwner,address spender) public view returns(uint256)
        - returns amount of token spender is allowed to withdrawal from tokenOwner
    - transfer(address _to, uint _value) public return (bool)
        - transfer `_value` of token to `_to` 
    - approve(address spender,uint _value) public returns (bool)
        - allow `spender` to withdraw `_value` amount of token from owners account
    - transferFrom(address from , address to , uint tokens) public returns (bool)
- Events
    - Approval(address indexed tokenOwner, address indexed spender, uint tokens)
    - Transfer(address indexed from, address indexed to , uint tokens)
- Variable
    - string public constant name
    - string public constant symbol
    - uint8 public constant decimals

### Deploy
- Update config file at `ERC20/config/config.yaml`
    - 
    ```
    node:
        type: ganache
        url: ws://localhost:7545
    deploy:
        contractAddress: 
        ownerKey: keys/contract_owner
        initialToken: 1000000000000000000
        tokenName: INR
        tokenSymbol: ₹
        tokenDecimal: 1  
    ```
    - Run `./bin/deploy`

### Start Event Watcher
Run `./bin/event`