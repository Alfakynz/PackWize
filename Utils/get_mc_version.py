from Utils.menu import menu
from pathlib import Path

def get_mc_version(all = True):
    # Get all Minecraft version directories (folders with digits in name, not hidden)
    folders = sorted([
        f.name
        for f in Path().iterdir()
        if f.is_dir()
        and not f.name.startswith('.')
        and any(c.isdigit() for c in f.name)
    ])

    choices = folders.copy()

    if all and len(folders) > 1:
        choices.insert(0, "All")

    minecraft_version = menu(choices, "Select a Minecraft version:")

    if minecraft_version == "All":
        return folders
    elif minecraft_version is None:
        return None

    return [minecraft_version]