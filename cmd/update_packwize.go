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
	Use:   "update-packwize",
	Aliases: []string{"upgrade-packwize", "up"},
	Long:  "Upgrade PackWize to the latest available version",
	Short: "Upgrade PackWize to the latest version",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mods.UpdatePackWize()
	},
}

func init() {
	rootCmd.AddCommand(updatePackWizeCmd)
}
