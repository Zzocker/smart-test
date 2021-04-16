package contract

import (
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Zzocker/smart-test/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

func TestDeploy(t *testing.T) {
	accounts := utils.SetupAccounts(1)
	sim := utils.SetupBlockchain(accounts, 1e18)
	defer sim.Close()

	initialToken := int64(1e10)
	name := "INR"
	symbol := "₹"
	decimal := uint8(1)
	address, _, contract, err := DeployERC20(accounts[0], sim, big.NewInt(initialToken), name, symbol, decimal)
	sim.Commit()

	if err != nil {
		t.Error(err)
	}
	if len(address.Bytes()) == 0 {
		t.Error("contract address should not be empty")
	}

	// name check
	if got, _ := contract.Name(nil); got != name {
		t.Error("Token name missmatch")
	}
	// symbol check
	if got, _ := contract.Symbol(nil); got != symbol {
		t.Error("Token symbol missmatch")
	}
	// decimals check
	if got, _ := contract.Decimals(nil); got != decimal {
		t.Error("Token name missmatch")
	}
	// total supply check
	if got, _ := contract.TotalSupply(nil); got.Int64() != initialToken {
		t.Error("Token total supply missmatch")
	}
}

func TestTransfer(t *testing.T) {
	accounts, sim, contract := setup(3)
	defer sim.Close()

	initialToken := int64(1e10)

	t.Run("HappyFlow", func(t *testing.T) {
		transferAmount := int64(1e5)
		_, err := contract.Transfer(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
		}, accounts[1].From, big.NewInt(transferAmount))
		if err != nil {
			t.Error(err)
		}
		sink := make(chan *ERC20Transfer)
		subscription, err := contract.WatchTransfer(nil, sink, []common.Address{accounts[0].From}, nil)
		if err != nil {
			t.Error(err)
		}
		var sinkErr error
		var timeout bool
		var eventLog *ERC20Transfer
		eventWaitTime := 30 * time.Millisecond
		var wg *sync.WaitGroup = new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			select {
			case eventLog = <-sink:
			case sinkErr = <-subscription.Err():
			case <-time.After(eventWaitTime):
				timeout = true
			}
			wg.Done()
		}()
		sim.Commit()
		wg.Wait()
		subscription.Unsubscribe()

		if sinkErr != nil {
			t.Error(sinkErr)
		}

		if timeout {
			t.Error("timeout no event received")
		}

		if eventLog != nil && (eventLog.To.Hex() != accounts[1].From.Hex()) {
			t.Error("event log has different `_to` address")
		}
		if eventLog != nil && (eventLog.Value.Int64() != transferAmount) {
			t.Error("event log has different transfer value")
		}

		if got, _ := contract.Balances(nil, accounts[0].From); got.Int64() != initialToken-transferAmount {
			t.Error("missmatch in balance of from account")
		}
		if got, _ := contract.Balances(nil, accounts[1].From); got.Int64() != transferAmount {
			t.Error("missmatch in balance of to account")
		}
	})

	t.Run("TransferHigherAmount", func(t *testing.T) {
		remainAmountBig, _ := contract.Balances(nil, accounts[0].From)
		remainAmount := remainAmountBig.Int64()
		_, err := contract.Transfer(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
		}, accounts[1].From, big.NewInt(remainAmount+1))
		if err == nil {
			t.Error("should not be able to transfer amount higher then owned token")
		}
	})
}

func TestBalanceOf(t *testing.T) {
	accounts, sim, contract := setup(3)
	defer sim.Close()
	initialToken := int64(1e10)

	if got, _ := contract.BalanceOf(nil, accounts[0].From); got.Int64() != initialToken {
		t.Error("balance of initial account doesn't matche")
	}
}

func TestApprove(t *testing.T) {
	accounts, sim, contract := setup(3)
	defer sim.Close()

	firstAllowance := int64(1e2)
	_, err := contract.Approve(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[2].From, big.NewInt(firstAllowance))
	sim.Commit()
	if err != nil {
		t.Error(err)
	}
	if got, _ := contract.Allowed(nil, accounts[1].From, accounts[2].From); got.Int64() != firstAllowance {
		t.Error("allowance missmatch")
	}
	// TODO add event log check
	secondAllowance := int64(1e2)
	contract.Approve(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[2].From, big.NewInt(secondAllowance))
	sim.Commit()
	if got, _ := contract.Allowed(nil, accounts[1].From, accounts[2].From); got.Int64() != firstAllowance+secondAllowance {
		t.Error("allowance was not added up")
	}
}

func TestAllowance(t *testing.T) {
	accounts, sim, contract := setup(3)
	defer sim.Close()

	allowance := int64(1e2)
	contract.Approve(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	}, accounts[2].From, big.NewInt(allowance))
	sim.Commit()

	if got, _ := contract.Allowance(nil, accounts[1].From, accounts[2].From); got.Int64() != allowance {
		t.Error("allowance missmatch")
	}
}

func TestTransferFrom(t *testing.T) {
	accounts, sim, contract := setup(3)
	defer sim.Close()
	initialToken := int64(1e10)

	t.Run("NoAllowance",func(t *testing.T) {
		_, err := contract.TransferFrom(&bind.TransactOpts{
			From:   accounts[1].From,
			Signer: accounts[1].Signer,
		}, accounts[0].From, accounts[2].From, big.NewInt(1))
		if err==nil{
			t.Error("should not able to tranfer with zero allowance")
		}
	})

	t.Run("SpendHigherThenAllowance",func(t *testing.T) {
		allowance := int64(1e2)
		contract.Approve(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
		}, accounts[2].From, big.NewInt(allowance))
		sim.Commit()
		_, err := contract.TransferFrom(&bind.TransactOpts{
			From:   accounts[2].From,
			Signer: accounts[2].Signer,
		}, accounts[0].From, accounts[1].From, big.NewInt(allowance+1))
		if err==nil{
			t.Error("should not able to tranfer token greater then allowed")
		}
	})

	t.Run("HappyFlow", func(t *testing.T) {
		allowance := int64(1e2)
		contract.Approve(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
		}, accounts[1].From, big.NewInt(allowance))
		sim.Commit()
		_, err := contract.TransferFrom(&bind.TransactOpts{
			From:   accounts[1].From,
			Signer: accounts[1].Signer,
		}, accounts[0].From, accounts[2].From, big.NewInt(allowance-1))
		sim.Commit()
		if err != nil {
			t.Error(err)
		}
		// check balance of `_from`
		if got,_ := contract.Balances(nil,accounts[0].From);got.Int64() != initialToken-allowance+1{
			t.Error("_from account balance missmatch")
		}
		// check balance of `_to`
		if got,_ := contract.Balances(nil,accounts[2].From);got.Int64() != allowance-1{
			t.Error("_to account balance missmatch")
		}
		// check allowance of sender
		if got,_ := contract.Allowed(nil,accounts[0].From,accounts[1].From);got.Int64() != 1{
			t.Error("allowance missmatch")
		}
		// TODO add event test
	})
}

// setup

func setup(count int) ([]*bind.TransactOpts, *backends.SimulatedBackend, *ERC20) {
	accounts := utils.SetupAccounts(count)
	sim := utils.SetupBlockchain(accounts, 1e18)

	initialToken := int64(1e10)
	name := "INR"
	symbol := "₹"
	decimal := uint8(1)
	_, _, contract, _ := DeployERC20(accounts[0], sim, big.NewInt(initialToken), name, symbol, decimal)
	sim.Commit()
	return accounts, sim, contract
}
