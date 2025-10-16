#!/bin/bash

# 从第一个参数获取版本号，如果没有就默认 latest
VERSION=$1
if [ -z "$VERSION" ]; then
  VERSION="latest"
fi

SERVER_NAME="journey"

#person-collect是命名空间，journey是仓库名
reso_addr="crpi-5vfnm5k3tdyjsxrh.cn-hangzhou.personal.cr.aliyuncs.com/person-collect/${SERVER_NAME}"
tag=$VERSION

container_name=${SERVER_NAME}

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-chat -v /easy-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8083 --net cmxnet --name=${container_name} -d ${reso_addr}:${tag}
echo "🚀 启动服务: ${container_name} (版本号: ${tag})"

# 调用公共镜像清理脚本
source ./clean_old_images.sh
clean_old_images "$reso_addr" "$VERSION" 3