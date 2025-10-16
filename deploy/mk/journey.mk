VERSION=latest

#需要修改这里
SERVER_NAME=journey

# 测试环境配置
# docker的镜像发布地址，person-collect是命名空间、SERVER_NAME是藏昆明
DOCKER_REPO_TEST=crpi-5vfnm5k3tdyjsxrh.cn-hangzhou.personal.cr.aliyuncs.com/person-collect/${SERVER_NAME}
# 测试版本
VERSION_TAG=$(VERSION)
# 编译的程序名称
APP_NAME=${SERVER_NAME}

# 测试下的编译文件
DOCKER_FILE=./deploy/dockerfile/Dockerfile_app

# 测试环境的编译发布
build:

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME} ./main.go
	docker build . -f ${DOCKER_FILE} --no-cache -t ${APP_NAME}

# 镜像的测试标签
tag:

	@echo 'create tag ${VERSION_TAG}'
	docker tag ${APP_NAME} ${DOCKER_REPO_TEST}:${VERSION_TAG}

publish:

	@echo 'publish ${VERSION_TAG} to ${DOCKER_REPO_TEST}'
	docker push $(DOCKER_REPO_TEST):${VERSION_TAG}

release: build tag publish
