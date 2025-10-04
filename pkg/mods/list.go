package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// ListMod manages listing mods for multiple versions and launchers
func ListMod(minecraftVersionArg string, launcherArg string) {
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
			fmt.Printf("Listing mods in %s/%s ...\n", v, l)

			cmd := exec.Command("packwiz", "list")
			cmd.Dir = fmt.Sprintf("%s/%s", v, l)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Printf("Error listing mods in %s/%s: %v\n", v, l, err)
			}
		}
	}
}