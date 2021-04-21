package main

import (
	"flag"
	"log"
	"os"

	"github.com/Zzocker/smart-test/CryptoZombies/src/config"
	"github.com/Zzocker/smart-test/CryptoZombies/stub"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	cfgPath = flag.String("cfg", "./config.yaml", "configuration file localtion")
	logger  = log.New(os.Stdout, "[DEPLOY] ", 0)
)

func main() {
	flag.Parse()
	logger.Printf("reading config present at %s", *cfgPath)
	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		panic(err)
	}
	logger.Print("deploying contract")
	err = deploy(cfg)
	if err != nil {
		panic(err)
	}
	
	err = config.UpdateConfig(*cfgPath, cfg)
	if err != nil {
		panic(err)
	}
}

func deploy(cfg *config.Cfg) error {
	logger.Printf("loading contract owner key from %s",cfg.Deploy.OwnerKeyPath)
	key,err := crypto.LoadECDSA(cfg.Deploy.OwnerKeyPath)
	if err!=nil{
		return err
	}
	auth := bind.NewKeyedTransactor(key)
	logger.Printf("connecting ethereum node at %s",cfg.EthNode.URL)
	client,err := ethclient.Dial(cfg.EthNode.URL)
	if err!=nil{
		return err
	}
	logger.Println("Deploying the contract")
	address,_,_,err := stub.DeployZombieFactory(&bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
	},client)
	if err!=nil{
		return err
	}
	logger.Printf("contract deployed at %s",address.String())
	cfg.Deploy.ContractAddress = address.String()
	return nil
}