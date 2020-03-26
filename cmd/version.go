package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vs4vijay/vizix/pkg/version"
)

var short bool

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.New().FormatVersion(short))
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	flags := versionCmd.Flags()
	flags.BoolVarP(&short, "short", "s", false, "shorten output version")
}
