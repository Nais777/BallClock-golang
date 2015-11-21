package ball

import (
	"testing"
)

func TestBall(t *testing.T){
	for i := uint8(0) ; i < 10 ; i++ {
		b := Ball{i};
		if b.Id != i {
			t.Errorf("Test ball failed! %d expected, %d recieved", i, b.Id)
		}
	} 
}

func TestNewBall(t *testing.T){
	for i:= uint8(0) ; i < 10 ; i++ {
		b := New(i);
		if b.Id != i {
			t.Errorf("Test New() failed! %d expected, %d recieved", i, b.Id)
		}
	}
}
