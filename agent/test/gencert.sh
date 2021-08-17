echo "enroll admin"
fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem --home /root/fabric_networks/organizations/org1.example.com/crypto

echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' >/root/fabric_networks/organizations/org1.example.com/crypto/msp/config.yaml

echo "register peer0"
fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem --home /root/fabric_networks/organizations/org1.example.com/crypto
echo "register org1admin"
fabric-ca-client register --caname ca-org1 --id.name Admin --id.secret Adminpw --id.type admin --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem --home /root/fabric_networks/organizations/org1.example.com/crypto
echo "register user1"
fabric-ca-client register --caname ca-org1 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem --home /root/fabric_networks/organizations/org1.example.com/crypto
echo "register orderer0"
fabric-ca-client register --caname ca-org1 --id.name orderer0 --id.secret orderer0pw --id.type orderer --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem --home /root/fabric_networks/organizations/org1.example.com/crypto
echo "generating peer0 msp"
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/msp --csr.hosts peer0.org1.example.com --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/msp/config.yaml /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/msp/config.yaml

echo "generating peer0 tls"
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls --enrollment.profile tls --csr.hosts peer0.org1.example.com --csr.hosts localhost --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/ca.crt
cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/signcerts/* /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/server.crt
cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/keystore/* /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/server.key

mkdir -p /root/fabric_networks/organizations/org1.example.com/crypto/msp/tlscacerts
cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/msp/tlscacerts/ca.crt

mkdir -p /root/fabric_networks/organizations/org1.example.com/crypto/tlsca
cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/tlsca/tlsca.org1.example.com-cert.pem

mkdir -p /root/fabric_networks/organizations/org1.example.com/crypto/ca
cp /root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/msp/cacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/ca/ca.org1.example.com-cert.pem

echo "generating orderer0 msp"
fabric-ca-client enroll -u https://orderer0:orderer0pw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/msp --csr.hosts orderer0.org1.example.com --csr.hosts localhost --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/msp/config.yaml /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/msp/config.yaml

echo "generating orderer0 tls"
fabric-ca-client enroll -u https://orderer0:orderer0pw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls --enrollment.profile tls --csr.hosts orderer0.org1.example.com --csr.hosts localhost --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/ca.crt
cp /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/signcerts/* /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/server.crt
cp /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/keystore/* /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/server.key

mkdir -p /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/msp/tlscacerts
cp /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/msp/tlscacerts/tlsca.org1.example.com-cert.pem

mkdir -p /root/fabric_networks/organizations/org1.example.com/crypto/msp/tlscacerts
cp /root/fabric_networks/organizations/org1.example.com/crypto/orderers/orderer0.org1.example.com/tls/tlscacerts/* /root/fabric_networks/organizations/org1.example.com/crypto/msp/tlscacerts/tlsca.org1.example.com-cert.pem

echo "generating org1admin msp"
fabric-ca-client enroll -u https://Admin:Adminpw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/msp/config.yaml /root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp/config.yaml

echo "generating user1 msp"
fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-org1 -M /root/fabric_networks/organizations/org1.example.com/crypto/users/user1@org1.example.com/msp --tls.certfiles /root/fabric_networks/fabric-ca/root-cert/tls-cert.pem

cp /root/fabric_networks/organizations/org1.example.com/crypto/msp/config.yaml /root/fabric_networks/organizations/org1.example.com/crypto/users/user1@org1.example.com/msp/config.yaml