import argparse
from PackWize.utils.menu import menu
from PackWize.utils.get_mcv_launchers import get_mcv_launchers
from PackWize.utils.convert_arguments import convert_arguments
from PackWize.utils.better_input import better_input
from PackWize.commands.add_mod import add_mod
from PackWize.commands.remove_mod import remove_mod
from PackWize.commands.update_mods import update_mods
from PackWize.commands.pin_mod import pin_mod
from PackWize.commands.unpin_mod import unpin_mod
from PackWize.commands.url_add import url_add
from PackWize.commands.accept_version import accept_version
from PackWize.commands.export_modpack import export_modpack
from PackWize.commands.list_modpack import list_modpack
from PackWize.commands.generate_pack_content import generate_pack_content
from PackWize.commands.update_modpack_version import update_modpack_version
from PackWize.commands.refresh_modpack import refresh_modpack
from PackWize.commands.migrate import migrate
from PackWize.commands.init_modpack import init_modpack
from PackWize.commands.update_packwize import update_packwize

VERSION="1.5.6"

# CLI function
def main() -> None:
    """
    CLI function, use `python main.py [argument]`
    
    Returns nothing
    """
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
    parser_update.add_argument("mod", nargs="?", help="Mod/resource pack/shaderpack you want to update. Use '--all' to update all)")
    parser_update.add_argument("--all", action="store_true", help="Update all mods/resource packs/shaderpacks in the modpack")

    # Pin mod
    parser_pin = subparsers.add_parser("pin", aliases=["hold"], parents=[common_parser], help="Pin a mod in the modpack to prevent it from being updated automatically")
    parser_pin.add_argument("mod", help="Mod/resource pack/shaderpack you want to pin.")

    # Unpin mod
    parser_unpin = subparsers.add_parser("unpin", aliases=["unhold"], parents=[common_parser], help="Unpin a mod in the modpack to allow it to be updated automatically")
    parser_unpin.add_argument("mod", help="Mod/resource pack/shaderpack you want to unpin.")

    # URL add
    parser_url_add = subparsers.add_parser("url-add", aliases=["url", "ua"], parents=[common_parser], help="Add a custom mod/resource pack/shaderpack in the modpack from a URL")
    parser_url_add.add_argument("mod", help="Name of the mod/resource pack/shaderpack.")
    parser_url_add.add_argument("url", help="URL of the mod/resource pack/shaderpack.")

    # Accept version
    parser_accept_version = subparsers.add_parser("accept-version", aliases=["av"], parents=[common_parser], help="Accept a specific version for the modpack")
    parser_accept_version.add_argument("version", help="Version to accept")

    # Export modpack
    parser_export = subparsers.add_parser("export", aliases=["ex", "build"], parents=[common_parser], help="Export the modpack content to a ZIP or MRPACK file. Find the file in the {Minecraft version}/{launcher} directory")

    # List modpack
    parser_list = subparsers.add_parser("list", aliases=["ls"], parents=[common_parser], help="List mods, resource packs and shaders in the modpack")

    # Generate modpack content list
    parser_generate = subparsers.add_parser("generate", aliases=["gen"], parents=[common_parser], help="Export the modpack's content list to an MD file. Find the file in the {Minecraft version} directory")

    # Update modpack version
    parser_update_version = subparsers.add_parser("update-version", aliases=["uv", "set-version", "sv", "change-version", "cv"], parents=[common_parser], help="Update the modpack version (not the Minecraft version)")

    # Refresh modpack
    parser_refresh = subparsers.add_parser("refresh", aliases=["rf"], parents=[common_parser], help="Refresh the pack.toml and index.toml files")

    # Migrate modpack
    parser_migrate = subparsers.add_parser("migrate", aliases=["mg"], parents=[common_parser], help="Migrate the modpack to a new Minecraft/loader version")
    parser_migrate.add_argument("target", help="Choose between Minecraft or loader version to migrate")
    parser_migrate.add_argument("version", help="Enter the target version to migrate to")

    # Init modpack
    parser_init = subparsers.add_parser("init", help="Initialize a new modpack and create directories")

    # Update PackWize
    parser_update_packwize = subparsers.add_parser("update-packwize", aliases=["up"], help="Update PackWize to the latest version")

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
            if args.all:
                update_mods(minecraft_versions, launchers, "--all")
            elif args.mod:
                update_mods(minecraft_versions, launchers, args.mod)
            else:
                print("Please specify a mod to update or use '--all' to update all mods.")

        case "pin" | "hold":
            pin_mod(minecraft_versions, launchers, args.mod)

        case "unpin" | "unhold":
            unpin_mod(minecraft_versions, launchers, args.mod)

        case "url-add" | "url" | "ua":
            url_add(minecraft_versions, launchers, args.mod, args.url)

        case "accept-version" | "av":
            accept_version(minecraft_versions, launchers, args.version)

        case "export" | "ex" | "build":
            export_modpack(minecraft_versions, launchers)

        case "list" | "ls":
            list_modpack(minecraft_versions, launchers)

        case "generate" | "gen":
            generate_pack_content(minecraft_versions, launchers)

        case "update-version" | "uv" | "set-version" | "sv" | "change-version" | "cv":
            update_modpack_version(minecraft_versions, launchers)

        case "refresh" | "rf":
            refresh_modpack(minecraft_versions, launchers)

        case "migrate" | "mg":
            migrate(minecraft_versions, launchers, args.target, args.version)

        case "init":
            init_modpack()

        case "update-packwize" | "up":
            update_packwize()

        case "tui" | _:
            try:
                selection()
            except KeyboardInterrupt:
                print("\nOperation aborted by user.")
                exit(-1)

