from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def url_add(minecraft_versions, launchers, mod_name, url):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'url', 'add', mod_name, url], modpack_dir)