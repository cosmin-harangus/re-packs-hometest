package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Load the configuration from the command line in a more complex
// scenario I would use Viper to load the configuration from a file
// and the environment variables but in this case it's overkill

var PORT string
var LOCATION string

var rootCmd = &cobra.Command{
	Use:   "shipment-calculator",
	Short: "Shipment calculator is a tool to calculate the minimum number of packs needed to fill an order",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "p", "8080", "Port to listen on")
	rootCmd.PersistentFlags().StringVarP(&LOCATION, "location", "l", "./priv/sizes.json", "Location of the sizes.json file")
	rootCmd.AddCommand(startCmd)
}

// Execute the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Unable to execute command")
	}
}
