from Utils.menu import menu
from Utils.get_mc_version import get_mc_version
from Utils.get_launcher import get_launcher
from Utils.get_modpack_dir import get_modpack_dir
from Utils.run_packwiz_cmd import run_packwiz_cmd

def refresh_modpack():
    minecraft_versions = get_mc_version()
    if minecraft_versions is None:
        return
    launchers = get_launcher()
    if launchers is None:
        return
    
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_packwiz_cmd(['packwiz', 'refresh'], modpack_dir)