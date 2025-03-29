#!/bin/bash
set -e

# Use the provided version (first argument) or default to "0.2-beta"
VERSION=${1:-"0.2-beta"}

echo "Building macOS binary..."
# Build the macOS binary (for Intel; adjust GOARCH if needed for arm64)
GOOS=darwin GOARCH=amd64 go build -o build/mcswain-darwin-amd64

chmod +x build/mcswain-darwin-amd64

echo "Creating package structure..."
# Create the directory structure as it would be installed (in this case /usr/local/bin)
mkdir -p package-root/usr/local/bin
cp build/mcswain-darwin-amd64 package-root/usr/local/bin/mcswain
chmod +x package-root/usr/local/bin/mcswain

echo "Building macOS installer pkg..."
# Create the installer package.
pkgbuild --root package-root --identifier com.johnmcswain.mcswain --version "$VERSION" build/mcswain.pkg

echo "Package built successfully: build/mcswain.pkg"