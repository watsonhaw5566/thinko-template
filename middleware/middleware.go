package middleware

import (
	"fmt"
	"github.com/think-go/tg"
	"github.com/think-go/tg/tglog"
	"net/http"
	"time"
)

// Cors 跨域处理中间件
func Cors() tg.HandlerFunc {
	return func(ctx *tg.Context) {
		method := ctx.Request.Method
		ctx.Response.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRFToken, Authorization, Token")
		ctx.Response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Response.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-ControlAllow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Response.Header().Set("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.Fail("OPTIONS", tg.FailOptions{
				StatusCode: http.StatusNoContent,
				ErrorCode:  http.StatusNoContent,
			})
		}
		ctx.Next()
	}
}

// Logger 日志中间件
func Logger() tg.HandlerFunc {
	return func(ctx *tg.Context) {
		start := time.Now()
		ctx.Next()
		stop := time.Now()
		log := fmt.Sprintf("[ThinkGO] | %s | %d | %s | %s | %s | %s |", time.Now().Format("2006-01-02 15:04:05"), 200, ctx.ClientIP(), stop.Sub(start), ctx.Request.Method, ctx.Request.RequestURI)
		tglog.Log().Info(log)
	}
}
