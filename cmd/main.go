package main

import (
	"api/performance_tune/pkg/config"
	router "api/performance_tune/pkg/router"
)

func main() {
	config.ConfigInit()
	router.RouterAdmin()
}
