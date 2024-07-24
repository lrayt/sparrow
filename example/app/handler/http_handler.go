package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/sparrow/example/app/protobuf/pb"
	"github.com/lrayt/sparrow/example/app/service"
	"github.com/lrayt/sparrow/helper"
	"io"
	"net/http"
	"time"
)

type HttpHandler struct {
	api              *gin.RouterGroup
	srv              *http.Server
	orderInfoService *service.OrderInfoService
}

func NewHttpHandler(orderInfoService *service.OrderInfoService) *HttpHandler {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()
	//r.Use(http_helper.LoggerMiddleware())
	return &HttpHandler{
		api:              r.Group("/api/v1"),
		srv:              &http.Server{Addr: ":8080", Handler: r},
		orderInfoService: orderInfoService,
	}
}

func (h HttpHandler) orderRouter() {
	rg := h.api.Group("/order")
	{
		rg.POST("/create", helper.GinHandle(&pb.OrderCreateRequest{}, h.orderInfoService.CreateOrder))
	}
}

func (h *HttpHandler) Run() error {
	h.api.Use(helper.LoggerMiddleware())
	h.orderRouter()
	if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (h HttpHandler) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return h.srv.Shutdown(ctx)
}
