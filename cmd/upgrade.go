/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// upgradeCmd is the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "update-packwize [version]",
	Aliases: []string{"upgrade-packwize", "up"},
	Short: "Upgrade the CLI to the latest version",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		var version string
		if len(args) == 0 {
			version = "latest"
		} else {
			version = args[0]
		}

		mods.Upgrade(version)
},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
