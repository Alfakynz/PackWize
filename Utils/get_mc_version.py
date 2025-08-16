from Utils.menu import menu
from pathlib import Path

def get_mc_version(all = True):
    # Get all Minecraft version directories (directories with digits in name, not hidden)
    directories = sorted([
        d.name
        for d in Path().iterdir()
        if d.is_dir()
        and not d.name.startswith('.')
        and any(c.isdigit() for c in d.name)
    ])

    choices = directories.copy()

    if all and len(directories) > 1:
        choices.insert(0, "All")

    minecraft_version = menu(choices, "Select a Minecraft version:")

    if minecraft_version == "All":
        return directories
    elif minecraft_version is None:
        return None

    return [minecraft_version]