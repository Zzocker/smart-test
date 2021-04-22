package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Zzocker/smart-test/CryptoZombies/config"
	"github.com/Zzocker/smart-test/CryptoZombies/stub/kitti"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	kittiCmd = &cobra.Command{
		Use:   "kitti",
		Short: "deploys kitti contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deployKitti(cfgFile)
		},
	}
)

func deployKitti(cfgPath string) error {
	logger := log.New(os.Stdout, "[DEPLOY_KITTI] ", 0)
	logger.Printf("reading config from %s", cfgPath)
	cfg, err := config.ReadConfig(cfgPath)
	if err != nil {
		return err
	}
	// get private of owner
	logger.Printf("reading owner's private key from %s", cfg.KittiContract.Config.OwnerKeyPath)
	key, err := crypto.LoadECDSA(cfg.KittiContract.Config.OwnerKeyPath)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(key)
	/// connect to ethereum node
	logger.Printf("connecting ethereum node at %s", cfg.EthNode.URL)
	client, err := ethclient.Dial(cfg.EthNode.URL)
	if err != nil {
		return err
	}

	/// deploy the contract
	logger.Println("deploying the contract")
	address, _, _, err := kitti.DeployKitti(auth, client)
	if err != nil {
		return err
	}
	if len(address.Bytes()) == 0 {
		return fmt.Errorf("empty kitti contract address")
	}
	/// update contract address
	logger.Printf("updating config file with contract address = %s", address.String())
	cfg.KittiContract.Config.Address = address.String()
	return config.UpdateConfig(cfgPath, cfg)
}
