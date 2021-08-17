#!/bin/bash

# create identities for org
curl -X POST -H "Content-Type:application/json" -d @create_org.json http://192.168.133.130:8081/createOrg
# create consortium
curl -X POST -H "Content-Type:application/json" -d @create_consortium.json http://192.168.133.130:8081/createConsortium
# start network
curl -X POST -H "Content-Type:application/json" -d @create_consortium.json http://192.168.133.130:8081/startConsortium
# create channel
curl -X POST -H "Content-Type:application/json" -d @create_channel.json http://192.168.133.130:8081/createChannel
