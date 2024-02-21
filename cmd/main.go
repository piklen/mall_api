package main

import (
	"mall_api/routes"
)

func main() {
	r := routes.NewRouter()
	_ = r.Run(":7999")
}
