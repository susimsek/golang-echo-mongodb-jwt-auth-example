#!/bin/bash

# DO NOT Execute this script with sudo
if [ $SUDO_USER ]; then
    echo "Please DO NOT execute with sudo !!!    ./install-prereqs.sh"
    echo "Aborting!!!"
    exit 0
fi
sudo ./docker.sh
sudo ./compose.sh
sudo ./golang.sh

echo "====== Please Logout & Logback in ======"