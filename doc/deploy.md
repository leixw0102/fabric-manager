# Fabric Manager 部署文档

## 服务架构
fabric manager区块链服务可拆分为5个子服务
- server： 前端服务器，用于接收用户的各类请求
- agent：安装在区块链网络的各个虚拟机上，接受server传来的各类请求并执行相应区块链命令
- ca：接收用户的证书生成请求并生成组织各成员的证书、私钥等
- etcd： 消息转发中间件，负责server与agent间通信
- explorer：可视化展示整个区块链网络状态

## 最小化部署
5个子服务均部署在一台机器上

### 机器配置：
- CentOS操作系统
- git
- yum
- vim
- curl
- ifconfig
- 联网

### 下载依赖
下载fabric-manager.tar.gz至部署机/root路径下并解压，下载地址：xxxxx，
```bash
cd /root
tar -xzvf fabric-manager.tar.gz
```

运行`install_dep.sh`下载环境依赖（docker,docker-compose, fabric bin）
```bash
cd /root/fabric-manager/deploy
chmod +x install_dep.sh
./install_dep.sh
```
看到`All dependencies are installed.`说明依赖安装成功。

### 部署CA与etcd
直接运行部署脚本`deploy.sh`,执行脚本需较长时间，请耐心等待。
```bash
cd /root/fabric-manager/deploy
./deploy.sh
```
如看到`fabric-manager successfully deployed!`说明部署成功。

### 启动server
开启一个新的bash窗口（server窗口），运行以下代码：
```bash
pushd /root/fabric-manager/server
chmod +x fabric-server
./fabric-server
popd
```

### 启动agent
开启一个新的bash窗口（agent窗口），运行以下代码：
```bash
pushd /root/fabric-manager/agent
chmod +x fabric-agent
./fabric-agent --etcd_addr 0.0.0.0:2379
popd
```

### 创建区块链网络
开启一个新的bash窗口，获取机器ip
```bash
ip route get 1 | awk '{print $NF;exit}'
```

修改/etc/hosts, 假设上一步获取的ip为`192.168.133.130`,添加以下内容：
```bash
192.168.133.130 peer0.org1.example.com
192.168.133.130 orderer0.org1.example.com
192.168.133.130 ca.org1.example.com
```

运行test_network.sh, 创建区块链网络
```bash
pushd /root/fabric-manager/deploy
chmod +x test_network.sh
./test_network.sh
popd
```

运行过后运行以下命令：
```bash
docker exec cli peer channel create -o orderer0.org1.example.com:7050 -c mychannel --ordererTLSHostnameOverride orderer0.org1.example.com -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/mychannel.tx --outputBlock /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/mychannel.block --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/orderers/orderer0.org1.example.com/msp/tlscacerts/tls-localhost-7054-ca-org1.pem
```

若看到以下输出说明网络创建成功：
```bash
2021-07-16 08:26:13.767 UTC [channelCmd] InitCmdFactory -> INFO 0a4 Endorser and orderer connections initialized
2021-07-16 08:26:13.968 UTC [msp.identity] Sign -> DEBU 0a5 Sign: plaintext: 0AE3070A1508051A0608A585C5870622...86B29C1704FB12080A021A0012021A00
2021-07-16 08:26:13.968 UTC [msp.identity] Sign -> DEBU 0a6 Sign: digest: CCD95BA248D4520CF30C922B42FF472B4F25B006C9D3A5152254335E883D862C
2021-07-16 08:26:13.971 UTC [cli.common] readBlock -> INFO 0a7 Received block: 0
```

## 多级部署
### 机器配置：
同最小化部署
### 部署etcd

### 部署CA