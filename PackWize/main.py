import argparse
from PackWize.utils.menu import menu
from PackWize.utils.get_mcv_launchers import get_mcv_launchers
from PackWize.utils.convert_arguments import convert_arguments
from PackWize.commands.add_mod import add_mod
from PackWize.commands.remove_mod import remove_mod
from PackWize.commands.update_mods import update_mods
from PackWize.commands.export_modpack import export_modpack
from PackWize.commands.generate_pack_content import generate_pack_content
from PackWize.commands.update_modpack_version import update_modpack_version
from PackWize.commands.refresh_modpack import refresh_modpack

# CLI function
def main():
    parser = argparse.ArgumentParser(prog="packwize", description="A CLI/TUI to manage modpack easier than just use Packwiz. Based on Packwiz")
    parser.add_argument("-v", "--version", action="version", version="Packwize 0.0.1", help="Show the version of PackWize")

    subparsers = parser.add_subparsers(dest="command", required=False, help="Available commands")

    # TUI menu
    parser_tui = subparsers.add_parser("tui", help="Show TUI menu")

    # Add mods
    parser_add = subparsers.add_parser("add", aliases=["install", "i"], help="Add a mod to the modpack")
    parser_add.add_argument("mod", help="Mod/resource pack/shaderpack you want to add")

    # Remove mods
    parser_remove = subparsers.add_parser("remove", aliases=["rm", "uninstall"], help="Remove a mod from the modpack")
    parser_remove.add_argument("mod", help="Mod/resource pack/shaderpack you want to remove")

    # Update mods
    parser_update = subparsers.add_parser("update", aliases=["upgrade"], help="Update mod in the modpack")
    parser_update.add_argument("mod", help="Mod/resource pack/shaderpack you want to update. Use '--all' to update all)")

    # Export modpack
    parser_export = subparsers.add_parser("export", aliases=["build"], help="Export the modpack content to a ZIP or MRPACK file. Find the file in the {Minecraft version}/{launcher} directory")

    # Generate modpack content list
    parser_generate = subparsers.add_parser("generate", aliases=["gen"], help="Export the modpack's content list to an MD file. Find the file in the {Minecraft version} directory")
    # Update modpack version
    parser_update_version = subparsers.add_parser("update-version", aliases=["uv", "set-version", "change-version"], help="Update the modpack version (not the Minecraft version)")
    # Refresh modpack
    parser_refresh = subparsers.add_parser("refresh", aliases=["rf"], help="Refresh the pack.toml and index.toml files")

    commands = [parser_add, parser_remove, parser_update, parser_export, parser_generate, parser_update_version, parser_refresh]
    
    for command in commands:
        command.add_argument("minecraft_version", help="Minecraft version to work in. Use 'all' to select all Minecraft versions")
        command.add_argument("launcher", help="Launcher to work in. Use 'all' to select all launchers")

    args = parser.parse_args()

    if hasattr(args, "minecraft_version") and hasattr(args, "launcher"):
        minecraft_version = convert_arguments("minecraft_version", args.minecraft_version)
        launcher = convert_arguments("launcher", args.launcher)

    match args.command:
        case "add" | "install" | "i":
            add_mod(minecraft_version, launcher, args.mod)
        
        case "remove" | "rm" | "uninstall":
            remove_mod(minecraft_version, launcher, args.mod)

        case "update" | "upgrade":
            update_mods(minecraft_version, launcher, args.mod)

        case "export" | "build":
            export_modpack(minecraft_version, launcher)

        case "generate" | "gen":
            generate_pack_content(minecraft_version, launcher)

        case "update-version" | "uv" | "set-version" | "change-version":
            update_modpack_version(minecraft_version, launcher)

        case "refresh" | "rf":
            refresh_modpack(minecraft_version, launcher)

        case "tui" | _:
            selection()

# TUI function
def selection():
    select = menu(["Add mod", "Remove mod", "Update mods", "Export modpack", "Generate pack content", "Update modpack version", "Refresh modpack"], "What do you want to do?")

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
            case "Generate pack content":
                generate_pack_content(minecraft_versions, launchers)
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
    try:
        main()
    except KeyboardInterrupt:
        print("\nOperation aborted by user.")
        exit(-1)