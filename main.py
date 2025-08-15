from Utils.menu import menu
from Functions.add_mod import add_mod
from Functions.remove_mod import remove_mod
from Functions.update_mods import update_mods
from Functions.build_pack_content import build_pack_content
from Functions.export_modpack import export_modpack
from Functions.update_modpack_version import update_modpack_version
from Functions.refresh_modpack import refresh_modpack

def main():
    select = menu(["Add mod", "Remove mod", "Update mods", "Export modpack", "Build pack content", "Update modpack version", "Refresh modpack"], "What do you want to do?")
    match select:
        case "Add mod":
            add_mod()
            main()
        case "Remove mod":
            remove_mod()
            main()
        case "Update mods":
            update_mods()
            main()
        case "Export modpack":
            export_modpack()
        case "Build pack content":
            build_pack_content()
            main()
        case "Update modpack version":
            update_modpack_version()
            main()
        case "Refresh modpack":
            refresh_modpack()
            main()
        case None:
            return
        case _:
            print("Invalid selection")

if __name__ == "__main__":
    main()
