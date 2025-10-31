package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// MigrateModpack migrates a modpack to another Minecraft or loader version
func MigrateModpack(minecraftVersionArg, launcherArg, target, version string) {
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
			fmt.Printf("Migrating the %s version to %s in %s/%s ...\n", target, version, v, l)

			cmd := exec.Command("packwiz", "migrate", target, version)
			cmd.Dir = fmt.Sprintf("%s/%s", v, l)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			if err := cmd.Run(); err != nil {
				log.Printf("Error migrating the modpack in %s/%s: %v\n", v, l, err)
			}
		}
	}
}