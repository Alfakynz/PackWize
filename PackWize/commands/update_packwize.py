import os
import subprocess
import json

def clean_version(v: str) -> str:
    """
    Clean version like v1.0.0, or PackWize 1.0.0

    Arguments:
        v: str. The version

    Returns:
        str. The cleaned version
    """
    v = v.strip()
    v = v.replace("Packwize", "").replace("packwize", "").strip()  # enlève le texte
    if v.startswith("v"):  # enlève le 'v' si présent
        v = v[1:]
    return v

def get_local_version() -> str:
    """
    Get your PackWize version. 

    Returns: str. Your PackWize version
    """
    try:
        result = subprocess.run(["packwize", "--version"], capture_output=True, text=True)
        if result.returncode == 0:
            return clean_version(result.stdout.strip())
    except FileNotFoundError:
        return "0.0.0"  # si packwize pas installé
    return "0.0.0"

def get_latest_version() -> str:
    """
    Get PackWize version from GitHub Releases. 

    Returns: str. The latest PackWize version
    """
    url = "https://api.github.com/repos/Alfakynz/PackWize/releases/latest"
    try:
        result = subprocess.run(
            ["curl", "-fsSL", url],
            capture_output=True,
            text=True,
            check=True
        )
        data = json.loads(result.stdout)
        return clean_version(data["tag_name"])
    except subprocess.CalledProcessError:
        print("❌ Failed to fetch latest version with curl.")
        return "0.0.0"


def version_tuple(v: str) -> tuple[int, ...]:
    """
    Transform a string version into a tuple. 

    Arguments:
        v: str. The version to convert
    
    Returns: tuple[int, ...]. The version converted
    """
    return tuple(int(x) for x in v.split(".") if x.isdigit())

def is_newer(latest: str, local: str) -> bool:
    """
    Check if a version is newer than an other. 

    Arguments:
        latest: str. The latest version
        local: str. The local version

    Returns:
        True. If the version is newer than the local
        False. If the version is not newer than the local
    """
    return version_tuple(latest) > version_tuple(local)

def update_packwize() -> None:
    """
    Update PackWize to the latest version if newer version exist.

    Return nothing
    """
    if os.name == 'nt':  # Windows
        print("Updating PackWize on Windows is not supported yet.")
        return
    
    local_version = get_local_version()
    latest_version = get_latest_version()

    if not is_newer(latest_version, local_version):
        print(f"PackWize is up to date ({local_version})")
        return

    print(f"Updating PackWize ({local_version} --> {latest_version}) ...")

    try:
        subprocess.run(
            ["bash", "-c", "curl -fsSL https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.sh | bash"],
            check=True
        )
        print("PackWize updated")
    except KeyboardInterrupt:
        print("\nUpdate cancelled by user.")
    except subprocess.CalledProcessError:
        print("Update failed.")
    """os.system("curl -fsSL https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.sh | bash")
    print("Packwize updated")"""