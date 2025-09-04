from PackWize.utils.get_mc_version import get_mc_version
from PackWize.utils.get_launcher import get_launcher

def get_mcv_launchers() -> tuple[list[str], list[str]] | tuple[None, None]:
    """
    Get both launchers and Minecraft versions to work with
    
    Returns
        tuple[list[str], list[str]]: if both 'minecraft_versions' and 'launchers' are provided
        tuple[None, None]: if either 'minecraft_versions' or 'launchers' do not match
    """
    minecraft_versions = get_mc_version()
    if minecraft_versions is None:
        return None, None

    launchers = get_launcher()
    if launchers is None:
        return None, None
    
    return minecraft_versions, launchers