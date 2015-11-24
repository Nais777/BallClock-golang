package balltrack

import (
	"testing"
	"github.com/Nais777/BallClock-golang/ball"
)

func TestBalltrack(t *testing.T){
	bt := ballTrack{make([]*ball.Ball, 0, 5)}
	if len(bt.balls) != 0 {
		t.Errorf("ballTrack Test Failed! Expected len(bt.balls) == 0, actual %d", len(bt.balls));
	}
	if cap(bt.balls) != 5 {
		t.Errorf("ballTrack Test Failed! Expected cap(bt.balls) == 5, actual %d", cap(bt.balls));
	}
}

func TestNewBallTrack(t *testing.T){
	bt := newBallTrack(5)
	if len(bt.balls) != 0 {
                t.Errorf("newBallTrack Test Failed! Expected len(bt.balls) == 0, actual %d", len(bt.balls));
        }
        if cap(bt.balls) != 5 {
                t.Errorf("newBallTrack Test Failed! Expected cap(bt.balls) == 5, actual %d", cap(bt.balls));
        }  
}

func TestIsFull(t *testing.T){
	bt := newBallTrack(1)
	suc := bt.isFull()
	if suc {
		t.Errorf("isFull() failed! Expected false, actual %t", suc);
	}

	bt.balls = append(bt.balls, ball.New(0))
	suc = bt.isFull()
	if !suc {
		t.Errorf("isFull() failed! Expected false, actual %t", suc)
	}
}

func TestAddBall(t *testing.T){
	bt := newBallTrack(1)
	b := ball.New(1)
	suc := bt.addBall(b)
	
	if !suc  {
		t.Errorf("Add Ball failed! Expected true, actual %t, len(bt.balls) == %d, cap(bt.balls) == %d", suc, len(bt.balls), cap(bt.balls))
	}

	b = ball.New(2)
	suc = bt.addBall(b)
	
	if suc {
		t.Errorf("Add Ball failed! Expected false, actual %t, len(bt.balls) == %d, cap(bt.balls) == %d", suc, len(bt.balls), cap(bt.balls))
	}
}

func TestGetContents(t *testing.T){
	bt := newBallTrack(5)
	e := []uint8{0,1,2,3}
	for i := range e {
		bt.addBall(ball.New(e[i]))
	}

	c := bt.GetContentIds()
	for i := range e {
		if c[i] != e[i] {
			t.Errorf("Get Contents failed, expected %d, actual %d", e[i], c[i])
		}
	}
}


