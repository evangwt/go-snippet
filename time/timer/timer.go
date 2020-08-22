package timer

import (
	"sync"
	"time"
)

var p = &pool{
	p: sync.Pool{
		New: func() interface{} {
			return time.NewTimer(0)
		},
	},
}

func Get(d time.Duration) *time.Timer {
	return p.Get(d)
}

func Put(timer *time.Timer) {
	p.Put(timer)
}

type pool struct {
	p sync.Pool
}

func (p *pool) Get(d time.Duration) *time.Timer {
	timer := p.p.Get().(*time.Timer)
	drain(timer)
	timer.Reset(d)
	return timer
}

func (p *pool) Put(timer *time.Timer) {
	drain(timer)
	p.p.Put(timer)
}

func drain(timer *time.Timer) {
	if !timer.Stop() {
		select {
		case <-timer.C: // try to drain from the channel
		default:
		}
	}
}
