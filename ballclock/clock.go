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
	timeTracks []*timeTrack
	ballQueue  *queueTrack
}

//NewClock returns a new instance of the clock type
func NewClock(ballCount int) (c *Clock, err error) {
	if ballCount < MinBalls || ballCount > MaxBalls {
		return nil, fmt.Errorf("Invalid ball count. %v provided, must be between %v and %v", ballCount, MinBalls, MaxBalls)
	}

	return &Clock{
		ballQueue: newQueueTrack(ballCount),
		timeTracks: []*timeTrack{
			newTimeTrack(minTrackCap),
			newTimeTrack(fiveAndHourTrackCap),
			newTimeTrack(fiveAndHourTrackCap)},
	}, nil
}

//GetTrackState returns a State object depicting the internal ball structure of the Clock
func (c *Clock) GetTrackState() *State {
	return &State{
		Min:     c.timeTracks[0].getContentIds(),
		FiveMin: c.timeTracks[1].getContentIds(),
		Hour:    c.timeTracks[2].getContentIds(),
		Main:    c.ballQueue.getContentIds(),
	}
}

//IsOriginalConfig returns true or false indicating if the Clock is in the same state immediatly after creation.
func (c *Clock) IsOriginalConfig() bool {
	return c.ballQueue.isOriginalConfig()
}

//Tick increments the Clock by one minute
func (c *Clock) Tick() {
	b := c.ballQueue.getBall()

	for t := 0; t < 3; t++ {
		overFlow := c.timeTracks[t].increment(b)
		if overFlow == nil {
			return
		}

		c.ballQueue.returnBalls(overFlow)
	}

	c.ballQueue.returnBall(b)
}
