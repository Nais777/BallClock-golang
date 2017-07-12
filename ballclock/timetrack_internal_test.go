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
		in, out []ball
	}{[]ball{0, 1, 2, 3, 4}, []ball{4, 3, 2, 1, 0}}

	tr := newTimeTrack(len(c.in))
	for b := range c.in {
		tr.addBall(c.in[b])
	}

	tr.reverseBalls()

	for i := range tr.balls {
		if tr.balls[i] != c.out[i] {
			t.Errorf("TestGetReverseBalls Failed. Expected %d, actual %d.", c.out[i], tr.balls[i])
		}
	}
}

func TestClearTimeTrack(t *testing.T) {
	tr := newTimeTrack(5)
	tr.balls = make([]ball, 5, 5)
	tr.clearTimeTrack()

	if len(tr.balls) != 0 {
		t.Errorf("clearTimeTrack failed! Expected tr.currentPos == 0, actual %d", len(tr.balls))
	}
}

func TestIncrement(t *testing.T) {
	tr := newTimeTrack(1)
	b := tr.increment(0)

	if b != nil {
		t.Errorf("Increment failed! Expected nil, got data")
	}

	b = tr.increment(1)

	if b == nil {
		t.Errorf("Increment failed! Expected data, got nil")
	}
}
