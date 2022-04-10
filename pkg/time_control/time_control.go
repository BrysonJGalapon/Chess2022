package time_control

import "fmt"

type TimeControl interface {
	GetHours() uint64
	GetMinutes() uint64
	GetSeconds() uint64
	GetIncrement() uint64
}

type timeControl struct {
	hours     uint64
	minutes   uint64
	seconds   uint64
	increment uint64
}

type TimeControlBuilder interface {
	Hours(hours uint64) TimeControlBuilder
	Minutes(uint64) TimeControlBuilder
	Seconds(uint64) TimeControlBuilder
	Increment(uint64) TimeControlBuilder
	Build() TimeControl
}

func (tc *timeControl) Hours(hours uint64) TimeControlBuilder {
	tc.hours = hours
	return tc
}

func (tc *timeControl) Minutes(minutes uint64) TimeControlBuilder {
	tc.minutes = minutes
	return tc
}

func (tc *timeControl) Seconds(seconds uint64) TimeControlBuilder {
	tc.seconds = seconds
	return tc
}

func (tc *timeControl) Increment(increment uint64) TimeControlBuilder {
	tc.increment = increment
	return tc
}

func (tc *timeControl) GetHours() uint64 {
	return tc.hours
}

func (tc *timeControl) GetMinutes() uint64 {
	return tc.minutes
}

func (tc *timeControl) GetSeconds() uint64 {
	return tc.seconds
}

func (tc *timeControl) GetIncrement() uint64 {
	return tc.increment
}

func (tc *timeControl) Build() TimeControl {
	return tc
}

func (tc *timeControl) String() string {
	return fmt.Sprintf("{hours: %d, minutes: %d, seconds: %d}", tc.hours, tc.minutes, tc.seconds)
}

func Builder() TimeControlBuilder {
	return &timeControl{}
}
