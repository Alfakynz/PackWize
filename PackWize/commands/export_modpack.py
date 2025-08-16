from PackWize.utils.get_modpack_dir import get_modpack_dir
from PackWize.utils.run_cmd import run_cmd
import sys
import os
import shutil

def export_modpack(minecraft_versions, launchers):
    for minecraft_version in minecraft_versions:
        for launcher in launchers:
            modpack_dir = get_modpack_dir(minecraft_version, launcher)

            config_dir = os.path.join(os.path.dirname(modpack_dir), "configurations")
            files_to_copy = []
            if os.path.exists(config_dir) and os.path.isdir(config_dir):
                for item in os.listdir(config_dir):
                    src_path = os.path.join(config_dir, item)
                    dst_path = os.path.join(modpack_dir, item)
                    files_to_copy.append({
                        "src": src_path,
                        "dst": dst_path,
                        "is_dir": os.path.isdir(src_path),
                        "desc": item
                    })
            else:
                print(f"Configurations directory {config_dir} does not exist.")
                return

            for item in files_to_copy:
                src, dst, is_dir, desc = item["src"], item["dst"], item["is_dir"], item["desc"]
                if not os.path.exists(src):
                    print(f"Source {desc} {'directory' if is_dir else 'file'} {src} does not exist.")
                    return
                try:
                    if is_dir:
                        if os.path.exists(dst):
                            shutil.rmtree(dst)
                        shutil.copytree(src, dst)
                    else:
                        shutil.copy2(src, dst)
                    print(f"Copied {src} to {dst}")
                except Exception as e:
                    print(f"Error copying {desc}: {e}")
                    sys.exit(1)

            # Run the two necessary commands
            run_cmd(['packwiz', 'refresh'], modpack_dir)
            run_cmd(['packwiz', launcher.lower(), 'export'], modpack_dir)