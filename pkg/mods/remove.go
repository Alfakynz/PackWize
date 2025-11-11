package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// RemoveMod manages removing a mod for multiple versions and launchers
func RemoveMod(minecraftVersionArg, launcherArg, modsArg string) {
	// Convert arguments
	versions := utils.ConvertArguments("minecraft_versions", minecraftVersionArg)
	launchers := utils.ConvertArguments("launchers", launcherArg)
	mods := utils.ConvertArguments("mods", modsArg)

	if versions == nil || launchers == nil || mods == nil {
		log.Printf("Invalid Arguments: %s / %s / %s\n", minecraftVersionArg, launcherArg, modsArg)
		return
	}

	// Loop over all mods
	for _, mod := range mods {
		if mod == "" {
			continue
		}

		// Loop over all combinations
		for _, v := range versions {
			for _, l := range launchers {
				fmt.Printf("Removing %s to %s/%s ...\n", mod, v, l)

				cmd := exec.Command("packwiz", "remove", mod)
				cmd.Dir = fmt.Sprintf("%s/%s", v, l)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					log.Printf("Error removing %s to %s/%s: %v\n", mod, v, l, err)
				}
			}
		}
	}
}