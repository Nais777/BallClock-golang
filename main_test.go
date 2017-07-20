package main

import (
	"strconv"
	"testing"

	"github.com/Nais777/BallClock-golang/ballclock"
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

func BenchmarkCalculateBallCycle123(b *testing.B) {
	m := []int{8, 62, 42, 87, 108, 35, 17, 6, 22, 75, 116, 112, 39, 119, 52, 60, 30, 88, 56, 36, 38, 26, 51, 31, 55, 120, 33, 99, 111, 24, 45, 21, 23, 34, 43, 41, 67, 65, 66, 85, 82, 89, 9, 25, 109, 47, 40, 0, 83, 46, 73, 13, 12, 63, 15, 90, 121, 2, 69, 53, 28, 72, 97, 3, 4, 94, 106, 61, 96, 18, 80, 74, 44, 84, 107, 98, 93, 103, 5, 91, 32, 76, 20, 68, 81, 95, 29, 27, 86, 104, 7, 64, 113, 78, 105, 58, 118, 117, 50, 70, 10, 101, 110, 19, 1, 115, 102, 71, 79, 57, 77, 122, 48, 114, 54, 37, 59, 49, 100, 11, 14, 92, 16}

	for n := 0; n < b.N; n++ {
		tmp := append([]int{}, m...)
		CalculateBallCycle(tmp)
	}
}

func BenchmarkCalculateBallCycle27(b *testing.B) {
	m := []int{12, 23, 9, 25, 26, 3, 13, 10, 15, 18, 20, 11, 17, 4, 5, 1, 16, 24, 14, 2, 19, 21, 22, 7, 8, 6, 0}

	for n := 0; n < b.N; n++ {
		tmp := append([]int{}, m...)
		CalculateBallCycle(tmp)
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
