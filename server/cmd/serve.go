package cmd

import "github.com/spf13/cobra"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve runs the server",
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
