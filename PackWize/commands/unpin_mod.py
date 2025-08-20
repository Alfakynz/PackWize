from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def unpin_mod(minecraft_versions, launchers, mod_name):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'unpin', mod_name], modpack_dir)