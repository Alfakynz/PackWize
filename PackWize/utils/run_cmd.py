import subprocess

def run_cmd(cmd, cwd):
    try:
        subprocess.run(cmd, check=True, cwd=cwd)
    except subprocess.CalledProcessError as e:
        print(f"Error running {' '.join(cmd)} in {cwd}: {e}")
        return # continue the script