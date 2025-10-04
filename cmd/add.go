/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// addModCmd represents the add command
var addModCmd = &cobra.Command{
	Use:   "add [minecraft_version] [launcher] [mod]",
	Aliases: []string{"install", "a", "i"},
	Short: "Add a mod to the modpack",
	Args:  cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]

		mods.AddMod(mcVersion, launcher, mod)
},
}

func init() {
	rootCmd.AddCommand(addModCmd)
}
