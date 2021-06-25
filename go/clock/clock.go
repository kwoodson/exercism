package clock

import (
	"fmt"
)

type Clock struct {
	hours   int
	minutes int
}

func mToH(m int) (int, int) {
	h := 0
	if m >= 60 {
		for i := m; i >= 60; i -= 60 {
			h += 1
		}
		m -= h * 60
	} else {
		i := 0
		for i = m; i < 0; i += 60 {
			h -= 1
		}
		m = i
	}
	return m, h
}

func New(h, m int) Clock {
	if m > 59 {
		tm, th := mToH(m)
		m = tm
		h += th
	} else if m < 0 {
		tm, th := mToH(m)
		m = tm
		h += th
	}
	if h < 0 {
		h = h % 24
		h = 24 + h
	}
	h = h % 24
	return Clock{hours: h, minutes: m}
}

func (c Clock) Add(minutes int) Clock {
	m, h := mToH(c.minutes + minutes)
	c.hours += h
	c.minutes = m

	return Clock{
		minutes: c.minutes,
		hours:   c.hours % 24,
	}
}

func (c Clock) Subtract(m int) Clock {
	i := 0
	for i = m; i > 60; i -= 60 {
		if c.hours == 0 {
			c.hours = 23
		} else {
			c.hours -= 1
		}
	}
	if c.minutes-i < 0 {
		if c.hours == 0 {
			c.hours = 23
		} else {
			c.hours -= 1
		}
		tm := c.minutes - i
		c.minutes = 60 + tm
	} else {
		c.minutes -= i
	}
	return Clock{
		minutes: c.minutes,
		hours:   c.hours,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
