package utils

import (
	"os"
	"sort"
	"strings"
	"unicode"
)

// ConvertArguments converts arguments like "minecraft_versions" or "launchers"
func ConvertArguments(argument, value string) []string {
	switch argument {
	case "minecraft_versions":
		switch strings.ToLower(value) {
		case "all":
			dirs := []string{}

			entries, err := os.ReadDir(".")
			if err != nil {
				return nil
			}

			for _, entry := range entries {
				if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") && containsDigit(entry.Name()) {
					dirs = append(dirs, entry.Name())
				}
			}

			sort.Strings(dirs)
			return dirs
		default:
			return []string{value}
		}

	case "launchers":
		switch strings.ToLower(value) {
		case "all":
			return []string{"CurseForge", "Modrinth"}
		case "modrinth", "mr":
			return []string{"Modrinth"}
		case "curseforge", "cf":
			return []string{"CurseForge"}
		}

	default:
		return nil
	}

	return nil
}

// containsDigit checks if a string contains any digit
func containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}