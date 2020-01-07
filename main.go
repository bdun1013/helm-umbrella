package main

import (
	"os"

	"github.com/bpdunni/helm-umbrella/cmd"
)

func main() {
	umbrellaCmd := cmd.NewRootCmd(os.Stdout, os.Args[1:])

	if err := umbrellaCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
