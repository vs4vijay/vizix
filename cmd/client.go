package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vs4vijay/vizix/pkg/tcp"
)

var serverHost, serverPort string

var clientCmd = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		initLogger(verbosity)
		address := fmt.Sprintf("%s:%s", serverHost, serverPort)
		log.Info("Connecting to server", address)
		tcp.Connect(address)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	flags := clientCmd.Flags()
	flags.StringVarP(&serverHost, "host", "H", "", "remote host address")
	flags.StringVarP(&serverPort, "port", "p", "", "port number to connect")
	cobra.MarkFlagRequired(flags, "host")
	cobra.MarkFlagRequired(flags, "port")
}
