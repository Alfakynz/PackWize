/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// updateVersionCmd represents the pin command
var updateVersionCmd = &cobra.Command{
	Use:   "update-version [minecraft_version] [launcher] [mod]",
	Aliases: []string{"uv"},
	Short: "Update the modpack version",
	Args:  cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		version := args[2]

		mods.UpdateVersion(mcVersion, launcher, version)
	},
}

func init() {
	rootCmd.AddCommand(updateVersionCmd)
}
