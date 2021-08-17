# Fabric Agent 安裝指南

## CA相关Setup
- 将ca证书copy到agent所在机器的/root/fabric-ca/root-certs目录下，命名为ca-cert.pem
- 下载 [fabric-ca binary](https://github.com/hyperledger/fabric-ca/releases/download/v1.5.0/hyperledger-fabric-ca-linux-amd64-1.5.0.tar.gz)，解压缩并将bin文件夹copy到agent机器的/root/
- 将/root/bin加入系统环境变量