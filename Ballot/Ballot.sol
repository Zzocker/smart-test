// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

// Voting with delegation
contract Ballot{
    // Voter : represents a single voter 
    struct Voter{
        uint Weight;
        bool Voted;
        address Delegate;
        uint Vote;
    }
    struct Proposal{
        string Name;
        uint VoteCount;
    }
    mapping(address=>Voter) public Voters;
    Proposal[] public Proposals;
    address public Chairperson;
 
    modifier onlyChairperson{
        require(msg.sender == Chairperson);
        _;
    }

    constructor(string[] memory _proposalsName){
        Chairperson = msg.sender;
        Voters[Chairperson].Weight = 1;

        for (uint i = 0;i<_proposalsName.length;i++){
            Proposals.push(Proposal({Name:_proposalsName[i],VoteCount:0}));
        }
    }


    function givenRightToVote(address voter)public onlyChairperson{
        require(Voters[voter].Weight == 0 && Voters[voter].Voted == false);
        Voters[voter].Weight = 1;
    }

    function vote(uint _proposal) public{
        Voter storage voter = Voters[msg.sender];

        require(voter.Weight != 0 && voter.Voted == false);
        voter.Voted = true;
        Proposals[_proposal].VoteCount += voter.Weight;
    }
    

    function delegate(address _to) public {
        Voter storage sender = Voters[msg.sender];
        require(sender.Weight != 0,"you dont have a voting right");
        require(sender.Voted == false,"you have already voted");
        require(msg.sender != _to, "Self-delegation is disallowed.");

        while (Voters[_to].Delegate != address(0)){
            _to = Voters[_to].Delegate;
            require(msg.sender != _to,"found loop");
        }
        
        sender.Voted = true;
        sender.Delegate = _to;
        Voter storage delegate_ = Voters[_to];
        if (delegate_.Voted){
            Proposals[delegate_.Vote].VoteCount+=sender.Weight;
        }else{
            delegate_.Weight +=sender.Weight;
        }
    }

    function winningProposal() public view returns(uint){
        uint max = 0;
        for (uint i=0;i<Proposals.length;i++){
            if (Proposals[i].VoteCount > Proposals[max].VoteCount){
                max = i;
            }
        }
        return max;
    }
}