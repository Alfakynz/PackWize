/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// updatePackWizeCmd is the command to update PackWize to the latest version
var updatePackWizeCmd = &cobra.Command{
	Use: "update-packwize [version]",
	Aliases: []string{"upgrade-packwize", "up"},
	Short: "Upgrade PackWize to the latest version",
	Args: cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		var version string
		if len(args) == 0 {
			version = "latest"
		} else {
			version = args[0]
		}

		mods.UpdatePackWize(version)
},
}

func init() {
	rootCmd.AddCommand(updatePackWizeCmd)
}
