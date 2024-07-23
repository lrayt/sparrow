//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lrayt/sparrow/example/Internal/database"
	"github.com/lrayt/sparrow/example/app/dao"
	"github.com/lrayt/sparrow/example/app/handler"
	"github.com/lrayt/sparrow/example/app/service"
)

var InternalProvider = wire.NewSet(
	database.NewDBManager,
)

// DaoProvider 数据库操作
var DaoProvider = wire.NewSet(
	dao.NewOrderInfoDao,
)

// ServiceProvider 业务处理
var ServiceProvider = wire.NewSet(
	service.NewOrderInfoService,
)

// HandlerProvider 获取参数
var HandlerProvider = wire.NewSet(
	handler.NewHttpHandler,
)

func InitExampleServer() (*ExampleServer, func(), error) {
	panic(wire.Build(InternalProvider, DaoProvider, ServiceProvider, HandlerProvider, NewExampleServer))
}
