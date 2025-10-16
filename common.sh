#编译发布(注意是私有化部署,需要本地docker登录阿里云制品仓库，且创建好命名空间和仓库)
make release VERSION=v1.0.2

#部署安装
make install-server VERSION=v1.0.5