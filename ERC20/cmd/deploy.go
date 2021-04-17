package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"

	"github.com/Zzocker/smart-test/ERC20/config"
	"github.com/Zzocker/smart-test/ERC20/stub"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	cfgPath = flag.String("config", "./config/config.yaml", "network configuration file")
)

func main() {
	flag.Parse()
	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		log.Printf("error reading config from %s : %v", *cfgPath, err)
		os.Exit(1)
	}
	if err = deploy(cfg); err != nil {
		log.Printf("error deploying smart contract : %v", err)
	}
	if err = config.UpdateConfig(*cfgPath, cfg); err != nil {
		log.Printf("error updating config at %s : %v", *cfgPath, err)
		os.Exit(1)
	}
}

func deploy(cfg *config.Config) error {
	log.Printf("loading private key of contract owner from %s", cfg.Deploy.OwnerKey)
	key, err := crypto.LoadECDSA(cfg.Deploy.OwnerKey)
	if err != nil {
		return err
	}
	log.Printf("dialing ethereum node at %s", cfg.Node.URL)
	client, err := ethclient.Dial(cfg.Node.URL)
	if err != nil {
		return err
	}
	defer client.Close()
	log.Println("getting chainID of network")
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	log.Printf("chainID of network is %v", chainID)

	log.Println("create transitionOpts using contract owner private key")
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return err
	}
	_ = auth
	log.Printf("Deploy contract with initialToken = %d , symbol = %s , name = %s", cfg.Deploy.InitialToken, cfg.Deploy.TokenSymbol, cfg.Deploy.TokenName)
	address, _, _, err := stub.DeployERC20(auth, client, big.NewInt(cfg.Deploy.InitialToken), cfg.Deploy.TokenName, cfg.Deploy.TokenSymbol, cfg.Deploy.TokenDecimal)
	if err!=nil{
		return err
	}
	log.Printf("ERC20 contract deployed at address = %s",address.String())
	cfg.Deploy.ContractAddress = address.String()
	return nil
}
