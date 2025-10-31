package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// findPackToml recursively searches root for a file named pack.toml and returns its full path if found.
func FindPackToml(root string) (string, error) {
	var packTomlPath string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.ToLower(info.Name()) == "pack.toml" {
			packTomlPath = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if packTomlPath == "" {
		return "", fmt.Errorf("pack.toml not found in %s", root)
	}
	return packTomlPath, nil
}