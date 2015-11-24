package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
)

type ballTrack struct {
	balls []*ball.Ball
}

func (t *ballTrack) isFull() bool {
	return len(t.balls) == cap(t.balls)
}

func (t *ballTrack) addBall(b *ball.Ball) bool {
	f := t.isFull()
	if !f {
		t.balls = append(t.balls, b)
	}

	return !f
}

func (t *ballTrack) GetContentIds() []uint8 {
	b := make([]uint8, len(t.balls), cap(t.balls))
	for i := range t.balls {
		b[i] = t.balls[i].Id
	}

	return b
}

func newBallTrack(cap uint8) *ballTrack {
	p := new(ballTrack)
	p.balls = make([]*ball.Ball, 0, cap)
	return p
}
