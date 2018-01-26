#!/bin/bash
DIR="/var/optima"
if [ ! -d ${DIR} ]; then
    export CONDUCTOR_IP="192.168.64.12"
    echo $CONDUCTOR_IP
    sudo mkdir -p ${DIR}
    sudo cd /var
    sudo cd /optima
    sudo ssh-keygen -t rsa -f optima_key
    sudo scp ./optima_key.pub root@$CONDUCTOR_IP:/var/optima/
fi