from Utils.get_mcv_launchers import get_mcv_launchers
from Utils.get_modpack_dir import get_modpack_dir
from Utils.run_cmd import run_cmd

def refresh_modpack():
    minecraft_versions, launchers = get_mcv_launchers()
    if minecraft_versions or launchers is None:
        return
    
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'refresh'], modpack_dir)