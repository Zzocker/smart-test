package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy [kitti|zombie] --config <configPath>",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "configuration file location")
	rootCmd.AddCommand(kittiCmd)
}
