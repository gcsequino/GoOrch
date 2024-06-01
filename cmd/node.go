/*
Copyright © 2024 Greyson Seuqino
*/
package cmd

import (
	"github.com/gcsequino/GoOrch/internal/node"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var (
	// flag vars
	registryAddr string

	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "Run the node component of GoOrch",
		Long:  `A longer description of the node component.`,
		Run: func(cmd *cobra.Command, args []string) {
			node.Hello(registryAddr)
		},
	}
)

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.Flags().StringVarP(
		&registryAddr,
		"registry",
		"r",
		"localhost:55655",
		"Address and port of the registry to connect to \"<ip>:<port>\"",
	)
}
