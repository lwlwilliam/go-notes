package timer

import (
	"time"
)

// StopWatch 是简易的时钟功能，它的零值与总时长为 0 的停止的时钟一样
type StopWatch struct {
	start int	//time.Time
	total time.Duration
	running bool
}

// Start 用来开启时钟
func (s *StopWatch) Start() {
	if !s.running {
		s.start = time.Now().Nanosecond()
		s.running = true
	}
}

func (s *StopWatch) Total() int{
	var total int
	if s.running {
		total = time.Now().Nanosecond() - s.start
		s.running = false
	}

	return total
}
