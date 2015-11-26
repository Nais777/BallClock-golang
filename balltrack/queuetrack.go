package balltrack

import (
		"github.com/Nais777/BallClock-golang/ball"
)

type QueueTrack struct {
	*ballTrack
}

func (q *QueueTrack) GetBall() *ball.Ball {
	b := q.balls[0]
	
	tmp := make([]*ball.Ball, len(q.balls), cap(q.balls))
	copy(tmp, q.balls[1:])
	q.balls = tmp

	q.currentPos -= 1
	
	return b
}

func (q *QueueTrack) ReturnBall(b *ball.Ball) {
	q.addBall(b)
}

func (q *QueueTrack) ReturnBalls(b []*ball.Ball) {
	for i := range b {
		q.addBall(b[i])
	}
}

func (q *QueueTrack) IsOriginalConfig() bool {
	if !q.isFull() {
		return false
	}
	
	for i := uint8(0); i < uint8(cap(q.balls)); i++ {
		if q.balls[i].Id != i+1 {
			return false
		}
	}
	
	return true
}

func NewQueueTrack(cap uint8) *QueueTrack {
	q := new(QueueTrack)
	q.ballTrack = newBallTrack(cap)
	
	for i := uint8(0); i < cap; i++ {
		q.addBall(ball.New(i+1))
	}
	
	return q
}
