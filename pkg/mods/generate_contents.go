package mods

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Alfakynz/PackWize/pkg/utils"
	"github.com/pelletier/go-toml"
)

// GenerateContents generates a markdown file listing all mods, resource packs, and shader packs used in the modpack.
func GenerateContents(minecraftVersionArg, launcherArg string) {
	types := []struct {
		Key   string
		Label string
	}{
		{"mods", "Mods"},
		{"resourcepacks", "Resource packs"},
		{"shaderpacks", "Shader packs"},
	}

	versions := utils.ConvertArguments("minecraft_versions", minecraftVersionArg)
	launchers := utils.ConvertArguments("launchers", launcherArg)

	for _, minecraftVersion := range versions {
		for _, launcher := range launchers {
			fmt.Printf("\033[1mGenerating pack content for Minecraft v%s and launcher %s...\033[0m\n", minecraftVersion, launcher)

			results := make(map[string][]string)

			for _, t := range types {
				items := listProjects(minecraftVersion, t.Key, launcher)
				results[t.Key] = items
				if len(items) > 0 {
					fmt.Printf("Found %d %s for Minecraft v%s and launcher %s.\n", len(items), strings.ToLower(t.Label), minecraftVersion, launcher)
				} else {
					fmt.Printf("No %s found for Minecraft v%s and launcher %s.\n", strings.ToLower(t.Label), minecraftVersion, launcher)
				}
				fmt.Println()
			}

			var fileName string
			switch launcher {
			case "Modrinth":
				fileName = "modrinth_contents.md"
			case "CurseForge":
				fileName = "curseforge_contents.md"
			default:
				fileName = "contents.md"
			}
			outputPath := filepath.Join("dist", minecraftVersion, fileName)
			file, err := os.Create(outputPath)
			if err != nil {
				log.Printf("Error creating %s: %v\n", outputPath, err)
				continue
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("# List of projects used (%s)\n\n", launcher))
			for _, t := range types {
				if len(results[t.Key]) > 0 {
					file.WriteString(fmt.Sprintf("- [%s used](#%s-used)\n", t.Label, strings.ReplaceAll(strings.ToLower(t.Label), " ", "-")))
				}
			}
			file.WriteString("\n")

			for _, t := range types {
				if len(results[t.Key]) == 0 {
					continue
				}
				file.WriteString(fmt.Sprintf("## %s used\n\n", t.Label))
				for _, item := range results[t.Key] {
					file.WriteString(item + "\n")
				}
				file.WriteString("\n")
				fmt.Printf("Added %d %s in %s\n", len(results[t.Key]), strings.ToLower(t.Label), outputPath)
			}

			fmt.Printf("Pack content generated in %s\n\n", outputPath)
		}
	}
}

// listProjects lists projects for the given Minecraft version, type, and launcher (Modrinth/CurseForge).
func listProjects(minecraftVersion, projectType, launcher string) []string {
	modrinthFolder := filepath.Join(minecraftVersion, "Modrinth", projectType)
	curseforgeFolder := filepath.Join(minecraftVersion, "CurseForge", projectType)

	var files []string
	if launcher == "Modrinth" {
		modrinthFiles, _ := filepath.Glob(filepath.Join(modrinthFolder, "*.toml"))
		files = append(files, modrinthFiles...)
	}
	if launcher == "CurseForge" {
		curseforgeFiles, _ := filepath.Glob(filepath.Join(curseforgeFolder, "*.toml"))
		files = append(files, curseforgeFiles...)
	}

	if len(files) == 0 {
		fmt.Printf("No files found in %s or %s.\n", modrinthFolder, curseforgeFolder)
		return []string{}
	}

	var projects []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Error reading %s: %v\n", file, err)
			continue
		}

		tomlTree, err := toml.LoadBytes(data)
		if err != nil {
			log.Printf("Error parsing TOML %s: %v\n", file, err)
			continue
		}

		name := tomlTree.Get("name")
		if name == nil {
			log.Printf("Missing name in %s\n", filepath.Base(file))
			continue
		}

		var url string
		if strings.Contains(file, "Modrinth") {
			projectID := tomlTree.Get("update.modrinth.mod-id")
			if projectID == nil {
				log.Printf("Missing Modrinth project ID in %s\n", filepath.Base(file))
				continue
			}
			url = fmt.Sprintf("https://modrinth.com/project/%v", projectID)
		} else if strings.Contains(file, "CurseForge") {
			fileName := strings.TrimSuffix(filepath.Base(file), ".pw.toml")
			url = fmt.Sprintf("https://www.curseforge.com/minecraft/mc-mods/%s", fileName)
		}

		projects = append(projects, fmt.Sprintf("- [%v](%s)", name, url))
	}

	slices.Sort(projects)
	return projects
}