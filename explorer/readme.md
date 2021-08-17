# Fabric Explorer

本文档用于部署区块链浏览器。

## 前期准备
- 一台联网的linux虚拟机，用于部署浏览器服务
- 工具：docker, docker-compose, git, go
- 下载源代码：http://10.20.5.5:10080/Data_Bank/fabric-manager/explorer
- 一个已经启动的区块链网络

## 编译可执行文件
在explorer目录下，执行build.sh
```bash
cd explorer
./build.sh -ip 172.38.40.240 -d /root/explorer -u root
```

参数说明：
```bash
-ip: 目标机器ip，将build好的可执行文件传至该机器
-d: 目标机器的上传路径，可执行文件上传至该路径下，请确保该路径存在
-u: 目标机器的用户名
```

根据命令提示输入机器的登录密码

## 将浏览器所需证书密钥下载至部署机
1. 需要下载区块链网络任一组织的三个文件：
    - admin私钥 (如/root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp/keystore/af225aaaa5967e4753b31e96ffc60536c99bd30697c6f3be07c5d889f4ff9f45_sk)
    - admin证书 (如/root/fabric_networks/organizations/org1.example.com/crypto/users/Admin@org1.example.com/msp/signcerts/cert.pem)
    - 任一peer的tlscacert (如/root/fabric_networks/organizations/org1.example.com/crypto/peers/peer0.org1.example.com/tls/ca.crt)
2. 将三个文件下载至explorer所在路径,并重命名
    - admin私钥重命名为private_key, 路径为$explorerDir/organizations/$orgName/admin/private_key
    - admin证书重命名为cert.pem, 路径为$explorerDir/organizations/$orgName/admin/cert.pem
    - peer的tlscacert重命名为tlsca.crt,路径为$explorerDir/organizations/$orgName/peer/tlsca.crt

    其中，$explorerDir为之前将可执行文件`fabric-explorer`上传的目录，如/root/explorer。$orgName为组织名

    样例路径如下
    ```bash    
        /root/explorer/
                ├── fabric-explorer
                ├── organizations
                │   └── Org1
                │       ├── admin
                │       │   ├── cert.pem
                │       │   └── private_key
                │       └── peer
                │           └── tlsca.crt
    ```

## 运行部署命令
在目标机器上执行`fabric-explorer`
```bash
cd /root/explorer
./fabric-explorer
```
如看到以下错误可忽略,原因为下载浏览器所需要的证书并没有通过mysql，当前是通过上一步手动下载至所需路径。mysql下载cert功能未来需要实现。
```bash
ERRO[0000] fail to connect to db: defaultdb, error:dial tcp 172.38.50.211:3306: connect: connection refused
ERRO[0000] fail to connect to database: defaultdb, error:dial tcp 172.38.50.211:3306: connect: connection refused
ERRO[0000] Fail to download org1 admin private key, error:open /root/explorer/orgazniations/org1/admin/private_key: no such file or directory
ERRO[0000] Fail to download cert
```

运行`./explorer -h` 可查看命令详细参数说明
```bash
cd /root/explorer
./fabric-explorer -h
```

## 验证执行结果
如果执行`explorer`命令后出现`Explorer started successfully.`则应该启动成功.

进一步验证,在浏览器中打开http://{目标机器ip}:8080/,如http://172.38.50.211:8080/,如果显示登录界面,则表明浏览器启动成功,浏览器用户名为exploreradmin, 密码为exploreradminpw


## 重启/关闭服务
通过docker-compose 进行服务重启/关闭即可
```bash
cd /root/explorer
docker-compose down --remove-orphan // 关闭
docker-compose up -d // 启动
```