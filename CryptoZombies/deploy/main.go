package main

import (
	"fmt"
	"os"

	"github.com/Zzocker/smart-test/CryptoZombies/deploy/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
