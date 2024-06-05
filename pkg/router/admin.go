package router

import (
	"api/performance_tune/pkg/config"
	"api/performance_tune/pkg/router/gocache"
	"api/performance_tune/pkg/router/normal"
	"api/performance_tune/pkg/router/syncpool"

	"github.com/gin-gonic/gin"
)

func RouterAdmin() {

	r := gin.Default()
	normal.InitRouter(r)
	gocache.InitGoCacheRouter(r)
	syncpool.InitSyncPoolRouter(r)
	r.Run(":" + config.GetSetting().Port)
}
