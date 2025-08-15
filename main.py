from Utils.menu import menu
from Utils.get_mcv_launchers import get_mcv_launchers
from Functions.add_mod import add_mod
from Functions.remove_mod import remove_mod
from Functions.update_mods import update_mods
from Functions.build_pack_content import build_pack_content
from Functions.export_modpack import export_modpack
from Functions.update_modpack_version import update_modpack_version
from Functions.refresh_modpack import refresh_modpack

def main():
    select = menu(["Add mod", "Remove mod", "Update mods", "Export modpack", "Build pack content", "Update modpack version", "Refresh modpack"], "What do you want to do?")

    if not select:
        return

    minecraft_versions, launchers = get_mcv_launchers()

    if not minecraft_versions or not launchers:
        main()
    else:
        match select:
            case "Add mod":
                mod_name = input("Enter the mod/resource pack/shader name: ")
                add_mod(minecraft_versions, launchers, mod_name)
                main()
            case "Remove mod":
                mod_name = input("Enter the mod/resource pack/shader name: ")
                remove_mod(minecraft_versions, launchers, mod_name)
                main()
            case "Update mods":
                mod_name = input("Enter the mod/resource pack/shader name (--all to update all): ")
                update_mods(minecraft_versions, launchers, mod_name)
                main()
            case "Export modpack":
                export_modpack(minecraft_versions, launchers)
                main()
            case "Build pack content":
                build_pack_content(minecraft_versions, launchers)
                main()
            case "Update modpack version":
                update_modpack_version(minecraft_versions, launchers)
                main()
            case "Refresh modpack":
                refresh_modpack(minecraft_versions, launchers)
                main()
            case None:
                return
            case _:
                print("Invalid selection")

if __name__ == "__main__":
    main()
