package ballclock

import (
	"testing"
)

func TestNewQueueTrack(t *testing.T) {
	q := newQueueTrack(127)

	if q.currentLen != 127 {
		t.Errorf("newQueueTrack Failed! Expected q.currentPos == 127, actual %d", q.currentLen)
	}

	for i := 0; i < cap(q.balls); i++ {
		if q.balls[i].id != i+1 {
			t.Errorf("newQueueTrack Failed! Expected ball.id == %d, actual %d", i+1, q.balls[i].id)
		}
	}
}

func TestGetBall(t *testing.T) {
	q := newQueueTrack(5)

	b := q.getBall()

	if b.id != 1 {
		t.Errorf("getBall Failed! Expected ball.id == 1, actual %d", b.id)
	}

	if q.currentLen != 4 {
		t.Errorf("getBall Failed! Expected q.currentPos == 4, actual %d", q.currentLen)
	}

	if cap(q.balls) != 4 {
		t.Errorf("getBall Failed! Expected cap(q.balls) == 5, actual %d", cap(q.balls))
	}
}

func TestReturnBall(t *testing.T) {
	e := []int{2, 3, 4, 5, 1}
	q := newQueueTrack(5)
	b := q.getBall()
	q.returnBall(b)

	if q.currentLen != 5 {
		t.Errorf("returnBall Failed. Expected q.currentLen == 5, actual %d", q.currentLen)
	}

	for i := 0; i < q.currentLen; i++ {
		if e[i] != q.balls[i].id {
			t.Errorf("returnBall Failed! Expected ball.id == %d, actual %d", e[i], q.balls[i].id)
		}
	}
}

func TestReturnBalls(t *testing.T) {
	e := []int{3, 4, 5, 1, 2}
	q := newQueueTrack(5)
	s := make([]*ball, 0, 2)
	s = append(s, q.getBall())
	s = append(s, q.getBall())

	if q.currentLen != 3 {
		t.Errorf("returnBalls Failed! Expected q.currentPos == 3, actual %d", q.currentLen)
	}

	q.returnBalls(s)

	if len(q.balls) != 5 {
		t.Errorf("returnBalls Failed! Expected q.currentLen == 5, actual %d", q.currentLen)
	}

	for i := 0; i < len(q.balls); i++ {
		if e[i] != q.balls[i].id {
			t.Errorf("ReturnBalls Failed! Expected ball.id == %d, actual %d", e[i], q.balls[i].id)
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
