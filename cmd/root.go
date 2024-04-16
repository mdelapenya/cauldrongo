package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultConfigFile is the default configuration file
	DefaultConfigFile = ".cauldrongo.yml"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cauldrongo",
	Short: "Cauldron Go is a client for the Cauldron APIs.",
	Long: `A Fast and Flexible Go client for the Cauldron APIs built with
				  love by mdelapenya and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		bs, err := os.ReadFile(cfgFile)
		if err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}

		fmt.Println(string(bs))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", DefaultConfigFile, "config file (default is .cauldrongo.yml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cauldrongo.yml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cauldrongo.yml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
