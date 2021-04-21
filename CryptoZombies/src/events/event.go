package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Zzocker/smart-test/CryptoZombies/src/config"
	"github.com/Zzocker/smart-test/CryptoZombies/stub"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	newZombiePrintForm = `[NEW_ZOMBIE] {Name : %s , Dna : %v , ID : %v} {BlockNumber : %v , TxHash : %v, BlockHash : %v}`
)

var (
	cfgPath = flag.String("config", "./config.yaml", "config path location")
	logger  = log.New(os.Stdout, "[EVENT] ", 0)
)

func main() {
	flag.Parse()
	logger.Printf("reading config file from %s", *cfgPath)
	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	logger.Printf("connecting zombie contract [address = %s] with ethereum node at %s", cfg.Deploy.ContractAddress, cfg.EthNode.URL)
	contract, err := getContract(cfg)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	doneChannel := make(chan os.Signal)
	signal.Notify(doneChannel, os.Interrupt)
	errChan := make(chan error)

	newZombieSink := make(chan *stub.ZombieFactoryNewZombie)

	go newZombieEvents(contract, newZombieSink, errChan)
	for {
		select {
		case <-doneChannel:
			logger.Println("event logger interrupted. exiting....")
			os.Exit(1)
		case err = <-errChan:
			logger.Println(err)
			os.Exit(1)
		case newZombie := <-newZombieSink:
			logger.Printf(newZombiePrintForm, newZombie.Name, newZombie.Dna, newZombie.Id, newZombie.Raw.BlockNumber, newZombie.Raw.TxHash, newZombie.Raw.BlockHash)
		}
	}
}

func newZombieEvents(contract *stub.ZombieFactory, sink chan<- *stub.ZombieFactoryNewZombie, errChan chan<- error) {
	sub, err := contract.WatchNewZombie(nil, sink)
	if err != nil {
		errChan <- err
		return
	}
	errChan <- (<-sub.Err())
}

func getContract(cfg *config.Cfg) (*stub.ZombieFactory, error) {
	client, err := ethclient.Dial(cfg.EthNode.URL)
	if err != nil {
		return nil, err
	}
	contract, err := stub.NewZombieFactory(common.HexToAddress(cfg.Deploy.ContractAddress), client)
	if err != nil {
		return nil, err
	}
	return contract, nil
}
