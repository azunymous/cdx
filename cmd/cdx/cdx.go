package main

import (
	"fmt"
	"github.com/azunymous/cdx/commands"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "cdx",
		Short:   "continous deployment tooling",
		Long:    `Tooling and scripts for continous deployment`,
		Version: "alpha",
	}

	commands.AddCommands(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
