---
title: 开发环境
icon: gears
order: 5
category:
  - Guide
tag:
  - disable

navbar: false

breadcrumb: false
pageInfo: false
contributors: false
editLink: false
lastUpdated: false
prev: false
next: false
comment: false
footer: false

backtotop: false
---

## 依赖注入（wire）

~~~ shell
  # 安装wire
  go install github.com/google/wire/cmd/wire@latest
  # 注意将GOPATH/bin添加到path环境变量中
  
  # 测试输出
  wire help
  
  # 修改privader
  # 生成代码
  make wire
~~~

## 接口协议

~~~ shell
# 1、安装protoc
# 前往https://github.com/protocolbuffers/protobuf/releases下载对应系统的zip包，然后解压并将bin路径添加到path中
# 验证安装是否成功
protoc --version

# 2、安装protoc-gen-go
go install  google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 注意：需要将$GOPATH/bin添加到环境变量path中
# 验证安装是否成功
protoc-gen-go --version

# 3、安装protoc-go-inject-tag
go install github.com/favadi/protoc-go-inject-tag@latest
# 验证安装是否成功
protoc-go-inject-tag -h

# 4、安装protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# 验证安装是否成功
protoc-gen-go-grpc --version
~~~