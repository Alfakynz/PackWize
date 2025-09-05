from PackWize.utils.run_cmd import run_cmd
import os

def remove_mod(minecraft_versions: list[str], launchers: list[str], mod_name: str) -> None:
    """
    Remove a specific mod for multiple Minecraft versions and launchers.

    Arguments:
        minecraft_versions: list[str]. List of Minecraft versions
        launchers: list[str]. List of launchers
        mod_name: str. The name of the mod to remove

    Returns nothing
    """
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = os.path.join(minecraft_version, launcher)
            if not os.path.isdir(modpack_dir):
                print(f"Directory {modpack_dir} does not exist.")
                continue

            run_cmd(['packwiz', 'remove', mod_name], modpack_dir)