package kitti

import (
	"context"
	"math/big"
	"testing"

	"github.com/Zzocker/smart-test/CryptoZombies/util"
)

func TestDeploy(t *testing.T) {
	accounts := util.SetupTestAccounts(1)
	sim := util.SetupTestBlockchain(accounts)
	defer sim.Close()
	///////////////
	address, _, _, err := DeployKitti(accounts[0], sim)
	sim.Commit()
	if err != nil {
		t.Error(err)
	}
	if len(address.Bytes()) == 0 {
		t.Error("address of deployed kitti contract should be non-empty")
	}

	/////// check if contract exists at given addres or not
	code, _ := sim.CodeAt(context.Background(), address, nil)
	if len(code) == 0 {
		t.Errorf("no code present at address %s", address.String())
	}
}

func TestGetKitti(t *testing.T) {
	accounts := util.SetupTestAccounts(1)
	sim := util.SetupTestBlockchain(accounts)
	_, _, contract, _ := DeployKitti(accounts[0], sim)
	sim.Commit()
	//////////////////////////
	gene, err := contract.GetKitti(nil, big.NewInt(5))
	if err != nil {
		t.Error(err)
	}
	// TODO remove this hard-coded
	if gene.Int64() != 6110039845060188 {
		t.Error("gene of kitti with ID = 5 should be 6110039845060188")
	}
}
