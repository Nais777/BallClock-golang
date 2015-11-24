package balltrack

import (
	"github.com/Nais777/BallClock-golang/ball"
	"testing"
)

func TestQueueTrack(t *testing.T){
	q := QueueTrack{}
	q.ballTrack = newBallTrack(127)
		
	if cap(q.balls) != 127 {
		t.Errorf("QueueTrack Failed! Expected cap(q.balls) == 127, actual %d", cap(q.balls))
	}
}

func TestNewQueueTrack(t *testing.T){
	q := NewQueueTrack(127)
	
	if cap(q.balls) != 127 {
		t.Errorf("NewQueueTrack Failed! Expected cap(q.balls) == 127, actual %d", cap(q.balls))
	}
	
	if len(q.balls) != 127 {
		t.Errorf("NewQueueTrack Failed! Expected len(q.balls) == 127, actual %d", len(q.balls))
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
	
	if len(q.balls) != 4 {
		t.Errorf("GetBall Failed! Expected len(q.balls) == 4, actual %d", len(q.balls))
	}
	
	if cap(q.balls) != 5 {
		t.Errorf("GetBall Failed! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}

	b = q.GetBall()

        if b.Id != 2 {
                t.Errorf("GetBall Failed! Expected Ball.Id == 2, actual %d", b.Id)
        }

        if len(q.balls) != 3 {
                t.Errorf("GetBall Failed! Expected len(q.balls) == 3, actual %d", len(q.balls))
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
	
	for i := 0; i < len(q.balls); i++ {
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
	
	if len(q.balls) != 3 {
		t.Errorf("ReturnBalls Failed! Expected len(q.balls) == 3, actual %d", len(q.balls))
	}
	
	if cap(q.balls) != 5 {
		t.Errorf("ReturnBalls Failed! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}
	
	q.ReturnBalls(s)
	
	if len(q.balls) != 5 {
		t.Errorf("ReturnBalls Failed! Expected len(q.balls) == 5, actual %d", len(q.balls))
	}
	
	if cap(q.balls) != 5 {
		t.Errorf("ReturnBalls Failed - 2! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}
	
	for i := 0; i < len(q.balls); i++ {
		if e[i] != q.balls[i].Id {
			t.Errorf("ReturnBalls Failed! Expected ball.Ball.Id == %d, actual %d", e[i], q.balls[i].Id)
		}
	}
}

func TestOriginalConfiguration(t *testing.T){
	q := NewQueueTrack(5)
	s := make([]*ball.Ball, 5, 5);
	copy(s, q.balls);
	
	suc := q.IsOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed! Expected true, actual %t", suc)
	}
	
	q.ReturnBall(q.GetBall());
	suc = q.IsOriginalConfig()
	if suc {
		t.Errorf("OriginalConfiguration Failed - 1! Expected false, actual %t", suc)
	}
	
	q.balls = s;
	suc = q.IsOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed - 2! Expected true, actual %t", suc)
	}
}
