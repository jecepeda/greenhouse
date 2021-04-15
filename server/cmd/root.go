package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()

	rootCmd.Flags().String("jwt_seed_key", "fake seed key", "The seed key needed to generate jwt keys")
	rootCmd.Flags().String("port", "4000", "The port where the server will run")
	rootCmd.Flags().String("db_host", "", "")
	rootCmd.Flags().Int("db_port", 5432, "")
	rootCmd.Flags().String("db_user", "", "")
	rootCmd.Flags().String("db_password", "", "")
	rootCmd.Flags().String("db_name", "", "")
}
