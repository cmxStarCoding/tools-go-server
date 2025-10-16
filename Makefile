VERSION ?= latest
#从当前git仓库取版本
#VERSION ?= $(shell git describe --tags --always --dirty)

#给伪目标加上 .PHONY 声明，告诉 make 无论如何都要执行。
.PHONY: journey release install-server

app:
	@make -f deploy/mk/journey.mk VERSION=$(VERSION) release

#编译发布
release: app

#构建生产
install-server:
	cd ./deploy/script && chmod +x publish.sh && ./publish.sh