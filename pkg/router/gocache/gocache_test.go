package gocache

import (
	"api/performance_tune/enum"
	"api/performance_tune/pkg/apilib"
	"strconv"
	"testing"
)

func TestGocache(t *testing.T) {
	apilib.PostItem(enum.GoCache, "1")
	apilib.GetItem(enum.GoCache, "1")
	apilib.PutItem(enum.GoCache, "1")
	apilib.DeleteItem(enum.GoCache, "1")
	apilib.GetItem(enum.GoCache, "1")
}

func BenchmarkGocache(b *testing.B) {
	// time.Sleep(10 * time.Second)
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		apilib.PostItem(enum.GoCache, id)
		apilib.GetItem(enum.GoCache, id)
		apilib.PutItem(enum.GoCache, id)
		apilib.DeleteItem(enum.GoCache, id)
	}
}
