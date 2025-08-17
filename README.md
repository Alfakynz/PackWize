# PackWize

## ✨ Features

Run packwiz commands for all Minecraft versions and launchers you want at the same time. Supports most packwiz commands.

- Add mods to the modpack
- Remove mods from the modpack
- Update mods in the modpack
- Export the modpack contents to a ZIP or MRPACK file
- Export the modpack's content list to an MD file
- Refresh the pack.toml and index.toml files

All of these functions can be run for one Minecraft version in one launcher or for all Minecraft versions in all launchers

## 🌲 Modpack tree example

1.21.1 <br />
┣ CurseForge <br />
┃ ┣ mods <br />
┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┗ text-placeholder-api.pw.toml <br />
┃ ┣ index.toml <br />
┃ ┗ pack.toml <br />
┣ Modrinth <br />
┃ ┣ mods <br />
┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┗ placeholder-api.pw.toml <br />
┃ ┣ index.toml <br />
┃ ┗ pack.toml <br />
┣ configurations <br />
┗ PACK_CONTENT.md <br />

## 📦 Installation

### 👤 For users

### 👨‍💻 For devs

- `git clone https://github.com/Alfakynz/PackWize.git`
- `cd PackWize`
- `pip install -r requirements.txt` (Windows only)
- `pip install pyinstaller` (optional)
- `pip install .` or `pyinstaller --onefile PackWize/main.py packwize` (if you've installed PyInstaller)

## ⚙️ Requirements

- [Packwiz](https://packwiz.infra.link/)
- [Python 3.13 (or higher)](https://python.org) (for devs)

## 🤝 Contributing

PackWize is a work in progress, and all suggestions are welcome.

If you'd like to contribute:

- Open an issue to suggest an idea, report a bug, or discuss a change
- Submit a pull request if you have an improvement to propose
