#!/usr/bin/env bash

FILE="micropython.uf2"

if [ ! -f $FILE ]; then
    echo "MicroPython not found"
    echo "Downloading MicroPython..."
    wget -c --show-progress --output-document=$FILE \
        https://micropython.org/resources/firmware/RPI_PICO_W-20240222-v1.22.2.uf2
    echo "Download Complete"
fi

echo "Installing MicroPython..."
cp micropython.uf2 /media/$USER/RPI-RP2
echo "Installation Complete"
