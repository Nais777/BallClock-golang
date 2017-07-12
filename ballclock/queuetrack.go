package ballclock

//isOriginalConfig
//getBall
//returnBall
//returnBalls

//queueTrack holds the available balls
type queueTrack struct {
	*ballTrack
}

//newQueueTrack returns a new instance of queueTrack
func newQueueTrack(cap uint8) *queueTrack {
	q := &queueTrack{
		ballTrack: newBallTrack(cap),
	}

	for i := uint8(0); i < cap; i++ {
		q.addBall(newBall(i + 1))
	}

	return q
}

//getBall gets a ball from the queuetrack
func (q *queueTrack) getBall() *ball {
	var b *ball
	b, q.balls = q.balls[0], q.balls[1:]

	q.currentLen--

	return b
}

//returnBalls returns a ball to the queue
func (q *queueTrack) returnBall(b *ball) {
	q.balls = append(q.balls, b)
	q.currentLen = len(q.balls)
}

//returnBalls returns multiple balls to the queue
func (q *queueTrack) returnBalls(b []*ball) {
	q.balls = append(q.balls, b...)
	q.currentLen = len(q.balls)
}

//isOriginalConfig returns true or false indicating if the track is full
//and all balls are in the proper order.
func (q *queueTrack) isOriginalConfig() bool {
	if !q.isFull() {
		return false
	}

	for i := uint8(0); i < uint8(q.capacity); i++ {
		if q.balls[i].id != i+1 {
			return false
		}
	}

	return true
}
