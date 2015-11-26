package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
)

type TimeTrack struct {
	*ballTrack
}

func (t *TimeTrack) ReverseBalls() {
	for i := (t.currentPos - 1)/2; i >= 0; i-- {
		opp := t.currentPos - 1 - i
		t.balls[i], t.balls[opp] = t.balls[opp], t.balls[i]
	}
} 

func (t *TimeTrack) clearTimeTrack(){
	t.balls = make([]*ball.Ball, cap(t.balls), cap(t.balls))
	t.currentPos = 0
}

func (t *TimeTrack) Increment(b *ball.Ball) []*ball.Ball{
	suc := t.addBall(b)
	if !suc {
		t.ReverseBalls()
		ret := t.balls
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
