import argparse
from PackWize.utils.menu import menu
from PackWize.utils.get_mcv_launchers import get_mcv_launchers
from PackWize.utils.convert_arguments import convert_arguments
from PackWize.commands.add_mod import add_mod
from PackWize.commands.remove_mod import remove_mod
from PackWize.commands.update_mods import update_mods
from PackWize.commands.accept_version import accept_version
from PackWize.commands.export_modpack import export_modpack
from PackWize.commands.list_modpack import list_modpack
from PackWize.commands.generate_pack_content import generate_pack_content
from PackWize.commands.update_modpack_version import update_modpack_version
from PackWize.commands.refresh_modpack import refresh_modpack
from PackWize.commands.init_modpack import init_modpack

VERSION="1.3.0"

# CLI function
def main():
    parser = argparse.ArgumentParser(
        prog="packwize",
        description="A CLI/TUI to manage modpack easier than just using Packwiz. Based on Packwiz"
    )
    parser.add_argument("-v", "--version", action="version", version=f"Packwize {VERSION}", help="Show the version of PackWize")

    subparsers = parser.add_subparsers(dest="command", required=False, help="Available commands")

    # Parent parser for common arguments
    common_parser = argparse.ArgumentParser(add_help=False)
    common_parser.add_argument("minecraft_versions", help="Minecraft version to work in. Use 'all' to select all Minecraft versions")
    common_parser.add_argument("launchers", help="Launcher to work in. Use 'all' to select all launchers")

    # TUI menu
    parser_tui = subparsers.add_parser("tui", help="Show TUI menu")

    # Add mods
    parser_add = subparsers.add_parser("add", aliases=["install", "i"], parents=[common_parser], help="Add a mod to the modpack")
    parser_add.add_argument("mod", help="Mod/resource pack/shaderpack you want to add")

    # Remove mods
    parser_remove = subparsers.add_parser("remove", aliases=["rm", "uninstall"], parents=[common_parser], help="Remove a mod from the modpack")
    parser_remove.add_argument("mod", help="Mod/resource pack/shaderpack you want to remove")

    # Update mods
    parser_update = subparsers.add_parser("update", aliases=["upgrade"], parents=[common_parser], help="Update mod in the modpack")
    parser_update.add_argument("mod", help="Mod/resource pack/shaderpack you want to update. Use '--all' to update all)")

    # Accept version
    parser_accept_version = subparsers.add_parser("accept-version", aliases=["av"], parents=[common_parser], help="Accept a specific version for the modpack")
    parser_accept_version.add_argument("version", help="Version to accept")

    # Export modpack
    parser_export = subparsers.add_parser("export", aliases=["build"], parents=[common_parser], help="Export the modpack content to a ZIP or MRPACK file. Find the file in the {Minecraft version}/{launcher} directory")

    # List modpack
    parser_list = subparsers.add_parser("list", aliases=["ls"], parents=[common_parser], help="List mods, resource packs and shaders in the modpack")

    # Generate modpack content list
    parser_generate = subparsers.add_parser("generate", aliases=["gen"], parents=[common_parser], help="Export the modpack's content list to an MD file. Find the file in the {Minecraft version} directory")

    # Update modpack version
    parser_update_version = subparsers.add_parser("update-version", aliases=["uv", "set-version", "change-version"], parents=[common_parser], help="Update the modpack version (not the Minecraft version)")

    # Refresh modpack
    parser_refresh = subparsers.add_parser("refresh", aliases=["rf"], parents=[common_parser], help="Refresh the pack.toml and index.toml files")

    # Init modpack
    parser_init = subparsers.add_parser("init", help="Initialize a new modpack and create directories")

    args = parser.parse_args()

    if hasattr(args, "minecraft_versions") and hasattr(args, "launchers"):
        minecraft_versions = convert_arguments("minecraft_versions", args.minecraft_versions)
        launchers = convert_arguments("launchers", args.launchers)

    match args.command:
        case "add" | "install" | "i":
            add_mod(minecraft_versions, launchers, args.mod)
        
        case "remove" | "rm" | "uninstall":
            remove_mod(minecraft_versions, launchers, args.mod)

        case "update" | "upgrade":
            update_mods(minecraft_versions, launchers, args.mod)

        case "accept-version" | "av":
            accept_version(minecraft_versions, launchers, args.version)

        case "export" | "build":
            export_modpack(minecraft_versions, launchers)

        case "list" | "ls":
            list_modpack(minecraft_versions, launchers)

        case "generate" | "gen":
            generate_pack_content(minecraft_versions, launchers)

        case "update-version" | "uv" | "set-version" | "change-version":
            update_modpack_version(minecraft_versions, launchers)

        case "refresh" | "rf":
            refresh_modpack(minecraft_versions, launchers)

        case "init":
            init_modpack()

        case "tui" | _:
            selection()

# TUI function
def selection():
    select = menu(["Add mod", "Remove mod", "Update mods", "Accept version", "Export modpack", "List modpack", "Generate pack content", "Update modpack version", "Refresh modpack", "Init modpack"], "What do you want to do?")

    if not select:
        return

    if select != "Init modpack":
        minecraft_versions, launchers = get_mcv_launchers()
    else:
        minecraft_versions, launchers = [], []

    if (not minecraft_versions or not launchers) and not (select == "Init modpack" and not minecraft_versions and not launchers):
        main()
    else:
        match select:
            case "Add mod":
                mod_name = input("Enter the mod/resource pack/shader name: ")
                add_mod(minecraft_versions, launchers, mod_name)
            case "Remove mod":
                mod_name = input("Enter the mod/resource pack/shader name: ")
                remove_mod(minecraft_versions, launchers, mod_name)
            case "Update mods":
                mod_name = input("Enter the mod/resource pack/shader name (--all to update all): ")
                update_mods(minecraft_versions, launchers, mod_name)
            case "Accept version":
                version = input("Enter the version to accept: ")
                accept_version(minecraft_versions, launchers, version)
            case "Export modpack":
                export_modpack(minecraft_versions, launchers)
            case "List modpack":
                list_modpack(minecraft_versions, launchers)
                return
            case "Generate pack content":
                generate_pack_content(minecraft_versions, launchers)
            case "Update modpack version":
                update_modpack_version(minecraft_versions, launchers)
            case "Refresh modpack":
                refresh_modpack(minecraft_versions, launchers)
            case "Init modpack":
                init_modpack()
            case None:
                return
            case _:
                print("Invalid selection")
        main()

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nOperation aborted by user.")
        exit(-1)