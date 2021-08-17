# 流程

## 前置条件
- etcd
    - server->agent路径：fabric-manager/agent/{orgDomain}/{peerIp}/{cmd}
    - agent->server: fabric-manager/server/{cmd}
- mysql
    - agent执行结果通过api server写入mysql, 方便cmd server检查执行结果

- 组网
    - 用户通过server注册身份(input:username, password,output:userid)
    - 界面创建org
    - agent通过userid,orgDomain (写入boltdb)启动
    - 上报ip/（心跳）/执行结果

## Server
### 1. CreateOrg
    - orgDomain
    - caAddress
    - identities[0]
        - name
        - password
        - idtype
    logic:
        - 编码
        - 写mysql
    NOTE: 需要每个peer的端口号，那个peer是orderer

    
### 2. CreateConsortium
    - configtx.yaml
        - ordererType
        - orgDomain(s)
    - consortium name

### 3. StartNetwork

        
### 4. CreateChannel
    - configtx.yaml
        - channel name

## Client
### 1. CreateOrg
    input:
        - consortium
        - orgDomain
        - caAddress
        - identities[0]
            - name
            - password
            - idtype
    logic:
        - 解码
        - 写boltdb
        - 上报执行结果

### 2. CreateConsortium
    - configtx.yaml
        - ordererType
        - orgDomain(s)
    - consortium name

### 3. StartNetwork
    - 生成docker-compose
    - 执行docker-comopse

### 4. CreateChannel
    - 

# 代码结构

## agent  
 - 负责所有命令的最终执行器
 - 接收etcd指令
 - 下载cryptogen 指定证书

## etcd 
 - 消息中间件

## server
 - 提供api与用户交互
 - 下发区块链指令

## 区块链浏览器
 - 区块链组网之后可以部署区块链浏览器，方便查看组网信息

## cert
 - cryptogen 证书服务
 - 提供生成证书
 - 提供下载证书接口


 # 分支说明
 ## cryptogen
  - 证书使用cryptogen 实现区块链
 ## master 
  - ca 证书实现，目前有问题