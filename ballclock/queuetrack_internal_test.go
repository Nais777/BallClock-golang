package ballclock

import (
	"testing"
)

func TestNewQueueTrack(t *testing.T) {
	q := newQueueTrack(127)

	if len(q.balls) != 127 {
		t.Errorf("newQueueTrack Failed! Expected q.currentPos == 127, actual %d", len(q.balls))
	}

	for i := 0; i < cap(q.balls); i++ {
		if q.balls[i] != i {
			t.Errorf("newQueueTrack Failed! Expected ball.id == %d, actual %d", i+1, q.balls[i])
		}
	}
}

func TestGetBall(t *testing.T) {
	q := newQueueTrack(5)

	b := q.getBall()

	if b != 0 {
		t.Errorf("getBall Failed! Expected ball.id == 1, actual %d", b)
	}

	if len(q.balls) != 4 {
		t.Errorf("getBall Failed! Expected q.currentPos == 4, actual %d", len(q.balls))
	}

	if cap(q.balls) != 4 {
		t.Errorf("getBall Failed! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}
}

func TestReturnBall(t *testing.T) {
	e := []int{1, 2, 3, 4, 0}
	q := newQueueTrack(5)
	b := q.getBall()
	q.returnBall(b)

	if len(q.balls) != 5 {
		t.Errorf("returnBall Failed. Expected len(q.balls) == 5, actual %d", len(q.balls))
	}

	for i := 0; i < len(q.balls); i++ {
		if e[i] != q.balls[i] {
			t.Errorf("returnBall Failed! Expected ball.id == %d, actual %d", e[i], q.balls[i])
		}
	}
}

func TestReturnBalls(t *testing.T) {
	e := []int{2, 3, 4, 0, 1}
	q := newQueueTrack(5)
	s := make([]int, 0, 2)
	s = append(s, q.getBall())
	s = append(s, q.getBall())

	if len(q.balls) != 3 {
		t.Errorf("returnBalls Failed! Expected q.currentPos == 3, actual %d", len(q.balls))
	}

	q.returnBalls(s)

	if len(q.balls) != 5 {
		t.Errorf("returnBalls Failed! Expected len(q.balls) == 5, actual %d", len(q.balls))
	}

	for i := 0; i < len(q.balls); i++ {
		if e[i] != q.balls[i] {
			t.Errorf("ReturnBalls Failed! Expected ball.id == %d, actual %d", e[i], q.balls[i])
		}
	}
}

func TestOriginalConfiguration(t *testing.T) {
	q := newQueueTrack(30)

	suc := q.isOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed! Expected true, actual %t", suc)
	}

	q.returnBall(q.getBall())
	suc = q.isOriginalConfig()
	if suc {
		t.Errorf("OriginalConfiguration Failed - 1! Expected false, actual %t", suc)
	}

	for i := 1; i < 30; i++ {
		q.returnBall(q.getBall())
	}

	suc = q.isOriginalConfig()
	if !suc {
		t.Errorf("OriginalConfiguration Failed - 2! Expected true, actual %t", suc)
	}
}
