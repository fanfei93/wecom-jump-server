package main

import (
	"wecom-jump-server/router"
)

func main() {
	r := router.InitRouter()
	err := r.Run(":8083")
	if err != nil {
		panic(err)
	}
}
