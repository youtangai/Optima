#!/bin/bash
curl -X POST -d '{"host_name":"'$HOSTNAME'"}' -H "Content-Type:application/json" $CONDUCTOR_IP:62070/leave
