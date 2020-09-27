package cmd

import (
	"github.com/spf13/cobra"
)

var (
	tagMessage string
	rootCmd    = &cobra.Command{
		Use:        "bump major | minor | patch <-m message>",
		Aliases:    nil,
		SuggestFor: nil,
		Short:      "Bump semantic version of a git repository",
		Long:       "",
		Example:    "",
		Version:    Version,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	majorCmd.Flags().StringVarP(&tagMessage, "message", "m", "", "Message for git annotated tag")
	minorCmd.Flags().StringVarP(&tagMessage, "message", "m", "", "Message for git annotated tag")
	patchCmd.Flags().StringVarP(&tagMessage, "message", "m", "", "Message for git annotated tag")
	rootCmd.AddCommand(majorCmd)
	rootCmd.AddCommand(minorCmd)
	rootCmd.AddCommand(patchCmd)
}
