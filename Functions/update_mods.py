from Utils.get_modpack_dir import get_modpack_dir
from Utils.run_cmd import run_cmd

def update_mods(minecraft_versions, launchers, mod_name):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'update', mod_name], modpack_dir)