package ballclock

import "fmt"

//MinBalls is the minimum amount of balls supported by the ball clock
const MinBalls = 27

//MaxBalls is the maximum amount of balls supported by the ball clock
const MaxBalls = 127

//minTrackCap is the ball capacity of the minute track
const minTrackCap int = 4

//fiveAndHourTrackCap is the ball capacity of the five minute and hour track
const fiveAndHourTrackCap int = 11

//Clock is an instance of a ballclock
type Clock struct {
	timeTracks [][]int
	ballQueue  []int
	ballCount  int
}

//NewClock returns a new instance of the clock type
func NewClock(ballCount int) (*Clock, error) {
	if ballCount < MinBalls || ballCount > MaxBalls {
		return nil, fmt.Errorf("Invalid ball count. %v provided, must be between %v and %v", ballCount, MinBalls, MaxBalls)
	}

	c := &Clock{
		ballQueue: make([]int, ballCount, ballCount+1152+253+11),
		timeTracks: [][]int{
			make([]int, 0, minTrackCap),
			make([]int, 0, fiveAndHourTrackCap),
			make([]int, 0, fiveAndHourTrackCap)},
		ballCount: ballCount,
	}

	for i := 0; i < ballCount; i++ {
		c.ballQueue[i] = i
	}

	return c, nil
}

//GetTrackState returns a State object depicting the internal ball structure of the Clock
func (c *Clock) GetTrackState() *State {
	return &State{
		Min:     c.timeTracks[0],
		FiveMin: c.timeTracks[1],
		Hour:    c.timeTracks[2],
		Main:    c.ballQueue,
	}
}

//IsOriginalConfig returns true or false indicating if the Clock is in the same state immediatly after creation.
func (c *Clock) IsOriginalConfig() bool {
	if len(c.ballQueue) != c.ballCount {
		return false
	}

	for i := 0; i < c.ballCount; i++ {
		if c.ballQueue[i] != i {
			return false
		}
	}

	return true
}

//Tick increments the Clock by one minute
func (c *Clock) Tick() {
	var b int
	b, c.ballQueue = c.ballQueue[0], c.ballQueue[1:]

	for i := 0; i < 3; i++ {
		cap := fiveAndHourTrackCap
		if i == 0 {
			cap = minTrackCap
		}

		if o := c.addOrOverflow(i, b, cap); !o {
			return
		}
	}

	c.ballQueue = append(c.ballQueue, b)
}

//TickFive causes the clock to tick 5 minutes
func (c *Clock) TickFive() {
	var b int

	reverseSlice(c.ballQueue[0:4])
	b, c.ballQueue = c.ballQueue[4], append(c.ballQueue[5:], c.ballQueue[0:4]...)

	for i := 1; i < 3; i++ {
		if o := c.addOrOverflow(i, b, fiveAndHourTrackCap); !o {
			return
		}
	}

	c.ballQueue = append(c.ballQueue, b)
}

func (c *Clock) addOrOverflow(t int, b int, cap int) bool {
	if len(c.timeTracks[t]) != cap {
		c.timeTracks[t] = append(c.timeTracks[t], b)
		return false
	}

	reverseSlice(c.timeTracks[t])
	c.ballQueue = append(c.ballQueue, c.timeTracks[t]...)

	c.timeTracks[t] = c.timeTracks[t][:0]

	return true
}

func reverseSlice(s []int) {
	l := len(s) - 1
	for i := l / 2; i >= 0; i-- {
		s[i], s[l-i] = s[l-i], s[i]
	}
}
