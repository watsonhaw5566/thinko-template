package controller

import (
	"github.com/think-go/tg"
	"think-go/api"
)

type HomeData struct {
	Count    int
	Title    string
	SubTitle string
}

type User struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// HomeView 首页宣传页
func HomeView(ctx *tg.Context) {
	ctx.View("index.html", &HomeData{
		Count:    9863763,
		Title:    "ThinkGO",
		SubTitle: "欢迎使用ThinkGO框架",
	})
}

// SayHello GET接口测试
func SayHello(ctx *tg.Context) {
	name := ctx.GetQuery("name")
	ctx.Success(api.SayHelloRes{Say: name + "说hello"})
}

// SayWorld POST接口测试
func SayWorld(ctx *tg.Context) {
	var req api.SayHelloReq
	ctx.BindStructValidate(&req)
	ctx.Success(req.Name)
}

func UserList(ctx *tg.Context) {
	var user []User
	err := tg.Db("user").Select(&user)
	if err != nil {
		ctx.Fail(err.Error())
		return
	}
	ctx.Success(user)
}
