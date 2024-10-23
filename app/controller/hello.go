package controller

import (
	"github.com/think-go/tg"
)

type HomeData struct {
	Count    int
	Title    string
	SubTitle string
}

// HomeView 首页宣传页
func HomeView(ctx *tg.Context) {
	ctx.View("index.html", &HomeData{
		Count:    9863763,
		Title:    "ThinkGO",
		SubTitle: "欢迎使用ThinkGO框架",
	})
}
