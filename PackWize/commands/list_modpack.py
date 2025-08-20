from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def list_modpack(minecraft_versions, launchers):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            print(f"\033[1mMods in {modpack_dir}\033[0m")
            run_cmd(['packwiz', 'list'], modpack_dir)
            print()