/*
Copyright © 2025 Alfakynz
*/
package cmd

import (
	"github.com/Alfakynz/PackWize/pkg/mods"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use: "init [modpack_name] [author] [version] [minecraft_version] [modloader] [modloaderVersion] [launcher]",
	Long:  "Init a new modpack with the good directories",
	Short: "Init a new modpack",
	Args: cobra.ExactArgs(7),
	Run: func(c *cobra.Command, args []string) {
		modpackName := args[0]
		author := args[1]
		version := args[2]
		mcVersion := args[3]
		modloader := args[4]
		modloaderVersion := args[5]
		launcher := args[6]

		mods.Init(modpackName, author, version, mcVersion, modloader, modloaderVersion, launcher)
},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
