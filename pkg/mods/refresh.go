package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// RefreshMod manages refreshing packwiz files for multiple versions and launchers
func RefreshMod(minecraftVersionArg, launcherArg string, quiet bool) {
	// Convert arguments
	versions := utils.ConvertArguments("minecraft_versions", minecraftVersionArg)
	launchers := utils.ConvertArguments("launchers", launcherArg)

	if versions == nil || launchers == nil {
		log.Printf("Invalid Arguments: %s / %s\n", minecraftVersionArg, launcherArg)
		return
	}

	// Loop over all combinations
	for _, v := range versions {
		for _, l := range launchers {
			if !quiet {
				fmt.Printf("Refreshing packwiz files in %s/%s ...\n", v, l)
			}
			cmd := exec.Command("packwiz", "refresh")
			cmd.Dir = fmt.Sprintf("%s/%s", v, l)
			if !quiet {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}

			if err := cmd.Run(); err != nil {
				log.Printf("Error refreshing packwiz files in %s/%s: %v\n", v, l, err)
			}
		}
	}
}