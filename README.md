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
