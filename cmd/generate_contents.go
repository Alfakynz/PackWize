/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// generateContentsCmd is the command to generate modpack content
var generateContentsCmd = &cobra.Command{
	Use: "generate [minecraft_version] [launcher]",
	Aliases: []string{"gen"},
	Short: "Generate all modpack content into a PACK_CONTENT.md file",
	Args: cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]

		mods.GenerateContents(mcVersion, launcher)
},
}

func init() {
	rootCmd.AddCommand(generateContentsCmd)
}
