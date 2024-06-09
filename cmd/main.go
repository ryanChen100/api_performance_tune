package main

import (
	"api/performance_tune/pkg/config"
	router "api/performance_tune/pkg/router"
)

func main() {
	config.ConfigInit()
	go router.FastRouterAdmin()
	go router.RouterAdmin()
	go router.GzipRouterAdmin()
	for {
	}
}
