package controller

import (
	"fmt"
	"github.com/think-go/tg"
)

func SayHello(ctx *tg.Context) {
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
	ctx.Fail("登录失败")
}
