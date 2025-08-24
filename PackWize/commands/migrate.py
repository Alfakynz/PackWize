from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd

def migrate(minecraft_versions, launchers, target, version):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            run_cmd(['packwiz', 'migrate', target, version], modpack_dir)