package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func SetupAccounts(count int) []*bind.TransactOpts {
	out := make([]*bind.TransactOpts, count)
	for i := range out {
		key, _ := crypto.GenerateKey()
		out[i] = bind.NewKeyedTransactor(key)
	}
	return out
}

func SetupBlockchain(accounts []*bind.TransactOpts, initialBalance int64) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	for i := range accounts {
		alloc[accounts[i].From] = core.GenesisAccount{
			Balance: big.NewInt(initialBalance),
		}
	}
	return backends.NewSimulatedBackend(alloc, 1000000)
}
