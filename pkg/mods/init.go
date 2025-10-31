package mods

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Alfakynz/PackWize/pkg/utils"
)

// Init creates the directory structure for the modpack version and launchers,
func Init(modpackNameArg, authorArg, versionArg, minecraftVersionArg, modloaderArg, modloaderVersionArg, launcherArg string) {
	// Normalize inputs for case-insensitive comparison
	mcVersionLower := strings.ToLower(minecraftVersionArg)
	modloaderLower := strings.ToLower(modloaderArg)
	modloaderVersionLower := strings.ToLower(modloaderVersionArg)

	// Parse launcher argument (expected formats: "Modrinth", "CurseForge", or both)
	launchers := utils.ConvertArguments("launchers", launcherArg)
	if launchers == nil {
		log.Printf("Invalid Argument: %s\n", launcherArg)
		return
	}

	var baseVersionDir string
	if mcVersionLower == "latest" {
		baseVersionDir = "latest"
	} else {
		baseVersionDir = minecraftVersionArg
	}

	// Create base version directory (e.g., ./1.21.1 or ./latest)
	if err := os.MkdirAll(baseVersionDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create version directory %s: %v", baseVersionDir, err)
	}

	// Create base configurations diectory
	configurationsDir := fmt.Sprintf("%s/configurations", baseVersionDir)
	if err := os.MkdirAll(configurationsDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create configurations directory %s: %v", configurationsDir, err)
	}

	configsDir := fmt.Sprintf("%s/config", configurationsDir)
	if err := os.MkdirAll(configsDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create config directory %s: %v", configsDir, err)
	}

	// Loop through launchers (e.g., Modrinth, CurseForge)
	for _, launcher := range launchers {
		launcherPath := fmt.Sprintf("%s/%s", baseVersionDir, launcher)

		// Create launcher directory (e.g., ./1.21.1/Modrinth or ./latest/Modrinth)
		if err := os.MkdirAll(launcherPath, os.ModePerm); err != nil {
			log.Printf("Failed to create launcher directory %s: %v", launcherPath, err)
			continue
		}

		fmt.Printf("Initializing %s for %s/%s...\n", modpackNameArg, minecraftVersionArg, launcher)

		// Build the packwiz command dynamically
		var cmdArgs []string
		cmdArgs = append(cmdArgs, "init", "--author", authorArg, "--name", modpackNameArg, "--version", versionArg, "--modloader", modloaderArg)

		// Handle Minecraft version argument
		if mcVersionLower == "latest" {
			cmdArgs = append(cmdArgs, "--latest")
		} else {
			cmdArgs = append(cmdArgs, "--mc-version", minecraftVersionArg)
		}

		// Handle modloader arguments
		if modloaderVersionLower == "latest" {
			cmdArgs = append(cmdArgs, fmt.Sprintf("--%s-latest", modloaderLower))
		} else {
			cmdArgs = append(cmdArgs, fmt.Sprintf("--%s-version", modloaderLower), modloaderVersionArg)
		}

		// Execute the command
		cmd := exec.Command("packwiz", cmdArgs...)
		cmd.Dir = launcherPath
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			log.Printf("Error initializing %s for %s/%s: %v\n", modpackNameArg, minecraftVersionArg, launcher, err)
			continue
		}
	}

	// Rename the directory if it is a latest version
	if mcVersionLower == "latest" && len(launchers) > 0 {
		// Attempt to read actual version from one of the launcher's pack.toml files
		launcherPath := fmt.Sprintf("%s/%s", baseVersionDir, launchers[0])
		packTomlPath, err := utils.FindPackToml(launcherPath)
		if err != nil {
			log.Printf("Failed to find pack.toml in %s: %v", launcherPath, err)
			return
		}

		file, err := os.Open(packTomlPath)
		if err != nil {
			log.Printf("Failed to open pack.toml at %s: %v", packTomlPath, err)
			return
		}

		scanner := bufio.NewScanner(file)
		inVersionsSection := false
		actualVersion := ""
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				inVersionsSection = (line == "[versions]")
				continue
			}
			if inVersionsSection && line != "" {
				// Expect line format: key = "value"
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					value = strings.Trim(value, `"`)
					if key == "minecraft" {
						actualVersion = value
						break
					}
				}
			}
		}
		file.Close()

		if actualVersion == "" {
			log.Printf("Could not find minecraft version in pack.toml at %s", packTomlPath)
			return
		}

		err = os.Rename(baseVersionDir, actualVersion)
		if err != nil {
			log.Printf("Failed to rename base version directory from %s to %s: %v", baseVersionDir, actualVersion, err)
			return
		}
	}
}