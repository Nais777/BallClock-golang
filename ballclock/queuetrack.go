package ballclock

//queueTrack holds the available balls
type queueTrack struct {
	*ballTrack

	maxCapacity int
}

//newQueueTrack returns a new instance of queueTrack
func newQueueTrack(cap int) *queueTrack {
	q := &queueTrack{
		ballTrack:   newBallTrack(cap),
		maxCapacity: cap,
	}

	for i := 0; i < cap; i++ {
		q.addBall(ball(i + 1))
	}

	return q
}

//getBall gets a ball from the queuetrack
func (q *queueTrack) getBall() ball {
	var b ball
	b, q.balls = q.balls[0], q.balls[1:]

	return b
}

//returnBalls returns a ball to the queue
func (q *queueTrack) returnBall(b ball) {
	q.balls = append(q.balls, b)
}

//returnBalls returns multiple balls to the queue
func (q *queueTrack) returnBalls(b []ball) {
	q.balls = append(q.balls, b...)
}

//isOriginalConfig returns true or false indicating if the track is full
//and all balls are in the proper order.
func (q *queueTrack) isOriginalConfig() bool {
	if len(q.balls) != q.maxCapacity {
		return false
	}

	for i := 0; i < q.maxCapacity; i++ {
		if int(q.balls[i]) != i+1 {
			return false
		}
	}

	return true
}
