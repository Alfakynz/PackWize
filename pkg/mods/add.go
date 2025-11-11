package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// AddMod manages adding one or multiple mods for multiple versions and launchers
func AddMod(minecraftVersionArg, launcherArg, modsArg string, forceModrinth, forceCurseforge bool) {
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
				fmt.Printf("Adding %s to %s/%s ...\n", mod, v, l)

				var cmd *exec.Cmd
				if forceModrinth {
					cmd = exec.Command("packwiz", "mr", "add", mod)
				} else if forceCurseforge {
					cmd = exec.Command("packwiz", "cf", "add", mod)
				} else {
					cmd = exec.Command("packwiz", strings.ToLower(l), "add", mod)
				}

				cmd.Dir = fmt.Sprintf("%s/%s", v, l)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin

				if err := cmd.Run(); err != nil {
					log.Printf("Error adding %s to %s/%s: %v\n", mod, v, l, err)
				}
			}
		}
	}
}