package normal

import (
	"api/performance_tune/enum"
	"api/performance_tune/pkg/apilib"
	"strconv"
	"testing"
)

func TestNormal(t *testing.T) {
	apilib.PostItem(enum.Normal, "1")
	apilib.GetItem(enum.Normal, "1")
	apilib.PutItem(enum.Normal, "1")
	apilib.DeleteItem(enum.Normal, "1")
	apilib.GetItem(enum.Normal, "1")
}

func BenchmarkNormal(b *testing.B) {
	// time.Sleep(10 * time.Second)
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := strconv.Itoa(i)
		apilib.PostItem(enum.Normal, id)
		apilib.GetItem(enum.Normal, id)
		apilib.PutItem(enum.Normal, id)
		apilib.DeleteItem(enum.Normal, id)
	}
}
