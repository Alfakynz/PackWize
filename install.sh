#!/usr/bin/env bash
set -euo pipefail

REPO="Alfakynz/PackWize"
INSTALL_DIR="$HOME/.local/bin"
BINARY="$INSTALL_DIR/packwize"

# Color codes for output
RED="\033[0;31m"
RESET="\033[0m"

if [ "${1:-}" = "uninstall" ]; then
  echo "Uninstalling Packwize..."
  if [ -f "$BINARY" ]; then
    rm -f "$BINARY"
    echo "Removed $BINARY"
  else
    echo "Packwize is not installed in $INSTALL_DIR"
  fi
  exit 0
fi

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux*)   OS="linux" ;;
  darwin*)  OS="macos" ;;
  *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

# Detect latest version from GitHub API
VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')
if [ -z "$VERSION" ]; then
  echo "Could not fetch latest version from GitHub"
  exit 1
fi

FILENAME="packwize-${OS}.zip"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILENAME"

echo "Downloading $URL..."
curl -L "$URL" -o "/tmp/$FILENAME"

echo "Extracting..."
rm -rf /tmp/packwize
unzip -o "/tmp/$FILENAME" -d /tmp/packwize >/dev/null

echo "Installing to $INSTALL_DIR"
mkdir -p "$INSTALL_DIR"
mv -f "/tmp/packwize/packwize" "$BINARY"
chmod +x "$BINARY"

if ! command -v packwiz >/dev/null 2>&1; then
  echo -e "${RED}Error: packwiz is not installed or not in PATH.${RESET}"
  echo "You can install it here: https://packwiz.infra.link/installation/"
fi

echo "Installation complete!"
echo "Make sure $INSTALL_DIR is in your PATH"
echo "Run: packwize --help"
