#!/bin/bash
# Simple install script for mcswain CLI

VERSION="v1.0.0"
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH="amd64"

if [ "$OS" == "linux" ]; then
    BINARY="mcswain-linux-${ARCH}"
elif [ "$OS" == "darwin" ]; then
    BINARY="mcswain-darwin-${ARCH}"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

URL="https://github.com/johnmcswain/cli/releases/download/${VERSION}/${BINARY}"
echo "Downloading $URL..."
curl -L -o mcswain "$URL"
chmod +x mcswain
sudo mv mcswain /usr/local/bin/mcswain
echo "mcswain installed successfully!"
