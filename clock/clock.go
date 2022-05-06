package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

func clock(h, m int) (int, int) {
	t := ((h%24+24)%24)*60 + m
	h = ((t / 60 % 24) + 24) % 24
	if t < 0 {
		h--
	}
	if h < 0 {
		h = 24 + h
	}
	m = (t%60 + 60) % 60
	return h, m
}

// New returns a new clock struct.
func New(h, m int) Clock {
	h, m = clock(h, m)
	return Clock{Hours: h, Minutes: m}
}

// Add adds minutes to a clock.
func (c Clock) Add(m int) Clock {
	var h int
	h, m = clock(c.Hours, c.Minutes+m)
	return Clock{Hours: h, Minutes: m}
}

// Add substracts minutes from a clock.
func (c Clock) Subtract(m int) Clock {
	var h int
	h, m = clock(c.Hours, c.Minutes-m)
	return Clock{Hours: h, Minutes: m}
}

// String returns a string representation of the clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
