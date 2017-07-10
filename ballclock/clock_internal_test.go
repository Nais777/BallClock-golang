package ballclock

import (
	"testing"
)

func TestNewClock(t *testing.T) {
	c, err := NewClock(27)

	if c == nil {
		t.Errorf("New Clock Failed. C should have been initiated. ERROR: %v", err.Error())
	}

	if c.timeTracks == nil {
		t.Errorf("New Clock Failed. TimeTracks not initiated")
	}

	if len(c.timeTracks) != 3 {
		t.Errorf("New Clcok Failed. Not all time tracks created")
	}

	if c.ballQueue == nil {
		t.Errorf("New Clock Failed. BallQueue not initiated")
	}

	c, err = NewClock(20)

	if c != nil {
		t.Errorf("New Clock Failed. Should have thrown INVALID BALL AMOUNT")
	}
}
