package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// AcceptVersion manages accepting a specific version for the modpack for multiple versions and launchers
func AcceptVersion(minecraftVersionArg, launcherArg, version string) {
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
			fmt.Printf("Accepting version %s to %s/%s ...\n", version, v, l)

			cmd := exec.Command("packwiz", "settings", "acceptable-versions", version)
			cmd.Dir = fmt.Sprintf("%s/%s", v, l)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Printf("Error accepting version %s to %s/%s: %v\n", version, v, l, err)
			}
		}
	}
}