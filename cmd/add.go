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
	Use: "add [minecraft_version] [launcher] [mod]",
	Aliases: []string{"install", "a", "i"},
	Long:  "Add a project using a Modrinth or CurseForge URL, slug, project ID, or search term",
	Short: "Add a mod to the modpack",
	Args: cobra.MinimumNArgs(3),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]

		modrinthFlag, _ := c.Flags().GetBool("modrinth")
		curseforgeFlag, _ := c.Flags().GetBool("curseforge")

		mods.AddMod(mcVersion, launcher, mod, modrinthFlag, curseforgeFlag)
	},
}

func init() {
	addModCmd.Flags().Bool("modrinth", false, "Force add mod from Modrinth")
	addModCmd.Flags().Bool("curseforge", false, "Force add mod from CurseForge")

	rootCmd.AddCommand(addModCmd)
}
