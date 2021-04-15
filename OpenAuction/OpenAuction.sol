// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

contract OpenAuction{
    // about auction
    uint public AuctionEndTime;
    address payable public Beneficiary;
    bool AuctionEnded;


    uint public HighestBid;
    address public HighestBidder;
    mapping(address => uint) PendingReturns;
    
    event HighestBidIncresed(address bidder,uint amount);

    constructor(uint _endTime,address payable _beneficiary){
        AuctionEndTime = _endTime;
        Beneficiary = _beneficiary;
    }

    function Bid() public payable{
        require(block.timestamp <= AuctionEndTime,"Auction has ended");
        require(msg.value >HighestBid,"Higher bid is already placed");
         
        if (HighestBidder != address(0)){
            PendingReturns[HighestBidder]+=HighestBid;
        }
        HighestBid = msg.value;
        HighestBidder = msg.sender;

        emit HighestBidIncresed(msg.sender,msg.value);
    }

    function withdraw()public returns(bool){
        uint amount = PendingReturns[msg.sender];
        if (amount > 0){
            PendingReturns[msg.sender] = 0;
            if (payable(msg.sender).send(amount) == false){
                PendingReturns[msg.sender] = amount;
                return false;
            }
        }
        return true;
    }
}