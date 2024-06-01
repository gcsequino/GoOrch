/*
Copyright © 2024 Greyson Seuqino
*/
package cmd

import (
	"github.com/gcsequino/GoOrch/internal/registry"
	"github.com/spf13/cobra"
)

// registryCmd represents the registry command
var (
	// flag vars
	port string

	registryCmd = &cobra.Command{
		Use:   "registry",
		Short: "Run the registry component of GoOrch",
		Long:  `A longer description of the registry command`,
		Run: func(cmd *cobra.Command, args []string) {
			registry.Hello(port)
		},
	}
)

func init() {
	rootCmd.AddCommand(registryCmd)
	registryCmd.Flags().StringVarP(&port, "port", "p", "55655", "port for the registry to run on")
}
