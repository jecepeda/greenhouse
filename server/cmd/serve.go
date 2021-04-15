package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		dc, close := getDependencyContainer()
		defer close()

		dc.Serve(viper.GetString("port"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
