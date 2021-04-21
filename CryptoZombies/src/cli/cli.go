package main

import (
	"fmt"
	"math/big"
	"strconv"

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
			auth, err := getAuth(key)
			if err != nil {
				return err
			}
			contract, err := getContract(cfg)
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
	getCMD = &cobra.Command{
		Use:   "get",
		Short: "get <index>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("invalid index")
			}
			index, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			cfg, err := rootCMD.Flags().GetString("config")
			if err != nil {
				return err
			}
			contract, err := getContract(cfg)
			if err != nil {
				return err
			}
			zombie, err := contract.Zombies(nil, big.NewInt(int64(index)))
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", zombie)
			return nil
		},
	}
)

func init() {
	rootCMD.AddCommand(createCmd, getCMD)

	createCmd.Flags().StringP("name", "n", "zzocker", "name of the zombie")
	rootCMD.PersistentFlags().StringP("config", "c", "./config.yaml", "config file path location")
	rootCMD.PersistentFlags().StringP("key", "k", "./keys/owner", "private key of the client")
}

func main() {
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}

func getContract(cfgPath string) (*stub.ZombieFactory, error) {
	cfg, err := config.ReadConfig(cfgPath)
	if err != nil {
		return nil, err
	}

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

func getAuth(keyPath string) (*bind.TransactOpts, error) {
	key, err := crypto.LoadECDSA(keyPath)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), nil
}
