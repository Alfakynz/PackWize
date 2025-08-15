from Utils.get_mc_version import get_mc_version
from Utils.get_launcher import get_launcher
import os

def update_modpack_version():
    minecraft_versions = get_mc_version()
    if minecraft_versions is None:
        return
    
    launchers = get_launcher()
    if launchers is None:
        return

    pack_toml_paths = []

    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = os.path.join(minecraft_version, launcher)
            pack_toml_path = os.path.join(modpack_dir, "pack.toml")
            pack_toml_paths.append((launcher, pack_toml_path))

    # Check if at least one pack.toml exists
    found_any = False
    for launcher, pack_toml_path in pack_toml_paths:
        if os.path.isfile(pack_toml_path):
            found_any = True
            break
    if not found_any:
        print("No pack.toml found for Modrinth or CurseForge in the selected version.")
        return

    # Read and display current versions
    current_versions = {}
    for launcher, pack_toml_path in pack_toml_paths:
        if os.path.isfile(pack_toml_path):
            with open(pack_toml_path, "r", encoding="utf-8") as f:
                lines = f.readlines()
            for line in lines:
                if line.strip().startswith("version ="):
                    current_versions[launcher] = line.strip().split("=", 1)[1].strip().strip('"')
                    break
        else:
            current_versions[launcher] = None

    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            if current_versions[launcher] is not None:
                print(f"Current {launcher} modpack version: {current_versions[launcher]}")
            else:
                print(f"pack.toml not found for {launcher}.")

    new_version = input("Enter new modpack version: ").strip()
    if not new_version:
        print("No version entered. Aborting.")
        return

    # Update the version line in both pack.toml files if they exist
    for minecraft_version in minecraft_versions:
        for launcher, pack_toml_path in pack_toml_paths:
            if not os.path.isfile(pack_toml_path):
                continue
            with open(pack_toml_path, "r", encoding="utf-8") as f:
                lines = f.readlines()
            version_line_idx = None
            for idx, line in enumerate(lines):
                if line.strip().startswith("version ="):
                    version_line_idx = idx
                    break
            if version_line_idx is not None:
                lines[version_line_idx] = f'version = "{new_version}"\n'
                with open(pack_toml_path, "w", encoding="utf-8") as f:
                    f.writelines(lines)
                print(f"Modpack version updated to {new_version} in {pack_toml_path}")
            else:
                print(f"No version line found in {pack_toml_path}, skipping.")