from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def update_mods(minecraft_versions: list[str], launchers: list[str], mod_name: str) -> None:
    """
    Update a specific mod for multiple Minecraft versions and launchers.

    Arguments:
        minecraft_versions: list[str]. List of Minecraft versions
        launchers: list[str]. List of launchers
        mod_name: str. The name of the mod to update

    Returns nothing
    """
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'update', mod_name], modpack_dir)