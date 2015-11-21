package ball

type Ball struct {
	Id uint8
}

func New(ballId uint8) *Ball {
	p := new(Ball)
	p.Id = ballId
	return p
}
