from Utils.run_cmd import run_cmd
import os

def remove_mod(minecraft_versions, launchers, mod_name):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = os.path.join(minecraft_version, launcher)
            if not os.path.isdir(modpack_dir):
                print(f"Directory {modpack_dir} does not exist.")
                continue

            run_cmd(['packwiz', 'remove', mod_name], modpack_dir)