package fast

import (
	"api/performance_tune/enum"
	"api/performance_tune/pkg/apilib"
	"api/performance_tune/pkg/config"
	"strconv"
	"testing"
)

func TestFastRouter(t *testing.T) {
	config.ConfigTestInit()
	apilib.FastPostItem(enum.FastHttp, "1")
	apilib.FastGetItem(enum.FastHttp, "1")
	apilib.FastPutItem(enum.FastHttp, "1")
	apilib.FastDeleteItem(enum.FastHttp, "1")
	apilib.FastGetItem(enum.FastHttp, "1")
}

func BenchmarkFastHttp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		apilib.FastPostItem(enum.FastHttp, id)
		apilib.FastGetItem(enum.FastHttp, id)
		apilib.FastPutItem(enum.FastHttp, id)
		apilib.FastDeleteItem(enum.FastHttp, id)
	}
}
