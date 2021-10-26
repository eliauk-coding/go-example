## 构建镜像
1. 手动构建
   1. 下载 https://github.com/etcd-io/etcd/releases/download/v3.5.0/etcd-v3.5.0-linux-amd64.tar.gz
   2. 解压到当前目录
   3. 复制解压文件里的etcd和etcdctl文件到当前目录
   4. 构建etcd镜像:docker build -t etcd .
   
2. 拉取官方镜像（速度比较慢）
   1. docker pull quay.io/coreos/etcd
   
## 运行
两种方式：使用docker run命令或者编写docker-compose文件
1. 启动etcd(单服务)
   docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
   --name etcd etcd /usr/local/bin/etcd \
   -name etcd0 \
   -advertise-client-urls http://192.168.3.3:2379,http://192.168.3.3:4001 \
   -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
   -initial-advertise-peer-urls http://192.168.3.3:2380 \
   -listen-peer-urls http://0.0.0.0:2380 \
   -initial-cluster-token etcd-cluster-1 \
   -initial-cluster etcd0=http://192.168.3.3:2380 \
   -initial-cluster-state new
2. 启用yml文件
   运行docker-compose -f etcd-compose.yml up -d
   查看容器启动情况：docker ps -a