package cmd

import (
	"log"

	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configurationFile string
	configuration     Configuration
	rootCmd           = &cobra.Command{
		Use:     "badip",
		Short:   "a script for quickly querying the VirusTotal API",
		Version: "v1.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// TODO: move this somewhere better
type Configuration struct {
	VTConfig  vtsdk.Configuration        `mapstructure:"virustotal"`
	Abuseipdb abuseipdbsdk.Configuration `mapstructure:"abuseaipdb"`
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if configurationFile != "" {
		viper.SetConfigFile(configurationFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/")
		viper.SetConfigName("badip")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config: %s\n", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v\n", err)
	}
}
