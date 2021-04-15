package openauction

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Zzocker/smart-test/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestDeploy(t *testing.T) {
	accounts := utils.SetupAccounts(2)
	sim := utils.SetupBlockchain(accounts, 2e18)
	defer sim.Close()

	endTime := time.Now().Add(1 * time.Minute).Unix()
	address, _, contract, err := DeployOpenAuction(accounts[0], sim, big.NewInt(endTime), accounts[1].From)
	sim.Commit()

	if err != nil {
		t.Error(err)
	}

	if len(address.Bytes()) == 0 {
		t.Errorf("empty contract address")
	}

	// check auction end time
	t.Run("endTime", func(t *testing.T) {
		gotEndTime, err := contract.AuctionEndTime(nil)
		if err != nil {
			t.Error(err)
		}
		if gotEndTime.Int64() != endTime {
			t.Errorf("end time miss-match ; want : %d , but got : %d", endTime, gotEndTime.Int64())
		}
	})

	// beneficiary check
	t.Run("beneficiaryAddress", func(t *testing.T) {
		gotBeneficiary, err := contract.Beneficiary(nil)
		if err != nil {
			t.Error(err)
		}
		if gotBeneficiary != accounts[1].From {
			t.Errorf("beneficiary address miss-match")
		}
	})

}

func TestBid(t *testing.T) {
	accounts := utils.SetupAccounts(2)
	sim := utils.SetupBlockchain(accounts, 2e18)
	defer sim.Close()

	// cannot bid after auction end time
	t.Run("AuctionEnd", func(t *testing.T) {
		endTime := time.Now().Add(1 * time.Second).Unix()
		_, _, contract, _ := DeployOpenAuction(accounts[0], sim, big.NewInt(endTime), accounts[1].From)
		sim.Commit()
		time.Sleep(1 * time.Second)
		_, err := contract.Bid(&bind.TransactOpts{
			From:   accounts[1].From,
			Signer: accounts[1].Signer,
			Value:  nil,
		})
		if err == nil {
			t.Error("bidding should not allowed after end time")
		}
	})

	endTime := time.Now().Add(1 * time.Minute).Unix()
	address, _, contract, _ := DeployOpenAuction(accounts[0], sim, big.NewInt(endTime), accounts[1].From)
	sim.Commit()

	var bidAmount int64 = 100000
	t.Run("FirstBid", func(t *testing.T) {
		_, err := contract.Bid(&bind.TransactOpts{
			From:   accounts[1].From,
			Signer: accounts[1].Signer,
			Value:  big.NewInt(bidAmount),
		})
		sim.Commit()
		if err != nil {
			t.Error(err)
		}
		if want, _ := contract.HighestBid(nil); want.Int64() != bidAmount {
			t.Error("bid amount missmatch")
		}
		if want, _ := contract.HighestBidder(nil); want.Hex() != accounts[1].From.Hex() {
			t.Error("Highest bidder missmatch")
		}
		if want, _ := sim.BalanceAt(context.Background(), address, nil); want.Int64() != bidAmount {
			t.Error("balance of contract doesn't changed")
		}
	})
	t.Run("LowerBid", func(t *testing.T) {
		_, err := contract.Bid(&bind.TransactOpts{
			From:   accounts[1].From,
			Signer: accounts[1].Signer,
			Value:  big.NewInt(bidAmount - 1),
		})
		if err == nil {
			t.Error("cannot bid amount lower then highest bid")
		}
	})
	t.Run("HigherBid", func(t *testing.T) {
		newBid := bidAmount + 1
		_, err := contract.Bid(&bind.TransactOpts{
			From:   accounts[0].From,
			Signer: accounts[0].Signer,
			Value:  big.NewInt(newBid),
		})
		sim.Commit()
		if err != nil {
			t.Error(err)
		}
		if want, _ := contract.HighestBid(nil); want.Int64() != newBid {
			t.Error("bid amount missmatch")
		}
		if want, _ := contract.HighestBidder(nil); want.Hex() != accounts[0].From.Hex() {
			t.Error("Highest bidder missmatch")
		}
		if want, _ := sim.BalanceAt(context.Background(), address, nil); want.Int64() != bidAmount+newBid {
			t.Error("balance of contract doesn't changed")
		}
	})
}

func TestWithdraw(t *testing.T) {
	accounts := utils.SetupAccounts(3)
	sim := utils.SetupBlockchain(accounts, 2e18)
	defer sim.Close()

	endTime := time.Now().Add(1 * time.Minute).Unix()
	address, _, contract, _ := DeployOpenAuction(accounts[0], sim, big.NewInt(endTime), accounts[1].From)
	sim.Commit()

	lowerBid := int64(100000)  // by 1
	higherBid := int64(200000) // by 2

	contract.Bid(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
		Value:  big.NewInt(lowerBid),
	})
	contract.Bid(&bind.TransactOpts{
		From:   accounts[2].From,
		Signer: accounts[2].Signer,
		Value:  big.NewInt(higherBid),
	})
	sim.Commit()
	_, err := contract.Withdraw(&bind.TransactOpts{
		From:   accounts[1].From,
		Signer: accounts[1].Signer,
	})
	sim.Commit()
	if err != nil {
		t.Error(err)
	}
	if want, _ := sim.BalanceAt(context.Background(), address, nil); want.Int64() != higherBid {
		t.Error("balance of contract doesn't changed")
	}

}
