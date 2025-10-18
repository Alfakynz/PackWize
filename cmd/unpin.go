/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// unpinModCmd represents the unpin command
var unpinModCmd = &cobra.Command{
	Use: "unpin [minecraft_version] [launcher] [mod]",
	Aliases: []string{"unhold", "unlock"},
	Long:  "Unpin a file so it receives updates for the specified Minecraft version and launcher",
	Short: "Unpin a file to allow updates",
	Args: cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]

		mods.UnpinMod(mcVersion, launcher, mod)
},
}

func init() {
	rootCmd.AddCommand(unpinModCmd)
}
