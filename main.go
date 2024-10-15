package main

import (
	"github.com/think-go/tg/tgsv"
	"think-go/router"
)

func main() {
	tgsv.Run(router.BindController)
}
