/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// listModCmd represents the list command
var listModCmd = &cobra.Command{
	Use: "list [minecraft_version] [launcher]",
	Aliases: []string{"ls"},
	Short: "List mods in the modpack",
	Args: cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]

		mods.ListMod(mcVersion, launcher)
},
}

func init() {
	rootCmd.AddCommand(listModCmd)
}
