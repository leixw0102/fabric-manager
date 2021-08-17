#chaincode 文档
## 易链安装后


- 第一步

```
cd /root/fabric_networks/consortiums/SampleConsortium/chaincode/go
```

- 第二步
 doc 目录下chaincode/go/basic 或者 chaincode/go/abstroe 目录下所有文件拷贝到/root/fabric_networks/consortiums/SampleConsortium/chaincode/go 目录下
 目录组织结构可以自己定义。会影响后面打包环节

- 第三步
 进入cli 容器

 ```
 1. docker exec -it cli bash
   结果：
   [root@peer0 ~]# docker exec -it cli bash
   bash-5.0# 
 2. cd abstroe/go && go mod vendor &&  peer lifecycle chaincode package fabcar.tar.gz --path /opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode/go/ --lang golang --label abs_1 
	结果
	bash-5.0# ls
	abs.tar.gz   go.mod         go.sum         abstore.go        vendor
```
 - 第四步 
 1.安装
 ```
 peer lifecycle chaincode install abs.tar.gz
   结果:
    2021-08-10 01:30:55.089 UTC [cli.lifecycle.chaincode] submitInstallProposal -> INFO 035 Installed remotely: response:<status:200 payload:"\nFfbcar:27b3fc25d061670358796e25225876165d2bfeb8d408bc5ec9373c9781e8f80e\022\005fbcar" > 
   2021-08-10 01:30:55.089 UTC [cli.lifecycle.chaincode] submitInstallProposal -> INFO 036 Chaincode code package identifier: abs_1:27b3fc25d061670358796e25225876165d2bfeb8d408bc5ec9373c9781e8f80e
 ```
 2. 查询

 ```
 peer lifecycle chaincode queryinstalled
 结果：
 Installed chaincodes on peer:
Package ID: abs_1:27b3fc25d061670358796e25225876165d2bfeb8d408bc5ec9373c9781e8f80e, Label: abs_1
 ```

 - 第六步
  1. 认可
  ```
  peer lifecycle chaincode approveformyorg --channelID mychannel --name abs_1 --version 1.0 --init-required --package-id abs_1:51538734bd36e4fb3345739427a5de3cfbbee1fce019903d0bcae247ef424299 --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
   结果：
   2021-08-10 01:33:44.377 UTC [chaincodeCmd] ClientWait -> INFO 0a6 txid [4e74bb40e94154f0e20a297badfad69f8c2c8a9cb6433d660ea09f4fd574efc7] committed with status (VALID) at 
  ```
  2. 查看链码的状态是否就绪
  ```
  peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name abs_1 --version 1.0 --init-required --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json
   结果：
2021-08-10 01:34:25.239 UTC [msp.identity] Sign -> DEBU 034 Sign: digest: A438020D76A1533A3C8D9E6B3035ABE07055A33BF3C76FB692B9355D01175C89 
{
        "approvals": {
                "Org1MSP": true
        }
}

  ```
 - 第7步
  1.提交智能合约定义

  ```
  peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name abs_1 --version 1.0 --sequence 1 --init-required --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

  结果：
  2021-08-10 01:38:34.030 UTC [chaincodeCmd] ClientWait -> INFO 046 txid [236641c621f0e9a47950e0fc9e5a88fbd50d5b1cc21b9d4fd6931d604ce5d2bb] committed with status (VALID) at peer0.org1.example.com:7051 
  ```

  2.查看

  - 第8步
  1.初始化
 
```
peer chaincode invoke -o orderer.example.com:7050 --isInit --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n abs_1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["Init","a","100","b","100"]}'


结果
- 2021-08-10 04:44:47.830 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> DEBU 044 ESCC invoke result: version:1 response:<status:200 > payload:"\n '\034p\020u\234\n\220\301>\367G\004\345\365\217\373fm\017X\201}h<\007^\336\333\342d\257\022\232\001\n\204\001\0226\n\n_lifecycle\022(\n&\n namespaces/fields/abs_1/Sequence\022\002\010\002\022J\n\005abs_1\022A\n\022\n\020\000\364\217\277\277initialized\032\027\n\020\000\364\217\277\277initialized\032\0031.0\032\010\n\001a\032\003100\032\010\n\001b\032\003100\032\003\010\310\001\"\014\022\005abs_1\032\0031.0" endorsement:<endorser:"\n\007Org1MSP\022\246\006-----BEGIN CERTIFICATE-----\nMIICJzCCAc6gAwIBAgIQYFRBUXhNK9rs0xTYJtwg8DAKBggqhkjOPQQDAjBzMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eu\nb3JnMS5leGFtcGxlLmNvbTAeFw0yMTA4MDIwNzIyMDBaFw0zMTA3MzEwNzIyMDBa\nMGoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T\nYW4gRnJhbmNpc2NvMQ0wCwYDVQQLEwRwZWVyMR8wHQYDVQQDExZwZWVyMC5vcmcx\nLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8z75Iz8o+eix\n6bnU4CvpB0nGID5SzGCvmNUrbd5lLnb4RQUMzbSo8PfDDNjj6oBc4uxWvtgRtQVM\nntcjIFxx9qNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYDVR0j\nBCQwIoAgE16uIAADWZkAR1Dg6qMDfpZv9j7i80xCHuVPv/np94YwCgYIKoZIzj0E\nAwIDRwAwRAIgHguGnglsVMDYM+9xEJeOVZx6Qzh4+r3mu1hs3o3OU0UCIBRdKuMt\nAdAhuKPEpsHIO4WrIm56ios/1kTSjzNjTW4M\n-----END CERTIFICATE-----\n" signature:"0D\002 $j^~\374o;\262\305\254Wt6\264'*\366\004\343\257\326f\2239\227q\373\352'(\021G\002 wo!\235\305\237\274\177\373?e\226U\251h\221\0201\224\353\307)iG\355\276\272\242\321%\261#" > 
2021-08-10 04:44:47.830 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 045 Chaincode invoke successful. result: status:200 

```
 2.查询
 ```
peer chaincode query -C mychannel -n fbcar -c '{"Args":["GetAllAssets"]}'
- abstore.go 
peer chaincode query -C mychannel -n abs_1 -c '{"Args":["query","a"]}'

结果：
2021-08-10 04:46:28.381 UTC [msp.identity] Sign -> DEBU 034 Sign: digest: AA50BF27D695F477EBF48FD8015B5D8B460CD3091880F2945BA3E51F34A7CB2A 
100
```

- 第9步
1.测试
```
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n abs_1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}' --waitForEvent

结果:
2021-08-10 04:49:16.503 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 045 Chaincode invoke successful. result: status:200 
```