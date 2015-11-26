package balltrack

import (
	"testing"
	"github.com/Nais777/BallClock-golang/ball"
)

func TestNewBallTrack(t *testing.T){
	bt := newBallTrack(5)
	if len(bt.balls) != 5 {
                t.Errorf("newBallTrack Test Failed! Expected len(bt.balls) == 5, actual %d", len(bt.balls));
        }
        if bt.balls[0] != nil {
                t.Errorf("ballTrack Test Failed! Expected bt.balls[0] == nil")
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

	bt.balls[0] = ball.New(0)
	bt.currentPos += 1
	suc = bt.isFull()
	if !suc {
		t.Errorf("isFull() failed! Expected true, actual %t", suc)
	}
}

func TestAddBall(t *testing.T){
	bt := newBallTrack(1)
	b := ball.New(1)
	suc := bt.addBall(b)
	
	if !suc  {
		t.Errorf("Add Ball failed! Expected true, actual %t", suc)
	}

	b = ball.New(2)
	suc = bt.addBall(b)
	
	if suc {
		t.Errorf("Add Ball failed! Expected false, actual %t", suc)
	}
}

func TestGetContents(t *testing.T){
	bt := newBallTrack(5)
	e := []int{0,1,2,3}
	for i := range e {
		bt.addBall(ball.New(uint8(e[i])))
	}

	c := bt.GetContentIds()
	for i := range e {
		if c[i] != e[i] {
			t.Errorf("Get Contents failed, expected %d, actual %d", e[i], c[i])
		}
	}
}


