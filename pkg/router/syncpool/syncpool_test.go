package syncpool

import (
	"api/performance_tune/enum"
	"api/performance_tune/pkg/apilib"
	"api/performance_tune/pkg/config"
	"strconv"
	"testing"
)

func TestSyncPool(t *testing.T) {
	config.ConfigTestInit()
	apilib.PostItem(enum.SyncPool, "1")
	apilib.GetItem(enum.SyncPool, "1")
	apilib.PutItem(enum.SyncPool, "1")
	apilib.DeleteItem(enum.SyncPool, "1")
	apilib.GetItem(enum.SyncPool, "1")
}

func BenchmarkSyncPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		apilib.PostItem(enum.SyncPool, id)
		apilib.GetItem(enum.SyncPool, id)
		apilib.PutItem(enum.SyncPool, id)
		apilib.DeleteItem(enum.SyncPool, id)
	}
}
