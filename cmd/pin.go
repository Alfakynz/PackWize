/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// pinModCmd represents the pin command
var pinModCmd = &cobra.Command{
	Use: "pin [minecraft_version] [launcher] [mod]",
	Aliases: []string{"hold", "lock"},
	Long:  "Pin a file so it does not get updated automatically for the specified Minecraft version and launcher",
	Short: "Pin a file to prevent automatic updates",
	Args: cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]

		mods.PinMod(mcVersion, launcher, mod)
},
}

func init() {
	rootCmd.AddCommand(pinModCmd)
}
