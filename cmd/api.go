package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vs4vijay/vizix/pkg/api"
)

var apiPort string

var apiCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		initLogger(verbosity)
		log.Info("Starting API Server on port ", apiPort)
		api.Start(apiPort)
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)

	flags := apiCmd.Flags()
	flags.StringVarP(&apiPort, "port", "p", "", "port number to listen")
	cobra.MarkFlagRequired(flags, "port")
}
