#!/bin/bash

# Script to install the 'ks' executable either system-wide or for the current user only.
# Ensures the target directory exists and handles user confirmation.

echo "Would you like to install 'ks' for all users or for this user only?"
echo "1) All users (requires sudo)"
echo "2) This user only"
read -p "Enter your choice (1 or 2): " choice

case "$choice" in
    1)
        if [ ! -f "ks" ]; then
            echo "Error: 'ks' not found in the current directory."
            exit 1
        fi
        sudo mkdir -p /usr/local/bin
        sudo cp ks /usr/local/bin/
        echo "'ks' has been installed system-wide in /usr/local/bin/."
        ;;
    2)
        if [ ! -f "ks" ]; then
            echo "Error: 'ks' not found in the current directory."
            exit 1
        fi
        mkdir -p ~/.local/bin
        cp ks ~/.local/bin/
        echo "'ks' has been installed for the current user in ~/.local/bin/."
        echo "Ensure that ~/.local/bin is in your PATH if it is not already."
        ;;
    *)
        echo "Invalid choice. Installation aborted."
        exit 1
        ;;
esac

