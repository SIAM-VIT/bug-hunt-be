package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	startTime time.Time
	duration  time.Duration
}

var GlobalTimer *Timer

func CreateTimer(duration time.Duration) *Timer {
	GlobalTimer = &Timer{
		startTime: time.Now(),
		duration:  duration,
	}
	return GlobalTimer
}

func (t *Timer) TimeLeft() (int, time.Time) {
	remaining := t.duration - time.Since(t.startTime)
	if remaining <= 0 {
		fmt.Println("Timer has expired!")
		return 0, t.startTime
	} else {
		fmt.Printf("Time remaining: %v\n", remaining.Round(time.Second))
		return int(remaining.Round(time.Second)), t.startTime
	}
}
