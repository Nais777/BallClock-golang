package ballclock

//ballTrack is a base structure for the time and queue tracks
type ballTrack struct {
	balls []*ball
}

//newBallTrack returns a pointer to a ball track
func newBallTrack(cap int) *ballTrack {
	return &ballTrack{
		balls: make([]*ball, 0, cap),
	}
}

//isFull returns true or false if the track is full to capacity
func (t *ballTrack) isFull() bool {
	return len(t.balls) == cap(t.balls)
}

//addBall checks if the track is full, adds the ball if it isnt and
//returns true or false if the ball was successfully added
func (t *ballTrack) addBall(b *ball) bool {
	f := t.isFull()
	if !f {
		t.balls = t.balls[0 : len(t.balls)+1]
		t.balls[len(t.balls)-1] = b
	}

	return !f
}

//getContentIds returns a slice containing the id's of the balls in the track
func (t *ballTrack) getContentIds() []int {
	b := make([]int, len(t.balls), cap(t.balls))
	for i := 0; i < len(t.balls); i++ {
		b[i] = t.balls[i].id
	}

	return b
}
