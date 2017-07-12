package ballclock

import (
	"testing"
)

func TestNewTimeTrack(t *testing.T) {
	tr := newTimeTrack(5)

	if tr.balls == nil {
		t.Errorf("NewTimeTrack Failed! Expected tr.balls != nil")
	}
}

func TestReverseBalls(t *testing.T) {
	c := struct {
		in, out []int
	}{[]int{0, 1, 2, 3, 4}, []int{4, 3, 2, 1, 0}}

	tr := newTimeTrack(len(c.in))
	for b := range c.in {
		tr.addBall(newBall(c.in[b]))
	}

	tr.reverseBalls()

	for i := range tr.balls {
		if tr.balls[i].id != c.out[i] {
			t.Errorf("TestGetReverseBalls Failed. Expected %d, actual %d.", c.out[i], tr.balls[i].id)
		}
	}
}

func TestClearTimeTrack(t *testing.T) {
	tr := newTimeTrack(5)
	tr.balls = make([]*ball, 5, 5)
	tr.clearTimeTrack()

	for i := range tr.balls {
		if tr.balls[i] != nil {
			t.Errorf("Clear Time Track Failed! tr.balls[%d] != nil", i)
		}
	}

	if tr.currentLen != 0 {
		t.Errorf("clearTimeTrack failed! Expected tr.currentPos == 0, actual %d", tr.currentLen)
	}
}

func TestIncrement(t *testing.T) {
	tr := newTimeTrack(1)
	b := tr.increment(newBall(0))

	if b != nil {
		t.Errorf("Increment failed! Expected nil, got data")
	}

	b = tr.increment(newBall(1))

	if b == nil {
		t.Errorf("Increment failed! Expected data, got nil")
	}
}
