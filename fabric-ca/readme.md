# Fabric CA操作手册

## Fabric CA Server
- 安装docker-compose
- 下载docker-image：fabric-ca-1.4.9
- 编写docker-compose.yaml, 参照该repo样例。
- fabric-ca-server-config.yaml的version不能超过fabirc-ca image的version
- docker-compose.yaml中voluems mount到/etc/hyperledger/fabric-ca-server的目录在fabric-ca-server启动后会自动生成fabric-ca-server-config.yaml
- 如果用docker 启动fabric-ca-server，要在docker-compose的start 命令上加上参数csr.hosts, 值为机器的ip, 例如：
```bash
sh -c 'fabric-ca-server start -b admin:adminpw -d --csr.hosts "172.38.50.212"'
```

## Fabric CA Client
- 下载 [fabric-ca binary](https://github.com/hyperledger/fabric-ca/releases/download/v1.4.9/hyperledger-fabric-ca-linux-amd64-1.4.9.tar.gz)
- 解压下载的压缩包至自己指定的路径
- 将自己指定的路径添加到`PATH`，方便后续调用fabric-ca-client

## 生成证书流程
```bash
fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles tls-cert.pem

fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles tls-cert.pem

fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M msp --csr.hosts peer0.org1.example.com --tls.certfiles tls-cert.pem

fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M msp --enrollment.profile tls --csr.hosts peer0.org1.example.com --csr.hosts localhost --tls.certfiles tls-cert.pem
```
