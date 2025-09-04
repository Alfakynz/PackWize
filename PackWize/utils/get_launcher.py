from PackWize.utils.menu import menu

def get_launcher(all: bool = True) -> list[str] | None:
    """
    A menu to select a launcher

    Arguments:
        all: bool (default: True). 
            If True, allows selecting from both launchers. 
            If False, only one launcher can be selected

    Returns:
        list[str]: A list containing 'CurseForge' and/or 'Modrinth'
        None: if nothing matches
    """
    choices = ["CurseForge", "Modrinth"]

    if all: choices.insert(0, "All")

    launcher = menu(choices, "Select a launcher:")

    if launcher == "All":
        return ["CurseForge", "Modrinth"]
    elif launcher is None:
        return None

    return [launcher]