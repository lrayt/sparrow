package http_helper

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lrayt/sparrow"
	"github.com/lrayt/sparrow/core/runtime"
	"github.com/lrayt/sparrow/ts_error"
	"net/http"
	"reflect"
	"time"
)

var validate = validator.New()

func GinHandle(param any, serviceFunc ServiceFunc) gin.HandlerFunc {
	fn := reflect.ValueOf(serviceFunc)
	if fn.Kind() != reflect.Func {
		return NotServiceFuncErr
	}
	return func(c *gin.Context) {
		// log common fields
		fields := map[string]interface{}{
			"run_env":   sparrow.GRunEnv(),
			"version":   sparrow.GBuildVersion(),
			"module":    c.Request.URL.String(),
			"trace_id":  c.GetString(RequestId),
			"user_id":   c.GetString(UserId),
			"client_id": c.GetString(ClientId),
		}

		// logger
		logger := sparrow.GLoggerProvider().NewLogger(fields)

		// bind param
		if err := c.Bind(param); err != nil {
			c.JSON(http.StatusOK, &ts_error.BaseResponse{
				Code: ts_error.ParamInvalid,
				Msg:  logger.NewErrorF("param bind err", err).Error(),
			})
			return
		}

		// validate param
		if err := validate.Struct(param); err != nil {
			c.JSON(http.StatusOK, &ts_error.BaseResponse{
				Code: ts_error.ParamInvalid,
				Msg:  logger.NewErrorF("param validate err", err).Error(),
			})
			return
		}

		// call service
		res := fn.Call([]reflect.Value{
			reflect.ValueOf(&runtime.Context{Logger: logger, CTX: context.Background()}),
			reflect.ValueOf(param),
		})
		if len(res) <= 0 {
			c.JSON(http.StatusInternalServerError, &ts_error.BaseResponse{
				Code: ts_error.SystemErr,
				Msg:  logger.NewError("service func return nil").Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res[0].Interface())
	}
}

// CORSMiddleware 跨域设置
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Signature, X-authorize-uuid, Client-Id")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	logger := sparrow.GLoggerProvider().NewLogger(map[string]interface{}{"ping": "123"})
	return func(c *gin.Context) {
		logger.Info("----", map[string]interface{}{
			"start_time": time.Now().Unix(),
			"method":     c.Request.Method,
			"client_ip":  c.ClientIP(),
		})
	}
}
