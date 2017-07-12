package ballclock

import (
	"testing"
)

func TestNewBallTrack(t *testing.T) {
	bt := newBallTrack(5)
	if len(bt.balls) != 0 {
		t.Errorf("newBallTrack Test Failed! Expected len(bt.balls) == 0, actual %d", len(bt.balls))
	}
	if cap(bt.balls) != 5 {
		t.Errorf("newBallTrack Test Failed! Expected cap(bt.balls) == 5, actual %d", cap(bt.balls))
	}
}

func TestIsFull(t *testing.T) {
	bt := newBallTrack(1)
	suc := bt.isFull()
	if suc {
		t.Errorf("isFull() failed! Expected false, actual %t", suc)
	}

	bt.balls = bt.balls[0 : len(bt.balls)+1]
	bt.balls[0] = 5
	suc = bt.isFull()
	if !suc {
		t.Errorf("isFull() failed! Expected true, actual %t", suc)
	}
}

func TestAddBall(t *testing.T) {
	bt := newBallTrack(1)
	suc := bt.addBall(1)

	if !suc {
		t.Errorf("Add Ball failed! Expected true, actual %t", suc)
	}

	suc = bt.addBall(2)

	if suc {
		t.Errorf("Add Ball failed! Expected false, actual %t", suc)
	}
}

func TestGetContents(t *testing.T) {
	bt := newBallTrack(5)
	e := []ball{0, 1, 2, 3}
	for i := range e {
		bt.addBall(e[i])
	}

	c := bt.getContentIds()
	for i := range e {
		if c[i] != int(e[i]) {
			t.Errorf("Get Contents failed, expected %d, actual %d", e[i], c[i])
		}
	}
}
