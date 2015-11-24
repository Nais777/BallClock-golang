package clock

import (
	"github.com/Nais777/BallClock-golang/balltrack"
)

const MIN_TRACK_CAP uint8 = 4
const FIVE_AND_HOUR_TRACK_CAP uint8 = 11

const MIN_BALLS = 27
const MAX_BALLS = 127

type Clock struct {
	timeTracks []*balltrack.TimeTrack
	ballQueue *balltrack.QueueTrack
}

func New(ballCount int) (clock *Clock, err string) {
	if ballCount < MIN_BALLS || ballCount > MAX_BALLS {
		return nil, "INVALID BALL COUNT"
	}

	c := new(Clock)
	c.ballQueue = balltrack.NewQueueTrack(uint8(ballCount))
	c.timeTracks = make([]*balltrack.TimeTrack, 3, 3)
	c.timeTracks[0] = balltrack.NewTimeTrack(MIN_TRACK_CAP)
	c.timeTracks[1] = balltrack.NewTimeTrack(FIVE_AND_HOUR_TRACK_CAP)
	c.timeTracks[2] = balltrack.NewTimeTrack(FIVE_AND_HOUR_TRACK_CAP)

	return c, ""
}

func (c *Clock) getTrackState() *ClockState {
	s := new(ClockState)
	s.Min = c.timeTracks[0].GetContentIds()
	s.FiveMin = c.timeTracks[1].GetContentIds()
	s.Hour = c.timeTracks[2].GetContentIds()
	s.Main = c.ballQueue.GetContentIds()

	return s
}

func (c *Clock) Run(runMin int) (s int, state *ClockState) {
	if runMin == 0 {
		return 0, c.getTrackState()
	}

	i := 0

	for {
		returnBall := true
		b := c.ballQueue.GetBall()
		for j := 0; j < 3; j++ {
			overFlow := c.timeTracks[j].Increment(b)
			if overFlow == nil {
				returnBall = false
				break
			}

			c.ballQueue.ReturnBalls(overFlow)
		}

		if returnBall {
			c.ballQueue.ReturnBall(b)

			if c.ballQueue.IsOriginalConfig() {
				return i+1, c.getTrackState()
			}
		}

		if runMin >= 0 && runMin == i+1 {
			return i+1, c.getTrackState()			
		}

		i++
	}
}
