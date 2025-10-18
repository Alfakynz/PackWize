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
	Use: "export [minecraft_version] [launcher]",
	Aliases: []string{"ex"},
	Long:  "Export the current modpack into a .mrpack file for Modrinth or a .zip file for CurseForge",
	Short: "Export the modpack to a .mrpack (Modrinth) or .zip (CurseForge)",
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
