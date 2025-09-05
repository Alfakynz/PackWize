from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def url_add(minecraft_versions: list[str], launchers: list[str], mod_name: str, url: str) -> None:
    """
    Add a mod via URL to multiple Minecraft versions and launchers.

    Arguments:
        minecraft_versions: list[str]. List of Minecraft versions
        launchers: list[str]. List of launchers
        mod_name: str. The name of the mod to add
        url: str. The URL of the mod to add

    Returns nothing
    """
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'url', 'add', mod_name, url], modpack_dir)