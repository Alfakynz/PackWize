from pathlib import Path

def convert_arguments(argument: str, value: str) -> list[str] | None:
    """
    Convert arguments like `minecraft_versions` or `launchers`

    Arguments: 
        argument: str. The argument type to check ('minecraft_versions' or 'launchers' are supported)
        value: str. The value to convert
    
    Returns:
        list[str]
        None: if nothing matchs
    """
    match argument:
        case "minecraft_versions":
            match value.lower():
                case "all":
                    directories = sorted([
                        d.name
                        for d in Path().iterdir()
                        if d.is_dir()
                        and not d.name.startswith('.')
                        and any(c.isdigit() for c in d.name)
                    ])

                    return directories
                case _:
                    return [value]
        
        case "launchers":
            match value.lower():
                case "all":
                    return ["CurseForge", "Modrinth"]
                case "modrinth" | "mr":
                    return ["Modrinth"]
                case "curseforge" | "cf":
                    return ["CurseForge"]

        case _:
            return
