# ETCD 部署

## 单机部署
```bash
sudo rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
  docker rmi gcr.io/etcd-development/etcd:v3.2.32 || true && \
  docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name etcd-gcr-v3.2.32 \
  quay.io/coreos/etcd:v3.2.32 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new
```
## 集群部署
三台机器ip假设为s1：`172.38.80.104`，s2:`172.38.80.103`, s3:`172.38.50.210`, 不同机器`advertise-client-urls`和`initial-advertise-peer-urls` 设为本机ip, `initial-cluster`设为各个机器自己的ip地址，所有端口均不变，具体如下所示：

第一台机器执行：
```bash
sudo rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
docker run \
-p 2379:2379 \
-p 2380:2380 \
--mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
--name etcd-gcr-v3.2.32 \
gcr.io/etcd-development/etcd:v3.2.32 \
/usr/local/bin/etcd \
--name s1 \
--data-dir /etcd-data \
--listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://172.38.80.104:2379 \
--listen-peer-urls http://0.0.0.0:2380 \
--initial-advertise-peer-urls http://172.38.80.104:2380 \
--initial-cluster s1=http://172.38.80.104:2380,s2=http://172.38.80.103:2380,s3=http://172.38.50.210:2380 \
--initial-cluster-token tkn \
--initial-cluster-state new
```

第二台机器执行：
```bash
sudo rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
docker run \
-p 2379:2379 \
-p 2380:2380 \
--mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
--name etcd-gcr-v3.2.32 \
gcr.io/etcd-development/etcd:v3.2.32 \
/usr/local/bin/etcd \
--name s2 \
--data-dir /etcd-data \
--listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://172.38.80.103:2379 \
--listen-peer-urls http://0.0.0.0:2380 \
--initial-advertise-peer-urls http://172.38.80.103:2380 \
--initial-cluster s1=http://172.38.80.104:2380,s2=http://172.38.80.103:2380,s3=http://172.38.50.210:2380 \
--initial-cluster-token tkn \
--initial-cluster-state new
```

第三台机器执行：
```bash
sudo rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
docker run \
-p 2379:2379 \
-p 2380:2380 \
--mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
--name etcd-gcr-v3.2.32 \
gcr.io/etcd-development/etcd:v3.2.32 \
/usr/local/bin/etcd \
--name s3 \
--data-dir /etcd-data \
--listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://172.38.50.210:2379 \
--listen-peer-urls http://0.0.0.0:2380 \
--initial-advertise-peer-urls http://172.38.50.210:2380 \
--initial-cluster s1=http://172.38.80.104:2380,s2=http://172.38.80.103:2380,s3=http://172.38.50.210:2380 \
--initial-cluster-token tkn \
--initial-cluster-state new
```