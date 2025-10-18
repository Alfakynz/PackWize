/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

var allFlag bool

// updateModCmd represents the update command
var updateModCmd = &cobra.Command{
	Use: "update [minecraft_version] [launcher] [mod]",
	Aliases: []string{"upgrade", "u"},
	Long:  "Update an external file (or all external files) in the modpack for the specified Minecraft version and launcher",
	Short: "Update an external file (or all) in the modpack",
	Args: cobra.MaximumNArgs(3),
	Run: func(c *cobra.Command, args []string) {
		var mod string
		if allFlag {
			mod = "--all"
		} else if len(args) == 3 {
			mod = args[2]
		} else {
			c.Help()
			return
		}
		mcVersion := args[0]
		launcher := args[1]

		mods.UpdateMod(mcVersion, launcher, mod)
},
}

func init() {
	updateModCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Update all mods")
	rootCmd.AddCommand(updateModCmd)
}
