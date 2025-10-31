package mods

import (
	"fmt"
	"os"
	"os/exec"
)

// UpdatePackWize manages the PackWize upgrade process
func UpdatePackWize(version string) {
	if version == "" {
		version = "latest"
	}

	fmt.Println("Upgrading PackWize to version:", version)

	var versionFlag string
	if version != "latest" {
		versionFlag = "@" + version
	} else {
		versionFlag = "@latest"
	}

	cmd := exec.Command("go", "install", "github.com/Alfakynz/PackWize/cmd/packwize"+versionFlag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error upgrading PackWize:", err)
		return
	}

	fmt.Println("Upgrade completed!")
	fmt.Println("Run `packwize --version` to verify your PackWize version.")
}