// Monotonically increasing time.
package monotime

import (
	"sync/atomic"
	"time"
)

type Time int64

const Max = Time(0x7fffffffffffffff)

var ctime int64

// Current monotonic time.
func Now() Time {
	return Time(atomic.LoadInt64(&ctime))
}

// Current monotonic time plus a duration.
func NowPlus(d time.Duration) Time {
	return Time(atomic.LoadInt64(&ctime)) + Time(d)
}

// Current unique monotonic time. All calls of NowUnique return distinct unique values.
func NowUnique() Time {
	return Time(atomic.AddInt64(&ctime, 1))
}

func init() {
	atomic.StoreInt64(&ctime, time.Now().UnixNano())
	go loop()
}

func loop() {
	tick := time.NewTicker(time.Second)
	for {
		<-tick.C
		atomic.AddInt64(&ctime, int64(time.Second))
	}
}
