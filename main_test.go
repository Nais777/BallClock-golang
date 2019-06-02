package main

import (
	"strconv"
	"testing"

	"github.com/lamebear/BallClock-golang/ballclock"
)

func BenchmarkCycleClock(b *testing.B) {
	for i := ballclock.MinBalls; i <= ballclock.MaxBalls; i++ {
		j := i
		b.Run("BallCount="+strconv.Itoa(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c, _ := ballclock.NewClock(j)

				CycleClock(c)
			}
		})
	}
}

func BenchmarkCycle24(b *testing.B) {
	for i := ballclock.MinBalls; i <= ballclock.MaxBalls; i++ {
		j := i
		b.Run("BallCount="+strconv.Itoa(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				c, _ := ballclock.NewClock(j)

				for k := 0; k < fiveMinutesPerDay; k++ {
					c.TickFive()
				}
			}
		})
	}
}

func BenchmarkRunForTickCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c, _ := ballclock.NewClock(ballclock.MaxBalls)

		RunForTickCount(c, 123456789)
	}
}
