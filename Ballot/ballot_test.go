package ballot

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

// index = 0 is chairperson
var (
	proposalNames = []string{"proposal_1", "proposal_2", "proposal_3", "proposal_4", "proposal_5"}
)

func TestDeploy(t *testing.T) {
	// only chairperson is required
	accounts := setupAccounts(1)
	sim := setupTestBlockchain(accounts)
	defer sim.Close()

	address, _, contract, err := DeployBallot(accounts[0], sim, proposalNames)
	sim.Commit()
	if err != nil {
		t.Error(err)
	}
	if len(address.Bytes()) == 0 {
		t.Errorf("invalid address of ballot contract")
	}

	// Chairperson address check
	chairpersonAdd, err := contract.Chairperson(nil)
	if err != nil {
		t.Error(err)
	}
	if chairpersonAdd.Hex() != accounts[0].From.Hex() {
		t.Errorf("Chairperson address should be same address who deployed the smart contract")
	}

	// proposals checks
	for i, v := range proposalNames {
		proposal, err := contract.Proposals(nil, big.NewInt(int64(i)))
		if err != nil {
			t.Error(err)
		}
		if proposal.Name != v {
			t.Errorf("Proposal Name are not matching; want = %s but got = %s", v, proposal.Name)
		}
		if proposal.VoteCount.Int64() != 0 {
			t.Errorf("vote count of newly added proposal should be zero , but got %d", proposal.VoteCount.Int64())
		}
	}
}

func TestGiveRightToVoter(t *testing.T) {
	// 0 :Chairperson , 1 : Voter to be added
	accounts := setupAccounts(3)
	sim := setupTestBlockchain(accounts)
	defer sim.Close()

	_, _, contract, _ := DeployBallot(accounts[0], sim, proposalNames)
	sim.Commit()

	// try adding voter with private key not owned by the chairperson
	_, err := contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[2].From,
		Signer: accounts[2].Signer,
		Value:  nil,
	}, accounts[1].From)
	sim.Commit()
	if err == nil {
		t.Error("other user should not be able to give voting right")
	}

	// happy flow
	_, err = contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[0].From,
		Signer: accounts[0].Signer,
		Value:  nil,
	}, accounts[1].From)
	sim.Commit()

	if err != nil {
		t.Errorf("chairperson should be able to give voting rights : %v", err)
	}
	voter, err := contract.Voters(nil, accounts[1].From)
	if err != nil {
		t.Error(err)
	}
	if voter.Weight.Int64() != 1 {
		t.Errorf("newly added voter should have weight = 1, but got otherwise = %d", voter.Weight.Int64())
	}
	if voter.Voted == true {
		t.Error("newly added voter should have voted filed as false, but got ortherwise")
	}

	// try giving right to already existing voter
	_, err = contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[0].From,
		Signer: accounts[0].Signer,
		Value:  nil,
	}, accounts[1].From)
	sim.Commit()
	if err == nil {
		t.Error("adding already existing voter should yield error")
	}
}

func TestVote(t *testing.T) {
	// 0 :Chairperson , 1 : Voter to be added
	accounts := setupAccounts(3)
	sim := setupTestBlockchain(accounts)
	defer sim.Close()

	_, _, contract, _ := DeployBallot(accounts[0], sim, proposalNames)
	sim.Commit()
	contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[0].From,
		Signer: accounts[0].Signer,
		Value:  nil,
	}, accounts[1].From)
	sim.Commit()

	// happy flow
	_, err := contract.Vote(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, big.NewInt(0))
	sim.Commit()

	if err != nil {
		t.Error(err)
	}
	proposal, _ := contract.Proposals(nil, big.NewInt(0))

	if proposal.VoteCount.Int64() != 1 {
		t.Error("vote count of voted proposal should increase")
	}

	// voter should only allowed to vote once
	_, err = contract.Vote(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, big.NewInt(0))
	if err == nil {
		t.Error("voter should only allowed to vote once")
	}

	// user with no voting right should be able to vote
	_, err = contract.Vote(&bind.TransactOpts{
		From:   accounts[2].From,
		Signer: accounts[2].Signer,
	}, big.NewInt(0))

	if err == nil {
		t.Error("voter with non voting right should not be allowed to vote")
	}
}

func TestDelegate(t *testing.T) {
	// 0 :Chairperson , 1 : Voter to be added
	accounts := setupAccounts(3)
	sim := setupTestBlockchain(accounts)
	defer sim.Close()

	_, _, contract, _ := DeployBallot(accounts[0], sim, proposalNames)
	sim.Commit()

	// non-voter not allowed to delegate
	_, err := contract.Delegate(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[0].From)
	if err == nil {
		t.Error("non-voter should not be allowed to delegate")
	}

	contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[0].From,
		Signer: accounts[0].Signer,
		Value:  nil,
	}, accounts[1].From)
	sim.Commit()

	// Self-delegation is disallowed.
	_, err = contract.Delegate(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[1].From)
	if err == nil {
		t.Error("Self-delegation is disallowed.")
	}

	// already voted voter cannot delegate
	contract.Vote(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, big.NewInt(0))
	sim.Commit()

	_, err = contract.Delegate(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[0].From)
	if err == nil {
		t.Error("already voted voter cannot delegate")
	}

	// happy flow
	contract.GivenRightToVote(&bind.TransactOpts{
		From:   accounts[0].From,
		Signer: accounts[0].Signer,
		Value:  nil,
	}, accounts[2].From)
	sim.Commit()
	_, err = contract.Delegate(&bind.TransactOpts{
		From:   accounts[2].From,
		Signer: accounts[2].Signer,
	}, accounts[0].From)
	sim.Commit()
	if err != nil {
		t.Error(err)
	}

	voter, _ := contract.Voters(nil, accounts[2].From)
	if voter.Voted != true {
		t.Errorf("state of delegateing voter should be changed to voted")
	}
	if voter.Delegate != accounts[0].From {
		t.Error("delegating miss-match")
	}
	voter, _ = contract.Voters(nil, accounts[0].From)
	if voter.Weight.Int64() != 2 {
		t.Error("delegated voter's weight should be increase")
	}
}

func TestWinningProposal(t *testing.T) {
	accounts := setupAccounts(4)
	sim := setupTestBlockchain(accounts)
	defer sim.Close()

	_, _, contract, _ := DeployBallot(accounts[0], sim, proposalNames)
	sim.Commit()

	for i := range accounts {
		if i == 0 {
			continue
		}
		contract.GivenRightToVote(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
			Value:  nil,
		}, accounts[i].From)
	}
	sim.Commit()
	const winnerWant = 2
	for i := range accounts {
		contract.Vote(&bind.TransactOpts{
			From:   accounts[i].From,
			Signer: accounts[i].Signer,
		}, big.NewInt(winnerWant))
	}
	sim.Commit()
	got,err := contract.WinningProposal(nil)
	if err!=nil{
		t.Error(err)
	}

	if got.Int64() != winnerWant{
		t.Errorf("proposal_%d should win, but proposal_%d won",winnerWant,got.Int64())
	}
}

func setupAccounts(testAccountCount int) []*bind.TransactOpts {
	testAccounts := make([]*bind.TransactOpts, testAccountCount)

	for i := 0; i < testAccountCount; i++ {
		key, _ := crypto.GenerateKey()
		testAccounts[i] = bind.NewKeyedTransactor((key))
	}
	return testAccounts
}

func setupTestBlockchain(testAccounts []*bind.TransactOpts) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	for i := range testAccounts {
		alloc[testAccounts[i].From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	}
	return backends.NewSimulatedBackend(alloc, 1000000)
}
