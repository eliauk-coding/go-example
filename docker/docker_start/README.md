# 1. 构建镜像
    docker build -t docker_go:v1 .
# 2. 验证镜像
    docker images
# 3. 创建并运行一个新容器
    docker run -p 8080:8080 docker_go