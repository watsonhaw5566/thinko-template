package main

import (
	thinkservice "github.com/watsonhaw5566/think-core/service"
	"think-go/router"
)

func main() {
	thinkservice.Run(router.BindController)
}
