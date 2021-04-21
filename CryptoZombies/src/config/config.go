package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type EthNode struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Deploy struct {
	OwnerKeyPath    string `yaml:"ownerKeyPath"`
	ContractAddress string `yaml:"contractAddress"`
}

type Cfg struct {
	EthNode EthNode `yaml:"ethNode"`
	Deploy  Deploy  `yaml:"deploy"`
}

func ReadConfig(filename string) (*Cfg, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg Cfg
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func UpdateConfig(filename string, cfg *Cfg) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewEncoder(file).Encode(cfg)
}
