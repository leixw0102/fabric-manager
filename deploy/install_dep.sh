#!/bin/bash -e
# execution dir: /root/fabric-manager/deploy
function CheckPrerequisits() {
    if [ ! git --version 2> /dev/null ]; then
        echo "Please install git."
        exit 1
    fi
    if [ ! curl --version 2> /dev/null ];then
        echo "Please install curl."
        exit 1
    fi
    if [ ! yum --version 2> /dev/null ];then
        echo "Please install yum."
        exit 1
    fi
    echo "All prerequisits satisfied."
}

function InstallDocker(){
    sudo yum install -y yum-utils
    sudo yum-config-manager \
        --add-repo \
        https://download.docker.com/linux/centos/docker-ce.repo
    echo y | sudo yum install docker-ce docker-ce-cli containerd.io  # bypass yes/no questions
    if ! docker --version;then
        echo "Fail to install docker. Please refer to https://docs.docker.com/engine/install/centos/ to install it mannually."
        exit 1
    fi
    echo "Docker installed successfully."
}

function InstallDockerCompose(){
    sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
}

# configtxgen, fabric-ca-client.
function InstallFabricBin() {
    cp ./bin/fabric-ca-client /usr/local/bin/
    cp ./bin/configtxgen /usr/local/bin/
}

CheckPrerequisits
InstallDocker
InstallDockerCompose
InstallFabricBin
echo "All dependencies are installed."