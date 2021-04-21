package stub

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestDeploy(t *testing.T) {
	accounts := setupAccounts(1)
	sim := setupBlockchain(accounts)
	defer sim.Close()
	///////////////////////////////

	address, _, _, err := DeployZombieFactory(accounts[0], sim)
	sim.Commit()
	if err != nil {
		t.Error(err)
	}
	if len(address.Bytes()) == 0 {
		t.Error("contract address should not be empty")
	}
}

func TestRandomZombie(t *testing.T){
	accounts := setupAccounts(1)
	sim := setupBlockchain(accounts)
	defer sim.Close()
	_, _, contract, _ := DeployZombieFactory(accounts[0], sim)
	sim.Commit()
	////////////////////////////

	zombieName := "zombie1"
	_,err := contract.CreateRandomZombie(&bind.TransactOpts{
		From: accounts[0].From,
		Signer: accounts[0].Signer,
	},zombieName)
	sim.Commit()
	if err!=nil{
		t.Error(err)
	}
	// Check if factory is updated
	factory,err := contract.Zombies(nil,big.NewInt(0))
	if err!=nil{
		t.Error(err)
	}
	if factory.Name != zombieName{
		t.Errorf("name of the zombie should be same, but got otherwise")
	}
}

func setupAccounts(count int) []*bind.TransactOpts {
	out := make([]*bind.TransactOpts, count)
	for i := 0; i < count; i++ {
		key, _ := crypto.GenerateKey()
		out[i] = bind.NewKeyedTransactor(key)
	}
	return out
}

func setupBlockchain(accounts []*bind.TransactOpts) *backends.SimulatedBackend {
	alloc := make(core.GenesisAlloc)
	for i := range accounts {
		alloc[accounts[i].From] = core.GenesisAccount{
			Balance: big.NewInt(1e18),
		}
	}
	return backends.NewSimulatedBackend(alloc, 1000000)
}
