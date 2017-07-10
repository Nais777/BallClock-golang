package ballclock

type ball struct {
	id uint8
}

func newBall(ballID uint8) *ball {
	return &ball{
		id: ballID,
	}
}
