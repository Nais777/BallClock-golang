package ballclock

//State represents the internal ball structure of the clock.
type State struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}
