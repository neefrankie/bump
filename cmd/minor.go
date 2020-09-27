package cmd

import (
	"github.com/neefrankie/bump/pkg/semver"
	"github.com/spf13/cobra"
)

var minorCmd = &cobra.Command{
	Use:                    "minor",
	Aliases:                nil,
	SuggestFor:             nil,
	Short:                  "Increase minor version",
	Long:                   "",
	Example:                "",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = Incr(semver.VerPartMinor, tagMessage, false)
	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}
