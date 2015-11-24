package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
)

type TimeTrack struct {
	*ballTrack
}

func (t *TimeTrack) getReverseBalls() []*ball.Ball{
	b := make([]*ball.Ball, len(t.balls), cap(t.balls))
	for i := len(t.balls) - 1; i >= 0; i-- {
		b[len(t.balls) - 1 - i] = t.balls[i]
	}

	return b
} 

func (t *TimeTrack) clearTimeTrack(){
	t.balls = make([]*ball.Ball, 0, cap(t.balls))
}

func (t *TimeTrack) Increment(b *ball.Ball) []*ball.Ball{
	suc := t.addBall(b)
	if !suc {
		ret := t.getReverseBalls()
		t.clearTimeTrack()
		return ret
	}

	return nil
}

func NewTimeTrack(cap uint8) *TimeTrack{
	t := new(TimeTrack)
	t.ballTrack = newBallTrack(cap)
	return t
}
