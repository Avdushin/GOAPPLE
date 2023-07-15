#!/bin/bash

SERVICE_NAME="openvpn"

if [ "$(systemctl is-active "$SERVICE_NAME")" != "active" ]
then
    echo '%{F#ff0000} '$SERVICE_NAME
else
    echo '%{F#00ff00} '$SERVICE_NAME
fi


enable () {
    sudo systemctl start "$SERVICE_NAME"
}

disable () {
    sudo systemctl stop "$SERVICE_NAME"
}

toggle () {
    if [ "$(systemctl is-active $SERVICE_NAME)" != "active" ]
    then
        sudo systemctl start $SERVICE_NAME
    else
        sudo systemctl stop $SERVICE_NAME
    fi
}

"$@"