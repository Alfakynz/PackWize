/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// removeModCmd represents the remove command
var removeModCmd = &cobra.Command{
	Use: "remove [minecraft_version] [launcher] [mod]",
	Aliases: []string{"uninstall", "rm"},
	Short: "Remove a mod to the modpack",
	Args: cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]

		mods.RemoveMod(mcVersion, launcher, mod)
},
}

func init() {
	rootCmd.AddCommand(removeModCmd)
}
