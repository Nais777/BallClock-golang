package balltrack

import (
        "testing"
	"github.com/Nais777/BallClock-golang/ball"
)

func TestNewTimeTrack(t *testing.T){
	tr := NewTimeTrack(5)
	
	if tr.balls == nil {
		t.Errorf("NewTimeTrack Failed! Expected tr.balls != nil")
	}
}

func TestReverseBalls(t *testing.T){
	c := struct {
		in, out []uint8
	}{[]uint8{0,1,2,3,4}, []uint8{4,3,2,1,0}}
	
	tr := NewTimeTrack(uint8(len(c.in)))
	for b := range c.in {
		tr.addBall(ball.New(c.in[b]));
	}
		
	tr.ReverseBalls()

	for i := range tr.balls {
		if tr.balls[i].Id != c.out[i] {
			t.Errorf("TestGetReverseBalls Failed. Expected %d, actual %d.", c.out[i], tr.balls[i].Id)
		}
	}
}

func TestClearTimeTrack(t *testing.T){
	tr := NewTimeTrack(5)
	tr.balls = make([]*ball.Ball, 5, 5);
	tr.clearTimeTrack()
	
	for i := range tr.balls {
		if tr.balls[i] != nil {
			t.Errorf("Clear Time Track Failed! tr.balls[%d] != nil", i)
		}
	}

	if tr.currentPos != 0 {
		t.Errorf("clearTimeTrack failed! Expected tr.currentPos == 0, actual %d", tr.currentPos)
	}
}

func TestIncrement(t *testing.T){
	tr := NewTimeTrack(1)
	b := tr.Increment(ball.New(0))
	
	if b != nil {
		t.Errorf("Increment failed! Expected nil, got data")
	}
	
	b = tr.Increment(ball.New(1))
	
	if b == nil {
		t.Errorf("Increment failed! Expected data, got nil")
	}
}
