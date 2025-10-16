# 读取外部传入版本号参数
VERSION=$1

# 如果没有传，则默认使用 latest
if [ -z "$VERSION" ]; then
  VERSION="latest"
fi

echo "🚀 当前部署版本号: $VERSION"

need_start_server_shell=(
  journey.sh
)

for i in ${need_start_server_shell[@]} ; do
    #chmod +x $i
    ./$i "$VERSION"
done

docker ps

#docker exec -it etcd etcdctl get --prefix ""