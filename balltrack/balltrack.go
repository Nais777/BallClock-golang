package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
)

type ballTrack struct {
	balls []*ball.Ball
	currentPos int
}

func (t *ballTrack) isFull() bool {
	return t.currentPos == cap(t.balls)
}

func (t *ballTrack) addBall(b *ball.Ball) bool {
	f := t.isFull()
	if !f {
		t.balls[t.currentPos] = b
		t.currentPos += 1	 
	}

	return !f
}

func (t *ballTrack) GetContentIds() []int {
	b := make([]int, t.currentPos, cap(t.balls))
	for i := 0; i < t.currentPos; i++ {
		b[i] = int(t.balls[i].Id)
	}

	return b
}

func newBallTrack(cap uint8) *ballTrack {
	t := new(ballTrack)
	t.balls = make([]*ball.Ball, cap, cap)
	t.currentPos = 0
	return t
}
