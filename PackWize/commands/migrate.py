from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def migrate(minecraft_versions: list[str], launchers: list[str], target: str, version: str) -> None:
    """
    Migrate modpacks to a different format for multiple Minecraft versions and launchers.

    Arguments:
        minecraft_versions: list[str]. List of Minecraft versions
        launchers: list[str]. List of launchers
        target: str. The target format to migrate to (e.g., "modrinth", "curseforge")
        version: str. The version of the target format to migrate to (e.g., "1.0.0")
    
    Returns nothing
    """
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'migrate', target, version], modpack_dir)