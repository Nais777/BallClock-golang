package ballclock

//ballTrack is a base structure for the time and queue tracks
type ballTrack struct {
	balls      []*ball
	currentLen int
}

//newBallTrack returns a pointer to a ball track
func newBallTrack(cap uint8) *ballTrack {
	return &ballTrack{
		balls:      make([]*ball, cap, cap),
		currentLen: 0,
	}
}

//isFull returns true or false if the track is full to capacity
func (t *ballTrack) isFull() bool {
	return t.currentLen == cap(t.balls)
}

//addBall checks if the track is full, adds the ball if it isnt and
//returns true or false if the ball was successfully added
func (t *ballTrack) addBall(b *ball) bool {
	f := t.isFull()
	if !f {
		t.balls[t.currentLen] = b
		t.currentLen++
	}

	return !f
}

//getContentIds returns a slice containing the id's of the balls in the track
func (t *ballTrack) getContentIds() []int {
	b := make([]int, t.currentLen, cap(t.balls))
	for i := 0; i < t.currentLen; i++ {
		b[i] = int(t.balls[i].id)
	}

	return b
}
