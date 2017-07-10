package ballclock

import (
	"testing"
)

func TestBall(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		b := ball{i}
		if b.id != i {
			t.Errorf("Test ball failed! %d expected, %d recieved", i, b.id)
		}
	}
}

func TestNewBall(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		b := newBall(i)
		if b.id != i {
			t.Errorf("Test New() failed! %d expected, %d recieved", i, b.id)
		}
	}
}
