#!/bin/bash
echo "Building fabric-manager project..."
./build.sh
echo "Packaging fabric-manager project..."
mkdir -p fabric-manager/server
mkdir -p fabric-manager/agent
mkdir -p fabric-manager/explorer
mkdir -p fabric-manager/ca/root-cert

# move build artifacts to temporary packaging folder
# server
cp ./server/fabric-server ./fabric-manager/server
cp ./server/config.json ./fabric-manager/server
# agent
cp ./agent/fabric-agent ./fabric-manager/agent
# explorer
cp ./explorer/fabric-explorer ./fabric-manager/explorer
# etcd
cp -r ./etcd ./fabric-manager
# ca
cp -r ./fabric-ca/docker-compose.yaml ./fabric-manager/ca
cp -r ./fabric-ca/fabric-ca-server-config.yaml ./fabric-manager/ca/root-cert
# deployment scripts
cp -r ./deploy ./fabric-manager

# package
tar -czvf fabric-manager.tar.gz fabric-manager
# clean up temp packaging folder
rm -rf fabric-manager
echo "package succeeded"