/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// exportModCmd represents the export command
var exportModCmd = &cobra.Command{
	Use: "export [minecraft_version] [launcher] [mod]",
	Aliases: []string{"ex"},
	Short: "Export the modpack to a mrpack file (Modrinth) or zip file (CurseForge)",
	Args: cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]

		mods.ExportMod(mcVersion, launcher)
},
}

func init() {
	rootCmd.AddCommand(exportModCmd)
}
