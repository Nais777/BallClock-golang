package ballclock

//timeTrack keeps track of balls for the one minute, five minute, and hour tracks
type timeTrack struct {
	*ballTrack
}

//newTimeTrack returns a pointer to a new instance of timeTrack
func newTimeTrack(cap int) *timeTrack {
	return &timeTrack{
		ballTrack: newBallTrack(cap),
	}
}

//reverseBalls reverses the balls in the slice
func (t *timeTrack) reverseBalls() {
	l := len(t.balls)
	for i := (l - 1) / 2; i >= 0; i-- {
		opp := l - 1 - i
		t.balls[i], t.balls[opp] = t.balls[opp], t.balls[i]
	}
}

//clearTimeTrack clears the track balls
func (t *timeTrack) clearTimeTrack() {
	t.balls = t.balls[:0]
}

//increment adds a ball to the track returning the overflow balls
//if the track is already full.
func (t *timeTrack) increment(b int) []int {
	suc := t.addBall(b)
	if !suc {
		t.reverseBalls()
		ret := t.balls
		t.clearTimeTrack()
		return ret
	}

	return nil
}
