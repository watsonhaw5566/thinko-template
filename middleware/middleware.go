package middleware

import (
	"github.com/think-go/tg"
	"net/http"
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
	//return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	return fmt.Sprintf(
	//		"[ThinkGO] | %s | %d | %s | %s | %s | %s | \n",
	//		params.TimeStamp.Format("2006-01-02 15:04:05"),
	//		params.StatusCode,
	//		params.ClientIP,
	//		params.Latency,
	//		params.Method,
	//		params.Path,
	//	)
	//})
	return func(ctx *tg.Context) {
		ctx.Next()
	}
}
