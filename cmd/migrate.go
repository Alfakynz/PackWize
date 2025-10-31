/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// migrateModpackCmd represents the migrate command
var migrateModpackCmd = &cobra.Command{
	Use: "migrate [minecraft_version] [launcher] [target] [version]",
	Aliases: []string{"mg"},
	Long:  "Migrate your Minecraft and loader versions to newer or older versions",
	Short: "Migrate your Minecraft and loader versions",
	Args: cobra.ExactArgs(4),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		target := args[2]
		version := args[3]

		mods.MigrateModpack(mcVersion, launcher, target, version)
},
}

func init() {
	rootCmd.AddCommand(migrateModpackCmd)
}
