package timer

import (
	"testing"
	"time"
)

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer := Get(time.Second)
		//<-timer.C
		Put(timer)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timer := time.NewTimer(time.Second)
		//<-timer.C
		timer.Stop()
		//timer.Reset(0)
	}
	b.StopTimer()
	b.ReportAllocs()
}
