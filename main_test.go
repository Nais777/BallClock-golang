package main

import (
	"strconv"
	"testing"

	"github.com/Nais777/BallClock-golang/ballclock"
)

func BenchmarkCycleClock(b *testing.B) {
	for i := ballclock.MinBalls; i <= ballclock.MaxBalls; i++ {
		b.Run("BallCount="+strconv.Itoa(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c, _ := ballclock.NewClock(i)

				CycleClock(c)
			}
		})
	}
}
