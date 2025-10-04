/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// refreshModCmd represents the refresh command
var refreshModCmd = &cobra.Command{
	Use:   "refresh [minecraft_version] [launcher]",
	Aliases: []string{"rf"},
	Short: "Refresh packwiz files in the modpack",
	Args:  cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]

		mods.RefreshMod(mcVersion, launcher)
},
}

func init() {
	rootCmd.AddCommand(refreshModCmd)
}
