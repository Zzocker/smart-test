package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"

	"github.com/Zzocker/smart-test/ERC20/config"
	"github.com/Zzocker/smart-test/ERC20/stub"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	cfgPath = flag.String("config", "./config/config.yaml", "network configuration file")
	logger  = log.New(os.Stdout, "[EVENT] ", log.LUTC)
)

func main() {
	logger.Println("reading config from %s", *cfgPath)
	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		logger.Panic(err)
	}
	logger.Printf("dialling ethereum node at %s", cfg.Node.URL)
	client, err := ethclient.Dial(cfg.Node.URL)
	if err != nil {
		logger.Panic(err)
	}
	defer client.Close()
	logger.Printf("connecting erc20 contract at %s", cfg.Deploy.ContractAddress)
	erc20, err := stub.NewERC20(common.HexToAddress(cfg.Deploy.ContractAddress), client)
	if err != nil {
		logger.Panic(err)
	}
	logger.Println("starting transfer event watcher")
	transferSink := make(chan *stub.ERC20Transfer)
	errChan := make(chan error)

	approvalSink := make(chan *stub.ERC20Approval)

	go approvalEvent(erc20, cfg, approvalSink, errChan)
	go transferEvent(erc20, cfg, transferSink, errChan)
	killChan := make(chan os.Signal)

	signal.Notify(killChan, os.Interrupt)
	decimal := big.NewInt(int64(cfg.Deploy.TokenDecimal))

	for {
		select {
		case err := <-errChan:
			logger.Panic(err)
		case <-killChan:
			logger.Println("event watcher killed")
			return
		case transfer := <-transferSink:
			logger.Printf("[TRANSFER] from %v to %v amount %v", transfer.From.String(), transfer.To.String(), new(big.Int).Div(transfer.Value, decimal))
		case approval := <-approvalSink:
			logger.Printf("[APPROVAL] %v token approved by %v to %v", new(big.Int).Div(approval.Value, decimal), approval.Owner.String(), approval.To.String())
		}
	}
}

func transferEvent(erc20 *stub.ERC20, cfg *config.Config, sink chan *stub.ERC20Transfer, errChan chan<- error) {
	sub, err := erc20.WatchTransfer(nil, sink, nil, nil)
	if err != nil {
		errChan <- err
		return
	}
	err = <-sub.Err()
	errChan <- fmt.Errorf("error in transfer event Subscription : %v", err)
}

func approvalEvent(erc20 *stub.ERC20, cfg *config.Config, sink chan *stub.ERC20Approval, errChan chan<- error) {
	sub, err := erc20.WatchApproval(nil, sink, nil, nil)
	if err != nil {
		errChan <- err
		return
	}
	err = <-sub.Err()
	errChan <- fmt.Errorf("error in approval event Subscription : %v", err)
}