# TUI function
def selection() -> None:
    """
    TUI function, use `python main.py` and select an option from the menu

    Returns nothing
    """
    select = menu(["Add mod", "Remove mod", "Update mods", "Pin mod", "Unpin mod", "Url add", "Accept version", "Export modpack", "List modpack", "Generate pack content", "Update modpack version", "Refresh modpack", "Migrate", "Init modpack", "Update PackWize"], "What do you want to do?")

    if not select:
        return

    if select != "Init modpack" and select != "Update PackWize":
        minecraft_versions, launchers = get_mcv_launchers()
    else:
        minecraft_versions, launchers = [], []

    if (not minecraft_versions or not launchers) and not select == "Init modpack" and not select == "Update PackWize":
        main()
    else:
        match select:
            case "Add mod":
                mod_name = better_input("Enter the mod/resource pack/shader name: ")
                add_mod(minecraft_versions, launchers, mod_name)
            case "Remove mod":
                mod_name = better_input("Enter the mod/resource pack/shader name: ")
                remove_mod(minecraft_versions, launchers, mod_name)
            case "Update mods":
                mod_name = better_input("Enter the mod/resource pack/shader name (--all to update all): ", default_value="--all")
                update_mods(minecraft_versions, launchers, mod_name)
            case "Pin mod":
                mod_name = better_input("Enter the mod/resource pack/shader name: ")
                pin_mod(minecraft_versions, launchers, mod_name)
            case "Unpin mod":
                mod_name = better_input("Enter the mod/resource pack/shader name: ")
                unpin_mod(minecraft_versions, launchers, mod_name)
            case "Url add":
                mod_name = better_input("Enter the mod/resource pack/shader name: ")
                url = better_input("Enter the URL to add: ")
                url_add(minecraft_versions, launchers, mod_name, url)
            case "Accept version":
                version = better_input("Enter the version to accept: ")
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
            case "Migrate":
                target = better_input("Enter the target to migrate (minecraft/loader): ")
                version = better_input("Enter the target version to migrate to: ")
                migrate(minecraft_versions, launchers, target, version)
            case "Init modpack":
                init_modpack()
            case "Update PackWize":
                update_packwize()
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