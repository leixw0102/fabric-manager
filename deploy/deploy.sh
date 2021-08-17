#!/bin/bash
packageDir=$PWD/..
function DeployEtcd() {
    sudo rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
    docker rmi quay.io/etcd-development/etcd:v3.2.32 || true && \
    docker run -d \
    -p 2379:2379 \
    -p 2380:2380 \
    --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
    --name etcd-gcr-v3.2.32 \
    quay.io/coreos/etcd:v3.2.32 \
    /usr/local/bin/etcd \
    --name s1 \
    --data-dir /etcd-data \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --initial-cluster s1=http://0.0.0.0:2380 \
    --initial-cluster-token tkn \
    --initial-cluster-state new
    if [ ! $? -eq 0 ]
    then
        echo "Fail to deploy etcd."
    else
        echo "Deploy etcd successfully."
    fi
}

function DeployCA() {
    docker-compose -f $packageDir/ca/docker-compose.yaml up -d
    if [ ! $? -eq 0 ]
    then
        echo "Fail to deploy CA."
    else
        echo "Deploy CA successfully."
        cp /root/fabric-manager/ca/root-cert/ca-cert.pem /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem
    fi
}

function DeployAgent() {
    pushd $packageDir/agent/
    chmod +x fabric-agent
    ./fabric-agent --etcd_addr 0.0.0.0:2379 &
    if [ ! $? -eq 0 ]
    then
        echo "Fail to deploy agent."
    else
        echo "Deploy agent successfully."
    fi
    popd
}

function DeployServer() {
    pushd $packageDir/server
    chmod +x fabric-server
    ./fabric-server &
    if [ ! $? -eq 0 ]
    then
        echo "Fail to deploy server."
    else
        echo "Deploy server successfully."
    fi
    popd
}

function DeployExplorer() {
    # copy certs to explorer folder
    mkdir -p /root/explorer/organizations/org1/admin
    mkdir -p /root/explorer/organizations/org1/peer
    cp /root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp/keystore/* /root/explorer/organizations/org1/admin/private_key
    cp /root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp/signcerts/cert.pem /root/explorer/organizations/org1/admin/cert.pem
    cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/ca.crt /root/explorer/organizations/org1/peer/tlsca.crt
}
    
echo "deploying etcd"
DeployEtcd
echo "deploying ca"
DeployCA
# echo "deploying server"
# DeployServer
# echo "deploying agent"
# DeployAgent