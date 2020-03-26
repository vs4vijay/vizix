package cmd

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	verbosity int
)

var rootCmd = &cobra.Command{
	Use: "vizix",
	Run: func(cmd *cobra.Command, args []string) {
		initLogger(verbosity)
		cmd.HelpFunc()(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vizix.yaml)")
	pFlags.CountVarP(&verbosity, "verbosity", "v", "set verbosity level")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".vizix")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger(verbosity int) {
	switch {
	case verbosity == 1:
		log.SetLevel(log.DebugLevel)
	case verbosity > 1:
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
