#!/bin/bash
ip=$2

echo "host ip is:$ip"
# create identities for org
echo "creating orgs"
curl -X POST -H "Content-Type:application/json" -d @create_crypto.json http://$ip:8081/createCrypto
sleep 5
# create consortium
echo "creating consortium"
curl -X POST -H "Content-Type:application/json" -d @create_consortium.json http://$ip:8081/createConsortium
sleep 5
# start network
echo "starting network"
curl -X POST -H "Content-Type:application/json" -d @start_network.json http://$ip:8081/startConsortium
sleep 30
# create channel
echo "creating channel"
curl -X POST -H "Content-Type:application/json" -d @create_channel.json http://$ip:8081/createChannel