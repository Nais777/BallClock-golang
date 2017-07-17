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

func TestTickFive(t *testing.T) {
	c, _ := NewClock(27)

	c.TickFive()

	if len(c.timeTracks[1].balls) != 1 {
		t.Errorf("Five minute track not correct after tick 5: %v", c.timeTracks[1].balls)
	}

	if c.timeTracks[1].balls[0] != 4 {
		t.Errorf("Wrong ball in 5 minute track: %v", c.timeTracks[1].balls[0])
	}

	if len(c.ballQueue.balls) != 26 {
		t.Errorf("Wrong length for ballQueue: %v", c.ballQueue.balls)
	}
}
