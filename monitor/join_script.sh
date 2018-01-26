#!/bin/bash
DIR="/var/optima"
if [ ! -d ${DIR} ]; then
    echo $CONDUCTOR_IP
    mkdir -p ${DIR}
    cd ${DIR}
    ssh-keygen -t rsa -f optima_key
    curl -X POST -d '{"host_name":"'$HOSTNAME'"}' -H "Content-Type:application/json" $CONDUCTOR_IP:62070/create_dir
    scp ./optima_key.pub root@$CONDUCTOR_IP:/var/optima/$HOSTNAME/
    curl -X POST -d '{"host_name":"'$HOSTNAME'"}' -H "Content-Type:application/json" $CONDUCTOR_IP:62070/join
fi