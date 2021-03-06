package pkg

import (
	"strconv"
	"time"
)

type StopWatch struct {
	start time.Time
	stop  time.Time
}

func NewStart() StopWatch {
	stopWatch := &StopWatch{time.Time{}, time.Time{}}
	return stopWatch.Start()
}

func (s *StopWatch) Start() StopWatch {
	s.start = time.Now()
	return *s
}

func (s *StopWatch) Stop() StopWatch {
	s.stop = time.Now()
	return *s
}

func (s *StopWatch) duration() time.Duration {
	return s.stop.Sub(s.start)
}

func (s *StopWatch) Show() string {
	hh := s.duration() / time.Hour
	mm := s.duration() / time.Minute
	ss := s.duration() / time.Second
	diff := strconv.Itoa(int(hh)) + "h " + strconv.Itoa(int(mm)) + "m " + strconv.Itoa(int(ss)) + "s"
	return diff
}

// ref: https://github.com/bradhe/stopwatch
