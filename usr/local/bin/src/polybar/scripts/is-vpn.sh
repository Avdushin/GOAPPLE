#!/bin/bash

SERVICE_NAME="windscribe"

if [ "$(curl -4 -sf 2ip.io)" != "your_ip" ]
then
    echo "%{F#00ff00} $(curl -4 -sf 2ip.io)"
else
    echo "%{F#ed5858} VPN OFF"
fi

"$@"