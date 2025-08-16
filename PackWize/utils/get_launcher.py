from PackWize.utils.menu import menu
from pathlib import Path

def get_launcher(all = True):
    choices = ["CurseForge", "Modrinth"]

    if all: choices.insert(0, "All")

    launcher = menu(choices, "Select a launcher:")

    if launcher == "All":
        return ["CurseForge", "Modrinth"]
    elif launcher is None:
        return None

    return [launcher]