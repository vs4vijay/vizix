package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/vs4vijay/vizix/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
