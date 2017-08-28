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
	BallQueue  []int
	ballCount  int
}

//NewClock returns a new instance of the clock type
func NewClock(ballCount int) (*Clock, error) {
	if ballCount < MinBalls || ballCount > MaxBalls {
		return nil, fmt.Errorf("Invalid ball count. %v provided, must be between %v and %v", ballCount, MinBalls, MaxBalls)
	}

	c := &Clock{
		BallQueue: make([]int, ballCount, ballCount+1152+253+11),
		timeTracks: [][]int{
			make([]int, 0, minTrackCap),
			make([]int, 0, fiveAndHourTrackCap),
			make([]int, 0, fiveAndHourTrackCap)},
		ballCount: ballCount,
	}

	for i := 0; i < ballCount; i++ {
		c.BallQueue[i] = i
	}

	return c, nil
}

//GetTrackState returns a State object depicting the internal ball structure of the Clock
func (c *Clock) GetTrackState() *State {
	return &State{
		Min:     c.timeTracks[0],
		FiveMin: c.timeTracks[1],
		Hour:    c.timeTracks[2],
		Main:    c.BallQueue,
	}
}

//IsOriginalConfig returns true or false indicating if the Clock is in the same state immediatly after creation.
func (c *Clock) IsOriginalConfig() bool {
	if len(c.BallQueue) != c.ballCount {
		return false
	}

	tmp := append(make([]int, 0, c.ballCount+1152+253+11), c.BallQueue...)
	c.BallQueue = tmp

	for i := 0; i < c.ballCount; i++ {
		if c.BallQueue[i] != i {
			return false
		}
	}

	return true
}

//Tick increments the Clock by one minute
func (c *Clock) Tick() {
	var b int
	b, c.BallQueue = c.BallQueue[0], c.BallQueue[1:]

	for i := 0; i < 3; i++ {
		cap := fiveAndHourTrackCap
		if i == 0 {
			cap = minTrackCap
		}

		if o := c.addOrOverflow(i, b, cap); !o {
			return
		}
	}

	c.BallQueue = append(c.BallQueue, b)
}

//TickFive causes the clock to tick 5 minutes
func (c *Clock) TickFive() {
	var b int

	reverseSlice(c.BallQueue[0:4])
	b, c.BallQueue = c.BallQueue[4], append(c.BallQueue[5:], c.BallQueue[0:4]...)

	for i := 1; i < 3; i++ {
		if o := c.addOrOverflow(i, b, fiveAndHourTrackCap); !o {
			return
		}
	}

	c.BallQueue = append(c.BallQueue, b)
}

func (c *Clock) addOrOverflow(t int, b int, cap int) bool {
	if len(c.timeTracks[t]) != cap {
		c.timeTracks[t] = append(c.timeTracks[t], b)
		return false
	}

	reverseSlice(c.timeTracks[t])
	c.BallQueue = append(c.BallQueue, c.timeTracks[t]...)

	c.timeTracks[t] = c.timeTracks[t][:0]

	return true
}

func reverseSlice(s []int) {
	l := len(s) - 1
	for i := l / 2; i >= 0; i-- {
		s[i], s[l-i] = s[l-i], s[i]
	}
}
