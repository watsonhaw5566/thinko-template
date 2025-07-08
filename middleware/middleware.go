package middleware

import (
	"fmt"
	"github.com/watsonhaw5566/think-core"
	thinkconfig "github.com/watsonhaw5566/think-core/config"
	thinklog "github.com/watsonhaw5566/think-core/log"
	thinktoken "github.com/watsonhaw5566/think-core/token"
	"net/http"
	"time"
)

// Cors 跨域处理中间件
func Cors() think.HandlerFunc {
	return func(ctx *think.Context) {
		method := ctx.Request.Method
		ctx.Response.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRFToken, Authorization, Token")
		ctx.Response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Response.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-ControlAllow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Response.Header().Set("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.Fail("OPTIONS", think.FailOption{
				StatusCode: http.StatusNoContent,
				ErrorCode:  http.StatusNoContent,
			})
		}
		ctx.Next()
	}
}

// Logger 日志中间件
func Logger() think.HandlerFunc {
	return func(ctx *think.Context) {
		start := time.Now()
		ctx.Next()
		stop := time.Now()
		log := fmt.Sprintf("[ThinkGO] | %s | %d | %s | %s | %s | %s |", time.Now().Format("2006-01-02 15:04:05"), 200, ctx.ClientIP(), stop.Sub(start), ctx.Request.Method, ctx.Request.RequestURI)
		thinklog.Log().Info(log)
	}
}

// Authorization 验证中间件
func Authorization() think.HandlerFunc {
	return func(ctx *think.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		token := thinktoken.GetAuthorization(authorization)
		thinktoken.ParseToken(token, thinkconfig.Config.Get("jwtKey").String())
		ctx.Next()
	}
}
