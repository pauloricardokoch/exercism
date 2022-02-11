package clock

import "fmt"

const (
	hourMinutes = 60
	dayMinutes  = hourMinutes * 24
)

// Define the Clock type here.
type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := ((h*hourMinutes+m)%dayMinutes + dayMinutes) % dayMinutes
	return Clock{minutes: minutes}
}

func (c Clock) Add(m int) Clock {
	minutes := c.minutes + m
	return New(minutes/hourMinutes, minutes%hourMinutes)
}

func (c Clock) Subtract(m int) Clock {
	minutes := c.minutes - m
	return New(minutes/hourMinutes, minutes%hourMinutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/hourMinutes, c.minutes%hourMinutes)
}
