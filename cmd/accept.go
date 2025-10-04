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
	Use:   "accept-version [minecraft_version] [launcher] [mod]",
	Aliases: []string{"av"},
	Short: "Accept a specific version for the modpack",
	Args:  cobra.ExactArgs(3),
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
