package clock

import (
	"fmt"
	"strconv"
)

type Clock struct {
	hours   int
	minutes int
}

func (c Clock) String() string {
	var hoursString, minutesString string

	if c.hours < 10 {
		hoursString = fmt.Sprintf("0%v", strconv.Itoa(c.hours))
	} else {
		hoursString = strconv.Itoa(c.hours)
	}

	if c.minutes < 10 {
		minutesString = fmt.Sprintf("0%v", strconv.Itoa(c.minutes))
	} else {
		minutesString = strconv.Itoa(c.minutes)
	}

	return fmt.Sprintf("%v:%v", hoursString, minutesString)
}

func New(h int, m int) Clock {
	totalMinutes := (h * 60) + m
	var totalHours int

	if totalMinutes < 0 {
		diffHours := (totalMinutes / 60) % 24
		totalHours = 24 + diffHours
	} else {
		totalHours = (totalMinutes / 60) % 24
	}

	m = totalMinutes % 60

	if m < 0 {
		m = 60 + m
		totalHours--
	}

	return Clock{totalHours, m}
}

func (c Clock) Add(m int) Clock {
	newMinutes := c.minutes + m
	newClock := New(c.hours, newMinutes)

	return newClock
}

func (c Clock) Subtract(m int) Clock {
	newMinutes := c.minutes - m
	newClock := New(c.hours, newMinutes)

	return newClock
}
