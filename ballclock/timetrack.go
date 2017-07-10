package ballclock

//increment

//timeTrack keeps track of balls for the one minute, five minute, and hour tracks
type timeTrack struct {
	*ballTrack
}

//newTimeTrack returns a pointer to a new instance of timeTrack
func newTimeTrack(cap uint8) *timeTrack {
	return &timeTrack{
		ballTrack: newBallTrack(cap),
	}
}

//reverseBalls reverses the balls in the slice
func (t *timeTrack) reverseBalls() {
	for i := (t.currentLen - 1) / 2; i >= 0; i-- {
		opp := t.currentLen - 1 - i
		t.balls[i], t.balls[opp] = t.balls[opp], t.balls[i]
	}
}

//clearTimeTrack clears the track balls
func (t *timeTrack) clearTimeTrack() {
	t.balls = make([]*ball, cap(t.balls), cap(t.balls))
	t.currentLen = 0
}

//increment adds a ball to the track returning the overflow balls
//if the track is already full.
func (t *timeTrack) increment(b *ball) []*ball {
	suc := t.addBall(b)
	if !suc {
		t.reverseBalls()
		ret := t.balls
		t.clearTimeTrack()
		return ret
	}

	return nil
}
