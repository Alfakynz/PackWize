# PackWize

## 📜 Description

This CLI/TUI was created to manage modpacks more easily than just using [Packwiz](https://packwiz.infra.link/). <br />
Run Packwiz commands across multiple directories at the same time. <br />
Easily create modpacks by placing you config file into a `configurations` directory, it will be copied into your modpack directory when exporting it. <br />
Made for creators who maintain multiple versions of a modpack.

## ✨ Features

Run packwiz commands for all Minecraft versions and launchers you want at the same time. Supports most packwiz commands.

- Add mods to the modpack (url add supported)
- Remove mods from the modpack
- Update mods in the modpack
- Pin/unpin mods in the modpack
- Set an acceptable version
- Export the modpack contents to a ZIP or MRPACK file (files are moved to: `dist/{version}`)
- List mods in the modpack
- Export the modpack's content list to an MD file
- Update the modpack version (not the minecraft version)
- Refresh the pack.toml and index.toml files
- Migrate Minecraft/lodaer version to another
- Initialize a new modpack and create directories

All of these functions can be run for one Minecraft version in one launcher or for all Minecraft versions in all launchers.

## 🌲 Modpack tree example

Modpack <br />
┣ 1.21.1 <br />
┃ ┣ CurseForge <br />
┃ ┃ ┣ mods <br />
┃ ┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┃ ┣ sodium.pw.toml <br />
┃ ┃ ┃ ┗ text-placeholder-api.pw.toml <br />
┃ ┃ ┣ index.toml <br />
┃ ┃ ┗ pack.toml <br />
┃ ┣ Modrinth <br />
┃ ┃ ┣ mods <br />
┃ ┃ ┃ ┣ fabric-api.pw.toml <br />
┃ ┃ ┃ ┣ modmenu.pw.toml <br />
┃ ┃ ┃ ┣ placeholder-api.pw.toml <br />
┃ ┃ ┃ ┗ sodium.pw.toml <br />
┃ ┃ ┣ index.toml <br />
┃ ┃ ┗ pack.toml <br />
┃ ┣ configurations <br />
┃ ┃ ┣ config <br />
┃ ┃ ┃ ┗ sodium-options.json <br />
┃ ┃ ┣ .packwizignore <br />
┃ ┃ ┣ icon.png <br />
┃ ┃ ┗ options.txt <br />
┃ ┗ PACK_CONTENT.md <br />
┣ dist <br />
┃ ┣ 1.21.1 <br />
┃ ┃ ┣ Modpack-1.0.0.mrpack <br />
┗ ┗ ┗ Modpack-1.0.0.zip

## ⚙️ Requirements

- [Packwiz](https://packwiz.infra.link/)
- [Go 1.21.5 (or higher)](https://go.dev/)

## 📦 Installation

### 👤 For users

```sh
go install github.com/Alfakynz/PackWize/cmd/packwize@main
```

### 👨‍💻 For devs

```sh
git clone https://github.com/Alfakynz/PackWize.git
cd PackWize
```

Run the command without compiling:

```sh
go run ./cmd/packwize
```

To Compile:

```sh
make build
```

## ❌ Uninstallation

### 👤 For users

```sh
rm $(go env GOPATH)/bin/packwize
```

### 👨‍💻 For devs

You just need to remove the `PackWize` directory

## 🤝 Contributing

PackWize is a work in progress, and all suggestions are welcome.

If you'd like to contribute:

- Open an issue to suggest an idea, report a bug, or discuss a change
- Submit a pull request if you have an improvement to propose

## To Do

- Documentation (website)
- Create a TUI
