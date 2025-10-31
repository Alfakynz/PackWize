package mods

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// ExportMod manages exporting the modpack for multiple versions and launchers
func ExportMod(minecraftVersionArg, launcherArg string) {
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
			// Copy configurations
			if err := utils.CopyConfigurations(v, l); err != nil {
				log.Printf("Error copying configurations from %s/configurations to %s/%s: %v\n", v, v, l, err)
				continue
			}

			// Export the modpack using packwiz
			fmt.Printf("Exporting the modpack to %s/%s ...\n", v, l)

			cmd := exec.Command("packwiz", strings.ToLower(l), "export")
			cmd.Dir = fmt.Sprintf("%s/%s", v, l)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Printf("Error exporting the modpack to %s/%s: %v\n", v, l, err)
				continue
			}

			outputDir := fmt.Sprintf("dist/%s", v)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
    			log.Printf("Error creating output directory %s: %v\n", outputDir, err)
    			continue
			}

			// Search for .mrpack or .zip files (no subfolders)
			pattern1 := fmt.Sprintf("%s/%s/*.mrpack", v, l)
			pattern2 := fmt.Sprintf("%s/%s/*.zip", v, l)

			filesMrpack, _ := filepath.Glob(pattern1)
			filesZip, _ := filepath.Glob(pattern2)
			exportedFiles := append(filesMrpack, filesZip...)

			if len(exportedFiles) == 0 {
			    fmt.Printf("No exported .mrpack or .zip file found in %s/%s\n", v, l)
			    continue
			}

			// Move each found file
			for _, f := range exportedFiles {
			    dst := filepath.Join(outputDir, filepath.Base(f))
			    if err := os.Rename(f, dst); err != nil {
        			log.Printf("Error moving %s to %s: %v\n", f, dst, err)
        			continue
			    }
			    fmt.Printf("Moved %s --> %s\n", f, dst)
			}
		}
	}
}