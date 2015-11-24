package clock

import (
	"testing"
)

func TestClock (t *testing.T){
	c, err := New(27)

	if c == nil {
		t.Errorf("New Clock Failed. C should have been initiated. ERROR: " + err)
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

        c, err = New(20)

        if c != nil {
                t.Errorf("New Clock Failed. Should have thrown INVALID BALL AMOUNT")
        }
}

func TestRun(t *testing.T){
	c, _ := New(30);

	i, _ := c.Run(-1);

	if i/1440 != 15 {
		t.Errorf("Clock Run 1 Failed, %d", i)
	}

	c, _ = New(45)
	i, _ = c.Run(-1);

        if i/1440 != 378 {
                t.Errorf("Clock Run 2 Failed, %d", i)
        }
	
}

func TestRunLimit(t *testing.T){
        c, _ := New(30)
        _, s := c.Run(325)

        e := new(ClockState)
        e.Min = []int{}
        e.FiveMin = []int{22,13,25,3,7}
        e.Hour = []int{6,12,17,4,15}
        e.Main = []int{11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9}

        if len(s.Min) != len(e.Min) {
		t.Errorf("Run Limit Failed. len(s.Min) == %d, expected %d", len(s.Min), len(e.Min))
	}
		
	if len(s.FiveMin) != len(e.FiveMin) {
		t.Errorf("Run Limit Failed. len(s.FiveMin) == %d, expected %d", len(s.FiveMin), len(e.FiveMin))
	}
		
	for i := range e.FiveMin {
		if s.FiveMin[i] != e.FiveMin[i] {
			t.Errorf("Run Limit Failed, s.FiveMin[%d] = %d, expected %d", i, s.FiveMin[i], e.FiveMin[i])
		}
	}
		
	if len(s.Hour) != len(e.Hour) {
		t.Errorf("Run Limit Failed. len(s.Hour) == %d, expected %d", len(s.Hour), len(e.Hour))
	}
		
	for i := range e.Hour {
		if s.Hour[i] != e.Hour[i] {
			t.Errorf("Run Limit Failed, s.Hour[%d] = %d, expected %d", i, s.Hour[i], e.Hour[i])
		}
	}
		
	if len(s.Main) != len(e.Main) {
		t.Errorf("Run Limit Failed. len(s.Main) == %d, expected %d", len(s.Main), len(e.Main))
	}
		
	for i := range e.Main {
		if s.Main[i] != e.Main[i] {
			t.Errorf("Run Limit Failed, s.Main[%d] = %d, expected %d", i, s.Main[i], e.Main[i])
		}
	}
}
