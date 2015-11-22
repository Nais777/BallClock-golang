package balltrack

import (
        "testing"
	"github.com/Nais777/BallClock-golang/ball"
)

func TestTimeTrack(t *testing.T){
	tr := TimeTrack{}
	tr.ballTrack = newBallTrack(5)
	
	if cap(tr.balls) != 5 {
		t.Errorf("TimeTrack Failed! Expected cap(tr.balls) == 5, actual %d", cap(tr.balls))
	}
}

func TestNewTimeTrack(t *testing.T){
	tr := NewTimeTrack(5)
	
	if cap(tr.balls) != 5 {
		t.Errorf("NewTimeTrack Failed! Expected cap(tr.balls) == 5, actual %d", cap(tr.balls))
	}
}

func TestGetReverseBalls(t *testing.T){
	c := struct {
		in, out []uint8
	}{[]uint8{0,1,2,3,4}, []uint8{4,3,2,1,0}}
	
	tr := NewTimeTrack(uint8(len(c.in)))
	for b := range c.in {
		tr.addBall(ball.New(c.in[b]));
	}
		
	rb := tr.getReverseBalls()
	for i := range rb {
		if rb[i].Id != c.out[i] {
			t.Errorf("TestGetReverseBalls Failed. Expected %d, actual %d.", c.out[i], rb[i].Id)
		}
	}
}

func TestClearTimeTrack(t *testing.T){
	tr := NewTimeTrack(5)
	tr.balls = make([]*ball.Ball, 5, 5);
	tr.clearTimeTrack()
	
	if len(tr.balls) != 0 {
		t.Errorf("clearTimeTrack failed! Expected len(tr.balls) == 0, actual %d", len(tr.balls))
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
