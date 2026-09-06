package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	hoursOffset, actualMinutes := itom(m)
	actualHours := mod(h+hoursOffset, 24)
	return Clock{
		h: actualHours,
		m: actualMinutes,
	}
}

func itom(m int) (int, int) {
	hoursOffset := 0
	actualMinutes := m
	if m >= 60 {
		for actualMinutes >= 60 {
			hoursOffset = hoursOffset + 1
			actualMinutes = actualMinutes - 60
		}
	} else if m < 0 {
		for actualMinutes < 0 {
			hoursOffset = hoursOffset - 1
			actualMinutes = actualMinutes + 60
		}
	}
	return mod(hoursOffset, 24), actualMinutes
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
