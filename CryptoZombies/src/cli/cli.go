package main

import (
	"github.com/Zzocker/smart-test/CryptoZombies/src/config"
	"github.com/Zzocker/smart-test/CryptoZombies/stub"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var (
	rootCMD = &cobra.Command{
		Use:   "zombie",
		Short: "Zombie is a ethereum based game like CryptoKitties",
	}
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "creates a new zombie with given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := rootCMD.Flags().GetString("config")
			if err != nil {
				return err
			}
			key, err := rootCMD.Flags().GetString("key")
			if err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			auth, contract, err := client(cfg, key)
			if err != nil {
				return err
			}
			_, err = contract.CreateRandomZombie(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
			}, name)
			return err
		},
	}
)

func init() {
	rootCMD.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "zzocker", "name of the zombie")
	rootCMD.PersistentFlags().StringP("config", "c", "./config.yaml", "config file path location")
	rootCMD.PersistentFlags().StringP("key", "k", "./keys/owner", "private key of the client")
}

func main() {
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}

func client(cfgPath string, keyPath string) (*bind.TransactOpts, *stub.ZombieFactory, error) {
	cfg, err := config.ReadConfig(cfgPath)
	if err != nil {
		return nil, nil, err
	}
	// read private key of client
	key, err := crypto.LoadECDSA(keyPath)
	if err != nil {
		return nil, nil, err
	}
	auth := bind.NewKeyedTransactor(key)

	client, err := ethclient.Dial(cfg.EthNode.URL)
	if err != nil {
		return nil, nil, err
	}
	contract, err := stub.NewZombieFactory(common.HexToAddress(cfg.Deploy.ContractAddress), client)
	if err != nil {
		return nil, nil, err
	}
	return auth, contract, nil
}
