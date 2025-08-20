# PackWize

This CLI/TUI was created to manage modpacks more easily than just using [Packwiz](https://packwiz.infra.link/). <br />
Run Packwiz commands across multiple directories at the same time. <br />
Easily create modpacks by placing you config file into a configuration directory, it will be copied into your modpack directory when exporting it. <br />
Made for creators who maintain multiple versions of a modpack.

## ✨ Features

Run packwiz commands for all Minecraft versions and launchers you want at the same time. Supports most packwiz commands.

- Add mods to the modpack
- Remove mods from the modpack
- Update mods in the modpack
- Export the modpack contents to a ZIP or MRPACK file
- Export the modpack's content list to an MD file
- Refresh the pack.toml and index.toml files
- Initialize a new modpack and create directories

All of these functions can be run for one Minecraft version in one launcher or for all Minecraft versions in all launchers.

## 🌲 Modpack tree example

1.21.1 <br />
┣ CurseForge <br />
┃ ┣ mods <br />
┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┣ sodium.pw.toml <br />
┃ ┃ ┗ text-placeholder-api.pw.toml <br />
┃ ┣ index.toml <br />
┃ ┗ pack.toml <br />
┣ Modrinth <br />
┃ ┣ mods <br />
┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┣ placeholder-api.pw.toml <br />
┃ ┃ ┗ sodium.pw.toml <br />
┃ ┣ index.toml <br />
┃ ┗ pack.toml <br />
┣ configurations <br />
┃ ┣ config <br />
┃ ┃ ┗ sodium-options.json <br />
┃ ┣ .packwizignore <br />
┃ ┣ icon.png <br />
┃ ┗ options.txt <br />
┗ PACK_CONTENT.md <br />

## ⚙️ Requirements

- [Packwiz](https://packwiz.infra.link/)
- [Python 3.13.5 (or higher)](https://python.org) (for devs)

## 📦 Installation

### 👤 For users

- Linux/macOS: `curl -fsSL https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.sh | bash`
- Windows: `irm https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.ps1 | iex`

### 👨‍💻 For devs

- `git clone https://github.com/Alfakynz/PackWize.git`
- `cd PackWize`
- `pip install -r requirements.txt` (Windows only)
- `pip install pyinstaller` (optional)
- `pip install .` or `pyinstaller --onefile PackWize/main.py packwize` (if you've installed PyInstaller)

## ❌ Uninstallation

### 👤 For users

- Linux/macos: `curl -fsSL https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.sh | bash -s uninstall`
- Windows: `iex "& { $(irm https://raw.githubusercontent.com/Alfakynz/PackWize/main/install.ps1) } -Uninstall"`

### 👨‍💻 For devs

- `pip uninstall packwize`

## 🤝 Contributing

PackWize is a work in progress, and all suggestions are welcome.

If you'd like to contribute:

- Open an issue to suggest an idea, report a bug, or discuss a change
- Submit a pull request if you have an improvement to propose

## To Do

- Support for other Packwiz commands
  - `packwiz settings acceptable-versions` ✅
  - `packwiz list` ✅
  - `packwiz migrate`
  - `packwiz pin` ✅
  - `packwiz unpin` ✅
  - `packwiz url add`
- Documentation
