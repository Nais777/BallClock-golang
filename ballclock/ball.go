package ballclock

type ball struct {
	id int
}

func newBall(ballID int) *ball {
	return &ball{
		id: ballID,
	}
}
