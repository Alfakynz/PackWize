from utils.get_modpack_dir import get_modpack_dir
from utils.run_cmd import run_cmd

def add_mod(minecraft_versions, launchers, mod_name):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', launcher.lower(), 'add', mod_name], modpack_dir)