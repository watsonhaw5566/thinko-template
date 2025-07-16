package main

import (
	tkService "github.com/watsonhaw5566/thinko/service"
	"think-go/router"
)

func main() {
	tkService.Run(router.BindController)
}
