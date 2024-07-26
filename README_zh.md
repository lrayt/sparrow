# 项目介绍
sparrow是一个轻量级golang后端开发骨架，麻雀虽小五脏俱全，sparrow集成了后端开发中常用的功能组件，同时这些组件也支持插拔替换

# 快速上手
## 1、安装
~~~shell
go get github.com/lrayt/sparrow@latest
mkdir you_server && cd you_server
go mod init github.com/xxx/you_server
~~~

## 2、工程结构
your-server
└─── cmd
│    └─── server_name
│         └─── main.go
│         └─── wire.go
│         └─── wire_gen.go
└─── app
│    └─── dao(数据操作层:如数据库、缓存)
│    └─── model(持久化数据的实体)
│    └─── service(业务处理)
│    └─── protobuf(协议)
│         └─── proto
│         └─── pb
│    └─── handler
│         └─── http_handler.go
│         └─── rpc_handler.go
│         └─── mq_handler.go
│         └─── ws_handler.go
└─── internal
│    └─── database
│    └─── http_client
│    └─── rpc_client
│    └─── file_resolver
│    └─── other...
└─── resource
│    └─── conf
│         └─── skeleton-local-conf.yaml
│         └─── skeleton-test-conf.yaml
│         └─── skeleton-prod-conf.yaml
│    └─── static（前端资源）
└─── pkg
│    └─── xxx_tools
└─── Makefile
└─── Dockerfile

## 3、构建
~~~shell
# 生成协议代码
make pb 
# 生成依赖注入代码
make wire 
# 项目构建
make build
# 构建docker镜像
make docker
~~~

# 架构设计
![struct_project](https://github.com/lrayt/sparrow/blob/master/docs/src/.vuepress/public/assets/image/struct_project.png)
# Feature
- 组件可插拔，项目运行依赖组件可替换更改；
- 业务调用维护单独context，全链路日志记录；