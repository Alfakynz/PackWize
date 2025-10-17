package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// AddMod manages adding a mod for multiple versions and launchers
func AddMod(minecraftVersionArg, launcherArg, mod string) {
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
			fmt.Printf("Adding %s to %s/%s ...\n", mod, v, l)

			cmd := exec.Command("packwiz", strings.ToLower(l), "add", mod)
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