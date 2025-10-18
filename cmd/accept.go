/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// acceptVersionCmd represents the settings acceptable-versions command
var acceptVersionCmd = &cobra.Command{
	Use: "accept-version [minecraft_version] [launcher] [version]",
	Aliases: []string{"av"},
	Long:  "Manage the acceptable Minecraft versions for your packs. This must be a comma-separated list of versions, e.g. 1.16.3,1.16.4,1.16.5",
	Short: "Accept specific Minecraft versions for the modpack",
	Args: cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		version := args[2]

		mods.AcceptVersion(mcVersion, launcher, version)
},
}

func init() {
	rootCmd.AddCommand(acceptVersionCmd)
}
