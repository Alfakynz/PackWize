import os

def get_modpack_dir(minecraft_version, launcher):
    modpack_dir = os.path.join(minecraft_version, launcher)
    if not os.path.isdir(modpack_dir):
        print(f"Directory {modpack_dir} does not exist.")
        return None
  
    return modpack_dir