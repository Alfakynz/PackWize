import os

def get_modpack_dir(minecraft_version: str, launcher: str) -> str | None:
    """
    Get the modpack directory for a specific Minecraft version and launcher

    Arguments:
        minecraft_version: str. The minecraft version of the modpack
        launcher: str. Thr launcher of the modpack

    Returns:
        str: the modpack directory if both inputs are valid
        None: If either input is invalid
    """
    modpack_dir = os.path.join(minecraft_version, launcher)
    if not os.path.isdir(modpack_dir):
        print(f"Directory {modpack_dir} does not exist.")
        return None
  
    return modpack_dir