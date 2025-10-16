### 工具项目
后端Go服务实现 ，阿狸工具-专注提高工作效率

### 运行环境
- Golang1.19.13(需要开启go.mod)
- Mysql8.0+
- Redis5.0+
- Rabbitmq


### 安装

使用以下命令安装：
```bash
git clone git@github.com:cmxStarCoding/tools-go-server.git

go mod download
```
配置mysql以及redis：
```bash
#配置redis、mysql(在common目录下新建config.ini文件写入如下配置,可复制config_example.ini文件内容)
[app]
domain = http://127.0.0.1:8083
name = 阿狸工具

[db]
host = 127.0.0.1
database = tools
username = root
password =
port = 3380

[redis]
host = 127.0.0.1
password =
port = 6379
```

编译运行：
```bash
cd tools-go-server/core

#编译二进制文件
go build -o core

#运行二进制文件
./core
```
### 作者
- 作者名字：崔明星
- 电子邮件：15638276200@163.com

### 贡献
如果你想为项目做出贡献，请通过邮箱15638276200@163.com联系。

