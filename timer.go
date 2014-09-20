package monotime

import (
	"time"
)

type Elapsed struct {
	t0 Time
	x0 int64
}

func NewElapsed() *Elapsed {
	return &Elapsed{Now(), time.Now().UnixNano()}
}

func (t Elapsed) Current() time.Duration {
	n := int64(Now() - t.t0)
	x := time.Now().UnixNano() - t.x0
	if x-2000000000 > n || x+2000000000 < n {
		return time.Duration(n)
	}
	return time.Duration(x)
}
