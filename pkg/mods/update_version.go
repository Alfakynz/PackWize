package mods

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// UpdateVersion manages the modpack version for multiple versions and launchers
func UpdateVersion(minecraftVersionArg, launcherArg, version string) {
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
			fmt.Printf("Updating version in %s/%s ...\n", v, l)

			// Define path to pack.toml
			packPath := fmt.Sprintf("%s/%s/pack.toml", v, l)

			// Read file content
			data, err := os.ReadFile(packPath)
			if err != nil {
				log.Printf("Error reading %s: %v\n", packPath, err)
				continue
			}

			// Replace the version line using regex
			re := regexp.MustCompile(`(?m)^version\s*=\s*".*"$`)
			updatedData := re.ReplaceAllString(string(data), fmt.Sprintf(`version = "%s"`, version))

			// Write back to file
			if err := os.WriteFile(packPath, []byte(updatedData), 0644); err != nil {
				log.Printf("Error writing %s: %v\n", packPath, err)
				continue
			}

			fmt.Printf("Version updated to %s in %s\n", version, packPath)
		}
	}
}