package cmd

import (
	"fmt"
	"github.com/vs4vijay/vizix/pkg/tcp_server"

	"github.com/spf13/cobra"
)

var port, serverType string

// CHECK: why &server.port can't be used
//type server struct {
//	port string
//}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server on port", port)
		tcp_server.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	flags := serverCmd.Flags()
	flags.StringVarP(&serverType, "type", "t", "tcp", "server type (tcp, udp)")
	flags.StringVarP(&port, "port", "p", "9999", "port number to listen")
	// CHECK: make "port" no. required
}
