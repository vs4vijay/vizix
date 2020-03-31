package main

import (
	"os"
	"time"

	"github.com/vs4vijay/vizix/cmd"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

var SentryDsn = os.Getenv("SENTRY_DSN")

func main() {
	if err := sentry.Init(sentry.ClientOptions{Dsn: SentryDsn}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("Sentry works!")

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
