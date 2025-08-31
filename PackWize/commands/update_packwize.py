import os

def update_packwize():
    if os.name == 'nt':  # Windows
        print("Updating PackWize on Windows is not supported yet.")
    else:  # macOS and Linux
        os.system("curl -fsSL https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.sh | bash")