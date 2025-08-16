from utils.get_mc_version import get_mc_version
from utils.get_launcher import get_launcher

def get_mcv_launchers():
    minecraft_versions = get_mc_version()
    if minecraft_versions is None:
        return None, None

    launchers = get_launcher()
    if launchers is None:
        return None, None
    
    return minecraft_versions, launchers