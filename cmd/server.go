package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vs4vijay/vizix/pkg/tcp"
)

var listenPort, serverType string

// CHECK: why &server.port can't be used
//type server struct {
//	port string
//}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		initLogger(verbosity)
		log.Info("Starting server on port ", listenPort)
		tcp.Start(listenPort)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	flags := serverCmd.Flags()
	flags.StringVarP(&serverType, "type", "t", "tcp", "server type (tcp, udp)")
	flags.StringVarP(&listenPort, "port", "p", "", "port number to listen")
	cobra.MarkFlagRequired(flags, "port")
}
