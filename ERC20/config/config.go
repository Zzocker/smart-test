// Package config : keeps all the config representation
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Node : config of ethereum node
type Node struct {
	Type NodeTypes `yaml:"type"`
	URL  string    `yaml:"url"`
}

// NodeTypes : types of node supported
type NodeTypes string

const (
	NODETYPE_GANACHE = "ganache"
)

type Deploy struct {
	ContractAddress string `yaml:"contractAddress"`
	OwnerKey        string `yaml:"ownerKey"`
	InitialToken    int64  `yaml:"initialToken"`
	TokenName       string `yaml:"tokenName"`
	TokenSymbol     string `yaml:"tokenSymbol"`
	TokenDecimal    uint8  `yaml:"tokenDecimal"`
}

type Config struct {
	Node   Node   `yaml:"node"`
	Deploy Deploy `yaml:"deploy"`
}

func ReadConfig(cfgPath string) (*Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func UpdateConfig(cfgPath string, cfg *Config) error {
	file, err := os.OpenFile(cfgPath, os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewEncoder(file).Encode(cfg)
}
