package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type EthNode struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Contract struct {
	OwnerKeyPath string `yaml:"ownerKeyPath"`
	Address      string `yaml:"address"`
}

type KittiContract struct {
	Config Contract `yaml:"config"`
	/// constructor arguments can also be added
}

type CFG struct {
	EthNode       EthNode       `yaml:"ethNode"`
	KittiContract KittiContract `yaml:"kittiContract"`
}

func ReadConfig(cfgPath string) (*CFG, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg CFG
	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func UpdateConfig(cfgPath string, cfg *CFG) error {
	file, err := os.OpenFile(cfgPath, os.O_WRONLY, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewEncoder(file).Encode(cfg)
}
