package mods

import (
	"fmt"
	"os"
	"os/exec"
)

// UpdatePackWize manages the PackWize upgrade process
func UpdatePackWize() {
	fmt.Println("Upgrading PackWize to version: main")

	cmd := exec.Command("go", "install", "github.com/Alfakynz/PackWize/cmd/packwize@main")
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