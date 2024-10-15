package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/think-go/tg/tgutl"
)

func SayHello(ctx *gin.Context) {
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
	tgutl.Fail(ctx, "登录失败")
}
