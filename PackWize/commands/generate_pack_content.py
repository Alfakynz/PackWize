from pathlib import Path
import tomllib

OUTPUT_FILE: str = "PACK_CONTENT.md"

def generate_pack_content(minecraft_versions: list[str], launchers: list[str]) -> None:
    """
    Generate a markdown file listing all mods, resource packs, and shader packs used in the modpack.

    Arguments:
        minecraft_versions: list[str]. List of Minecraft versions
        launchers: list[str]. List of launchers
        
    Returns nothing
    """
    types = [
        ("mods", "Mods"),
        ("resourcepacks", "Resource packs"),
        ("shaderpacks", "Shader packs"),
    ]

    results = {}
    
    for minecraft_version in minecraft_versions:
        print(f"\033[1mGenerating pack content for Minecraft v{minecraft_version}...\033[0m")
        for key, label in types:
            items = list_projects(minecraft_version, key)
            results[key] = items
            print(f"Found {len(items)} {label.lower()} for Minecraft v{minecraft_version}." if items else f"No {label.lower()} found for Minecraft v{minecraft_version}.")
            print()

        output_file = Path(f"{minecraft_version}/{OUTPUT_FILE}")
        with open(output_file, "w", encoding="utf-8") as f:
            f.write("# List of projects used\n\n")
            for _, label in types:
                f.write(f"- [{label} used](#{label.lower().replace(' ', '-')}-used)\n")
            f.write("\n")
            for key, label in types:
                f.write(f"## {label} used\n\n")
                f.write("\n".join(sorted(results[key])))
                f.write("\n\n")
            for key, label in types:
                print(f"Added {len(results[key])} {label.lower()} in {output_file}")
        print(f"Pack content generated in {output_file}\n")

# List projects in the specified Minecraft version and type (mods, resourcepacks, shaderpacks).
def list_projects(minecraft_version, type):
    projects_folder = Path(f"{minecraft_version}/Modrinth/{type}")
    if not projects_folder.exists():
        print(f"The {projects_folder} folder doesn't exist.")
        return []

    projects = []
    for file in projects_folder.glob("*.toml"):
        with open(file, "rb") as f:
            data = tomllib.load(f)
        try:
            name = data["name"]
            project_id = data["update"]["modrinth"]["mod-id"]
            url = f"https://modrinth.com/project/{project_id}"
            projects.append(f"- [{name}]({url})")
        except KeyError as e:
            print(f"Missing key in {file.name}: {e}")

    return sorted(projects)