package controller

import (
	"github.com/watsonhaw5566/thinko"
)

// HomeData 首页宣传页测试数据
type HomeData struct {
	Count    int
	Title    string
	SubTitle string
}

// HomeView 首页宣传页
func HomeView(ctx *thinko.Context) {
	ctx.View("index.html", &HomeData{
		Count:    9863763,
		Title:    "Thinko",
		SubTitle: "欢迎使用 Thinko 框架",
	})
}

// SayHello GET接口测试
func SayHello(ctx *thinko.Context) {
	name := ctx.GetQuery("name")
	ctx.Success(name + "说hello")
}
