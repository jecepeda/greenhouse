package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var createDeviceCmd = &cobra.Command{
	Use:   "create-device",
	Short: "creates a new device ready to be used",
	Run: func(cmd *cobra.Command, args []string) {
		dc, close := getDependencyContainer()
		defer close()

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()
		d, err := dc.GetDeviceService().SaveDevice(ctx, name, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(
			"Device:\n\tID: %d\n\tPassword%s\n\tName:%s\n",
			d.ID, password, d.Name,
		)
	},
}

func init() {
	createDeviceCmd.Flags().String("name", "", "the name of the device")
	createDeviceCmd.Flags().String("password", "", "the password to apply")

	rootCmd.AddCommand(createDeviceCmd)
}
