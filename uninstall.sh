#!/bin/bash
# uninstall.sh: Remove the installed mcswain binary

TARGET="/usr/local/bin/mcswain"

if [ -f "$TARGET" ]; then
  echo "Removing $TARGET..."
  sudo rm "$TARGET"
  echo "mcswain has been uninstalled."
else
  echo "mcswain is not installed at $TARGET."
fi
