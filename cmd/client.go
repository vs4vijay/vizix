package cmd

import (
	"fmt"
	"github.com/vs4vijay/vizix/pkg/tcp"

	"github.com/spf13/cobra"
)

var serverHost, serverPort string

var clientCmd = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%s:%s", serverHost, serverPort)
		fmt.Println("Connecting to server", address)
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
