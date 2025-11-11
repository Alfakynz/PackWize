package utils

import (
	"os"
	"slices"
	"sort"
	"strings"
	"unicode"
)

// ConvertArguments converts arguments like "minecraft_versions" or "launchers"
func ConvertArguments(argument, value string) []string {
	switch argument {
	case "minecraft_versions":
		switch strings.ToLower(value) {
		case "all", "a":
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
		case "all", "a":
			return []string{"CurseForge", "Modrinth"}
		case "modrinth", "mr":
			return []string{"Modrinth"}
		case "curseforge", "cf":
			return []string{"CurseForge"}
		}
	
	case "minecraft_versions_init":
		value = strings.ToLower(value)

		// Build a flat ordered list of all known subversions from the minecraftVersions map
		majors := make([]string, 0, len(minecraftVersions))
		for m := range minecraftVersions {
			majors = append(majors, m)
		}

		// sort majors semantically (1.7 before 1.10, etc.)
		sort.Slice(majors, func(i, j int) bool {
			a := strings.Split(majors[i], ".")
			b := strings.Split(majors[j], ".")
			for k := 0; k < len(a) && k < len(b); k++ {
				ai := 0
				bi := 0
				for _, r := range a[k] {
					if r >= '0' && r <= '9' {
						ai = ai*10 + int(r-'0')
					}
				}
				for _, r := range b[k] {
					if r >= '0' && r <= '9' {
						bi = bi*10 + int(r-'0')
					}
				}
				if ai != bi {
					return ai < bi
				}
			}
			return len(a) < len(b)
		})

		flat := []string{}
		for _, m := range majors {
			if vers, ok := minecraftVersions[m]; ok {
				flat = append(flat, vers...)
			}
		}

		// helper: find index of target in flat (exact match)
		findIndex := func(target string) int {
			for i, v := range flat {
				if v == target {
					return i
				}
			}
			return -1
		}

		// helper: resolve a bound string to a concrete version present in flat
		resolveBound := func(b string, isStart bool) string {
			b = strings.TrimSpace(b)
			// case "1.21.x" -> first/last version of that major
			if strings.HasSuffix(b, ".x") {
				major := strings.TrimSuffix(b, ".x")
				if vers, ok := minecraftVersions[major]; ok && len(vers) > 0 {
					if isStart {
						return vers[0]
					}
					return vers[len(vers)-1]
				}
			}
			// case "1.21" -> treat as first/last of major if exists
			if vers, ok := minecraftVersions[b]; ok && len(vers) > 0 {
				if isStart {
					return vers[0]
				}
				return vers[len(vers)-1]
			}
			// direct version like "1.21.4"
			if idx := findIndex(b); idx != -1 {
				return flat[idx]
			}
			// fallback: try to find first version that has this as prefix
			for i, v := range flat {
				if strings.HasPrefix(v, b) {
					if isStart {
						return flat[i]
					}
					// if is end, find last with that prefix
					last := i
					for j := i; j < len(flat) && strings.HasPrefix(flat[j], b); j++ {
						last = j
					}
					return flat[last]
				}
			}
			// not found: return empty string
			return ""
		}

		result := []string{}
		parts := strings.Split(value, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			// range case "start-end"
			if strings.Contains(part, "-") {
				bounds := strings.SplitN(part, "-", 2)
				startStr := strings.TrimSpace(bounds[0])
				endStr := strings.TrimSpace(bounds[1])

				startV := resolveBound(startStr, true)
				endV := resolveBound(endStr, false)
				if startV == "" || endV == "" {
					// skip invalid bounds
					continue
				}
				startIdx := findIndex(startV)
				endIdx := findIndex(endV)
				if startIdx == -1 || endIdx == -1 || startIdx > endIdx {
					continue
				}
				for i := startIdx; i <= endIdx; i++ {
					result = append(result, flat[i])
				}
				continue
			}

			// major.x like "1.21.x"
			if strings.HasSuffix(part, ".x") {
				major := strings.TrimSuffix(part, ".x")
				if vers, ok := minecraftVersions[major]; ok {
					result = append(result, vers...)
				}
				continue
			}

			// plain "1.21" (treat as major) or direct "1.21.4"
			if vers, ok := minecraftVersions[part]; ok {
				result = append(result, vers...)
				continue
			}
			// direct version if present in flat
			if idx := findIndex(part); idx != -1 {
				result = append(result, flat[idx])
				continue
			}
			// otherwise ignore unknown token
		}

		// keep order as in flat and remove duplicates while preserving order
		seen := map[string]struct{}{}
		final := []string{}
		for _, v := range flat {
			if _, want := seen[v]; want {
				// already added
				continue
			}
			// add only if in result slice
			if slices.Contains(result, v) {
				final = append(final, v)
				seen[v] = struct{}{}
			}
		}
		return final

	case "mods":
		// Split mods by comma and trim spaces
		mods := strings.Split(value, ",")
		for i, m := range mods {
			mods[i] = strings.TrimSpace(m)
		}
		return mods

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