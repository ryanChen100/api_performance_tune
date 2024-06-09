package gzip

import (
	"api/performance_tune/enum"
	"api/performance_tune/pkg/apilib"
	"api/performance_tune/pkg/config"
	"strconv"
	"testing"
)

func TestGzip(t *testing.T) {
	config.ConfigTestInit()
	apilib.GzipPostItem(enum.Normal, "1")
	apilib.GzipGetItem(enum.Normal, "1")
	apilib.GzipPutItem(enum.Normal, "1")
	apilib.GzipDeleteItem(enum.Normal, "1")
	apilib.GzipGetItem(enum.Normal, "1")
}

func BenchmarkGzip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		apilib.GzipPostItem(enum.Normal, id)
		apilib.GzipGetItem(enum.Normal, id)
		apilib.GzipPutItem(enum.Normal, id)
		apilib.GzipDeleteItem(enum.Normal, id)
	}
}
