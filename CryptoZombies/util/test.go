package util

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	initBalance = "100000000000000000000" // 100 eth
)

func SetupTestAccounts(count int) []*bind.TransactOpts {
	keys := make([]*bind.TransactOpts, count)
	for i := range keys {
		key, _ := crypto.GenerateKey()
		keys[i] = bind.NewKeyedTransactor(key)
	}
	return keys
}

func SetupTestBlockchain(accounts []*bind.TransactOpts) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	for i := range accounts {
		balance, _ := new(big.Int).SetString(initBalance, 10)
		alloc[accounts[i].From] = core.GenesisAccount{
			Balance: balance,
		}
	}
	return backends.NewSimulatedBackend(alloc, 10000000)
}
