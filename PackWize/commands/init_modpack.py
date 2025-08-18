from pathlib import Path
import shutil
import re
from PackWize.utils.run_cmd import run_cmd

MODRINTH = ["modrinth", "mr"]
CURSEFORGE = ["curseforge", "cf"]
LAUNCHERS = MODRINTH + CURSEFORGE + ["all"]

def init_modpack():
    while True:
        minecraft_version = input("PackWize: Enter Minecraft version: ")
        
        # Check if the version is valid
        if re.match(r'^\d+(\.\d+){1,2}$', minecraft_version):
            break
        else:
            print("PackWize: Invalid Minecraft version. Example: 1.21, 1.21.1")
    
    while True:
        launcher = input("PackWize: Enter launcher (All (default), Modrinth or CurseForge): ").lower()
        
        if launcher in LAUNCHERS:
            break
        elif launcher == "":
            launcher = "all"
            break
        else:
            print("PackWize: Invalid launcher. Please enter 'Modrinth', 'CurseForge', or 'All'.")

    # Create the modpack directory and initialize it
    if launcher == "all":
        launchers = ["CurseForge", "Modrinth"]

        for launcher in launchers:
            modpack_dir = Path(f"{minecraft_version}/{launcher}")
            modpack_dir.mkdir(parents=True, exist_ok=True)

        run_cmd(['packwiz', 'init'], modpack_dir)

        source_index = Path(f"{modpack_dir}/index.toml")
        source_pack = Path(f"{modpack_dir}/pack.toml")

        modpack_dir_dest = Path(f"{minecraft_version}/{launchers[0]}")
        dest_index = modpack_dir_dest / "index.toml"
        dest_pack = modpack_dir_dest / "pack.toml"

        if source_index.exists() and source_pack.exists():
            shutil.copy(source_index, dest_index)
            shutil.copy(source_pack, dest_pack)
            print(f"PackWize: Copied index.toml and pack.toml to {modpack_dir_dest}")
        else:
            print("PackWize: index.toml or pack.toml not found in current directory")

    else:
        if launcher in MODRINTH:
            launcher = "Modrinth"
        elif launcher in CURSEFORGE:
            launcher = "CurseForge"
        else:
            print("PackWize: Invalid launcher. Please enter 'Modrinth', 'CurseForge', or 'All'.")
            return
        
        modpack_dir = Path(f"{minecraft_version}/{launcher}")
        modpack_dir.mkdir(parents=True, exist_ok=True)

        run_cmd(['packwiz', 'init'], modpack_dir)


# Copier les fichiers pack.toml et index.toml dans le modpack_dir
def copy_packwiz_files(modpack_dir: Path):
    source_index = Path("index.toml")
    source_pack = Path("pack.toml")

    dest_index = modpack_dir / "index.toml"
    dest_pack = modpack_dir / "pack.toml"

    if source_index.exists() and source_pack.exists():
        shutil.copy(source_index, dest_index)
        shutil.copy(source_pack, dest_pack)
        print(f"PackWize: Copied index.toml and pack.toml to {modpack_dir}")
    else:
        print("PackWize: index.toml or pack.toml not found in current directory")

