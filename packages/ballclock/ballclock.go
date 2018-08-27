//Package ballclock implements the logic necessary to simulate a ball clock.
package ballclock

import (
	"encoding/json"
	"fmt"
)

type clock struct {
	ballCount        int
	minutesSimulated int
	originalQueue    []int
	MinutesQueue     []int `json:"Min"`
	FiveMinuteQueue  []int `json:"FiveMin"`
	HoursQueue       []int `json:"Hour"`
	MainQueue        []int `json:"Main"`
}

//NewClock initializes a clock structure. It expects an integer parameter
//that has been previously validated to fall in the range of 27..127.
func NewClock(ballCount int) *clock {
	c := new(clock)
	c.ballCount = ballCount
	c.MinutesQueue = make([]int, 0)
	c.FiveMinuteQueue = make([]int, 0)
	c.HoursQueue = make([]int, 0)

	for i := 1; i <= int(ballCount); i++ {
		c.MainQueue = append(c.MainQueue, i)
		c.originalQueue = append(c.originalQueue, i)
	}
	return c
}

//RunMode1 runs the ballclock simulator until the order of the main queue returns to the original order.
//reports the state of each queue.
func (c *clock) RunMode1() string {
	cycleCount := 0
	finished := false
	for !finished {
		ball := c.getBallFromQueue()
		cycleCount++
		c.addMinute(ball)
		if len(c.MainQueue) == len(c.originalQueue) {
			finished = true
			for x, y := range c.MainQueue {
				if y != c.originalQueue[x] {
					finished = false
					break
				}
			}
		}
	}

	return fmt.Sprintf("%d balls cycle after %d days.", c.ballCount, cycleCount/(60*24))
}

//RunMode2 runs the ballclock simulator for a given number of "minutes" and then
//reports the state of each queue.
func (c *clock) RunMode2(minutes int) string {
	for i := 0; i < minutes; i++ {
		ball := c.getBallFromQueue()
		c.addMinute(ball)
	}
	result, err := c.printJSON()
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}

func (c *clock) getBallFromQueue() int {
	b := c.MainQueue[0]
	c.MainQueue = c.MainQueue[1:]
	return b
}

func (c *clock) addMinute(ball int) {
	c.minutesSimulated++
	if len(c.MinutesQueue) < 4 {
		c.MinutesQueue = append(c.MinutesQueue, ball)
		return
	}
	//dump minutes
	for i := len(c.MinutesQueue) - 1; i >= 0; i-- {
		c.MainQueue = append(c.MainQueue, c.MinutesQueue[i])

	}
	c.MinutesQueue = c.MinutesQueue[:0]
	c.addFiveMinute(ball)
}

func (c *clock) addFiveMinute(ball int) {
	if len(c.FiveMinuteQueue) < 11 {
		c.FiveMinuteQueue = append(c.FiveMinuteQueue, ball)
		return
	}
	//dump fiveminutes
	for i := len(c.FiveMinuteQueue) - 1; i >= 0; i-- {
		c.MainQueue = append(c.MainQueue, c.FiveMinuteQueue[i])

	}
	c.FiveMinuteQueue = c.FiveMinuteQueue[:0]
	c.addHour(ball)
}

func (c *clock) addHour(ball int) {
	if len(c.HoursQueue) < 11 {
		c.HoursQueue = append(c.HoursQueue, ball)
		return
	}
	//dump hours
	for i := len(c.HoursQueue) - 1; i >= 0; i-- {
		c.MainQueue = append(c.MainQueue, c.HoursQueue[i])

	}
	c.HoursQueue = c.HoursQueue[:0]
	c.MainQueue = append(c.MainQueue, ball) //trigger ball falls back into main queue
}

//used for debugging
func (c *clock) status() string {
	s := fmt.Sprintf("      minutes: %4d \t queue: %v\n", len(c.MinutesQueue), c.MinutesQueue)
	s += fmt.Sprintf(" five minutes: %4d \t queue: %v\n", len(c.FiveMinuteQueue)*5, c.FiveMinuteQueue)
	s += fmt.Sprintf("        hours: %4d \t queue: %v\n", len(c.HoursQueue)+1, c.HoursQueue)
	s += fmt.Sprintf("minutes simul: %6d \t Time: %02d:%02d:00 \t days: %d\n\n", c.minutesSimulated, len(c.HoursQueue)+1, (len(c.FiveMinuteQueue)*5)+len(c.MinutesQueue), int(float32(c.minutesSimulated)/1440.)+1)
	s += fmt.Sprintf("   main queue: %4d \t queue: %v\n", len(c.MainQueue), c.MainQueue)
	s += fmt.Sprintf("initial queue: %4d \t queue: %v\n", len(c.originalQueue), c.originalQueue)
	return s
}

func (c *clock) printJSON() (string, error) {
	j, err := json.Marshal(c)
	if err != nil {
		return "", err

	}
	return string(j), nil
}
