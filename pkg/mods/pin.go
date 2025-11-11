package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// PinMod manages pinning a mod for multiple versions and launchers
func PinMod(minecraftVersionArg, launcherArg, modsArg string) {
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

		// Loop over all combinations of version/launcher
		for _, v := range versions {
			for _, l := range launchers {
				fmt.Printf("Pinning %s to %s/%s ...\n", mod, v, l)

				cmd := exec.Command("packwiz", "pin", mod)
				cmd.Dir = fmt.Sprintf("%s/%s", v, l)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					log.Printf("Error pinning %s to %s/%s: %v\n", mod, v, l, err)
				}
			}
		}
	}
}