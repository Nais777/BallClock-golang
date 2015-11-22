package balltrack

import (
		"github.com/Nais777/BallClock-golang/ball"
)

type QueueTrack struct {
	*ballTrack
}

func (q *QueueTrack) GetBall() *ball.Ball {
	b := q.balls[0]
	
	tmp := make([]*ball.Ball, len(q.balls) -1, cap(q.balls))
	copy(tmp, q.balls[1:])
	q.balls = tmp
	
	return b
}

func (q *QueueTrack) ReturnBall(b *ball.Ball) {
	q.addBall(b)
}

func (q *QueueTrack) ReturnBalls(b []*ball.Ball) {
	q.balls = append(q.balls, b...)
}

func (q *QueueTrack) IsOriginalConfig() bool {
	if !q.isFull() {
		return false
	}
	
	for i := uint8(0); i < uint8(cap(q.balls)); i++ {
		if q.balls[i].Id != i {
			return false
		}
	}
	
	return true
}

func NewQueueTrack(cap uint8) *QueueTrack {
	q := new(QueueTrack)
	q.ballTrack = newBallTrack(cap)
	
	for i := uint8(0); i < cap; i++ {
		q.balls = append(q.balls, ball.New(i))
	}
	
	return q
}
