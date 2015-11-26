package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
	"testing"
)

func TestNewQueueTrack(t *testing.T){
	q := NewQueueTrack(127)
	
	if q.currentPos != 127 {
		t.Errorf("NewQueueTrack Failed! Expected q.currentPos == 127, actual %d", q.currentPos)
	}
	
	for i := uint8(0); i < uint8(cap(q.balls)) ; i++ {
		if q.balls[i].Id != i+1 {
			t.Errorf("NewQueueTrack Failed! Expected Ball.Id == %d, actual %d", i+1, q.balls[i].Id)
		}
	}
}

func TestGetBall(t *testing.T){
	q := NewQueueTrack(5)
	
	b := q.GetBall()
	
	if b.Id != 1 {
		t.Errorf("GetBall Failed! Expected Ball.Id == 1, actual %d", b.Id)
	}
	
	if q.currentPos != 4 {
		t.Errorf("GetBall Failed! Expected q.currentPos == 4, actual %d", q.currentPos)
	}
	
	if cap(q.balls) != 5 {
		t.Errorf("GetBall Failed! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}

}

func TestReturnBall(t *testing.T){
	e := []uint8{2,3,4,5,1}
	q := NewQueueTrack(5)
	b := q.GetBall()
	q.ReturnBall(b);
	
	if q.currentPos != 5 {
		t.Errorf("ReturnBall Failed. Expected q.currentPos == 5, actual %d", q.currentPos)
	}

	for i := 0; i < q.currentPos; i++ {
		if e[i] != q.balls[i].Id {
			t.Errorf("ReturnBall Failed! Expected ball.Ball.Id == %d, actual %d", e[i], q.balls[i].Id)
		}
	}
}

func TestReturnBalls(t *testing.T){
	e := []uint8{3,4,5,1,2}
	q := NewQueueTrack(5)
	s := make([]*ball.Ball, 0, 2)
	s = append(s, q.GetBall())
	s = append(s, q.GetBall())
	
	if q.currentPos != 3 {
		t.Errorf("ReturnBalls Failed! Expected q.currentPos == 3, actual %d", q.currentPos)
	}
	
	q.ReturnBalls(s)
	
	if q.currentPos != 5 {
		t.Errorf("ReturnBalls Failed! Expected q.currentPos == 5, actual %d", q.currentPos)
	}
	
	for i := 0; i < len(q.balls); i++ {
		if e[i] != q.balls[i].Id {
			t.Errorf("ReturnBalls Failed! Expected ball.Ball.Id == %d, actual %d", e[i], q.balls[i].Id)
		}
	}
}

func TestOriginalConfiguration(t *testing.T){
	q := NewQueueTrack(30)
	
	suc := q.IsOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed! Expected true, actual %t", suc)
	}
	
	q.ReturnBall(q.GetBall());
	suc = q.IsOriginalConfig()
	if suc {
		t.Errorf("OriginalConfiguration Failed - 1! Expected false, actual %t", suc)
	}
	
	for i := 1; i < 30; i++ {
        	q.ReturnBall(q.GetBall());
	}

	suc = q.IsOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed - 2! Expected true, actual %t", suc)
	}
}
