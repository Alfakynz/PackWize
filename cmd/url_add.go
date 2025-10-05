/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// urlAddCmd represents the url add command
var urlAddCmd = &cobra.Command{
	Use: "url [minecraft_version] [launcher] [mod]",
	Aliases: []string{"url-add", "ua"},
	Short: "Add a custom mod (with url) to the modpack",
	Args: cobra.ExactArgs(4),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]
		mod := args[2]
		url := args[3]

		mods.UrlAdd(mcVersion, launcher, mod, url)
},
}

func init() {
	rootCmd.AddCommand(urlAddCmd)
}
