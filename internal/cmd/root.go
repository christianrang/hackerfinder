package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/christianrang/hackerfinder/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configurationFile string
	configuration     internal.Configuration
	rootCmd           = &cobra.Command{
		Use:     "hackerfinder",
		Short:   "a script for quickly querying VirusTotal and Abuseipdb APIs",
		Version: "v1.5.2",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

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
		viper.SetConfigName("hackerfinder")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config: %s\n", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("unable to decode into struct: %v\n", err)
		os.Exit(1)
	}
}
