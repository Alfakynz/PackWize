import subprocess

def run_cmd(cmd: list[str], cwd: str) -> None:
    """
    Run shell commands

    Arguments:
        cmd: list[str]. The commands to run, provided as a list of strings
        cwd: str. The working directory in which to run the command
    
    Returns:
        None if an error occurs
    """
    try:
        subprocess.run(cmd, check=True, cwd=cwd)
    except subprocess.CalledProcessError as e:
        print(f"Error running {' '.join(cmd)} in {cwd}: {e}")
        return # continue the script