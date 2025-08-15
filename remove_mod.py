from Utils.menu import menu
from Utils.get_mc_version import get_mc_version
from Utils.get_launcher import get_launcher
from Utils.run_packwiz_cmd import run_packwiz_cmd
import os

def remove_mod():
    minecraft_versions = get_mc_version()
    if minecraft_versions is None:
        return

    launchers = get_launcher()
    if launchers is None:
        return

    mod_name = input("Enter the mod/resource pack/shader name: ")

    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = os.path.join(minecraft_version, launcher)
            if not os.path.isdir(modpack_dir):
                print(f"Directory {modpack_dir} does not exist.")
                continue

            run_packwiz_cmd(['packwiz', 'remove', mod_name], modpack_dir)