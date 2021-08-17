pushd /root/fabric_networks/fabric-ca
docker-compose down
popd
pushd /root/fabric_networks/organizations/org1.example.com/docker
docker-compose -f orderer0.org1.example.com.yaml -f peer0.org1.example.com.yaml down --remove-orphan
popd
rm -rf /root/fabric_networks/consortiums/* 
rm -rf /root/fabric_networks/organizations/*
rm -rf /root/fabric_networks/fabric-ca/client/*
mv /root/fabric_networks/fabric-ca/root-cert/*.yaml /root/fabric_networks/fabric-ca
rm -rf /root/fabric_networks/fabric-ca/root-cert/*
mv /root/fabric_networks/fabric-ca/fabric-ca-server-config.yaml /root/fabric_networks/fabric-ca/root-cert