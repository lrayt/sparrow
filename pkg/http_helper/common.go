package http_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/lrayt/sparrow/ts_error"
	"net/http"
)

const (
	RequestId = "Request-Id"
	TraceId   = "Trace-Id"
	BizId     = "Biz-Id"
	UserId    = "User-Id"
	ClientId  = "Client-Id"
)

type ServiceFunc any

var NotServiceFuncErr = func(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, &ts_error.BaseResponse{
		Code: ts_error.SystemErr,
		Msg:  "not a service func",
	})
	c.Abort()
}
