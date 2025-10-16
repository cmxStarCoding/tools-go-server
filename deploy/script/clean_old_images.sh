#!/bin/bash
clean_old_images() {
    IMAGE_REPO=$1
    CURRENT_VERSION=$2
    #第三个参数不传默认为3
    KEEP_COUNT=${3:-3}

    if [ -z "$IMAGE_REPO" ] || [ -z "$CURRENT_VERSION" ]; then
        echo "Usage: clean_old_images <IMAGE_REPO> <CURRENT_VERSION> [KEEP_COUNT]"
        return 1
    fi

    if [ "$CURRENT_VERSION" != "latest" ]; then
        echo "🧹 清理旧版本镜像，只保留最新 $KEEP_COUNT 个版本..."

        images=$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "${IMAGE_REPO}:v" | sort -V)

        delete_images=($(echo "$images" | head -n -$KEEP_COUNT))

        for img in "${delete_images[@]}"; do
            if docker images --format '{{.Repository}}:{{.Tag}}' | grep -q "^${img}$"; then
                echo "🗑️ 删除旧镜像: $img"
                docker rmi "$img"
            else
                echo "⚠️ 镜像不存在，不删除: $img"
            fi
        done
    fi

    echo "✅ 镜像管理完成: $IMAGE_REPO"
}
