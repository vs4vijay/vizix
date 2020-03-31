package main

import (
	"os"
	"time"

	"github.com/vs4vijay/vizix/cmd"

	log "github.com/sirupsen/logrus"
	"github.com/getsentry/sentry-go"
)

var SENTRY_DSN = os.Getenv("SENTRY_DSN")

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: SENTRY_DSN,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	sentry.CaptureMessage("Sentry ran")

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
