package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:        "bump",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "Bump makes git tagging easier",
	Long:       "",
	Example:    "",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmdMajor)
}
