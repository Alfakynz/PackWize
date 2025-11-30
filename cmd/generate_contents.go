/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

var outputDir string
var name string

// generateContentsCmd is the command to generate modpack content
var generateContentsCmd = &cobra.Command{
	Use: "generate [minecraft_version] [launcher]",
	Aliases: []string{"gen"},
	Long:  "Generate all modpack content into modrinth_contents.md and curseforge_contents.md files",
	Short: "Generate all modpack content into Markdown files",
	Args: cobra.MinimumNArgs(2),
	Run: func(c *cobra.Command, args []string) {
		mcVersion := args[0]
		launcher := args[1]

		quietFlag, _ := c.Flags().GetBool("quiet")

		mods.GenerateContents(mcVersion, launcher, outputDir, name, quietFlag)
	},
}

func init() {
	generateContentsCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "", "Output directory")
	generateContentsCmd.Flags().StringVarP(&name, "name", "n", "", "Base name for output files")

	generateContentsCmd.Flags().Bool("quiet", false, "Suppress output messages")
	
	rootCmd.AddCommand(generateContentsCmd)
}
