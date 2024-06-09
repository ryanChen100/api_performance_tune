package router

import (
	"api/performance_tune/pkg/config"
	"api/performance_tune/pkg/router/fast"
	"api/performance_tune/pkg/router/gocache"
	"api/performance_tune/pkg/router/normal"
	"api/performance_tune/pkg/router/syncpool"
	"compress/gzip"
	"io"
	"log"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
)

func RouterAdmin() {

	r := gin.Default()
	normal.InitRouter(r)
	gocache.InitGoCacheRouter(r)
	syncpool.InitSyncPoolRouter(r)
	r.Run(":" + config.GetSetting().GinPort)

}

func FastRouterAdmin() {
	fr := fasthttprouter.New()
	fast.InitFastHttpRouter(fr)
	if err := fasthttp.ListenAndServe(":"+config.GetSetting().FastPort, fr.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func GzipRouterAdmin() {
	r := gin.Default()
	r.Use(gzipMiddleware)
	normal.InitRouter(r)
	gocache.InitGoCacheRouter(r)
	syncpool.InitSyncPoolRouter(r)
	r.Run(":" + config.GetSetting().GzipPort)
}

// 手动实现 Gzip 压缩的中间件
func gzipMiddleware(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
		gzipWriter := gzip.NewWriter(c.Writer)
		defer gzipWriter.Close()
		c.Header("Content-Encoding", "gzip")
		c.Writer = &gzipResponseWriter{
			Writer:         gzipWriter,
			ResponseWriter: c.Writer,
		}
	}
	c.Next()
}

// 自定义的 gzip 响应写入器，嵌入 http.ResponseWriter
type gzipResponseWriter struct {
	io.Writer
	gin.ResponseWriter
}

// 实现 Write 方法以满足 io.Writer 接口
func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
