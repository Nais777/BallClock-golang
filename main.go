package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"encoding/json"

	"errors"
	"time"

	"github.com/Nais777/BallClock-golang/ballclock"
)

var reader *bufio.Reader

const minutesPerDay = 1440
const fiveMinutesPerDay = minutesPerDay / 5

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	PrintMenu()

	for {
		args, err := ParseInput()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		c, err := ballclock.NewClock(args.BallCount)
		if err != nil {
			fmt.Println("Error creating clock: " + err.Error())
			continue
		}

		if args.UseTickCount {
			RunForTickCount(c, args.TickCount)
			state, err := json.Marshal(c.GetTrackState())
			if err != nil {
				fmt.Printf("Unable to marshal clock state: %v\n", err.Error())
				continue
			}

			fmt.Printf("Clock state after %v ticks:\n", args.TickCount)
			fmt.Printf("%v\n", string(state))
		} else {
			i := CycleClock(c)
			fmt.Printf("Clock cycles after %v days.\n", i)
		}
	}
}

//PrintMenu prints the menu
func PrintMenu() {
	fmt.Println("Commands are expected to be in the format # #, with the second number being optional.")
	fmt.Printf("The first number dictactes the amount of balls in the clock and must be between %v and %v inclusive\n", ballclock.MinBalls, ballclock.MaxBalls)
	fmt.Println("In the case of a single number (ex: 30), the output will be the number of 24 hr periods until the clock returns to its original configuration.")
	fmt.Println("If two numbers are entered on the same line, the clock runs for the amount of minutes (ticks) specified by the second number.")
	fmt.Println("The output of the second input mode will be the ball configuration of the clock after the specified minutes have passed.")
}

//Args contains the run arguments
type Args struct {
	BallCount    int
	UseTickCount bool
	TickCount    int
}

//ParseInput parses the user input and returns the amount of balls and the amount of cycles if specified
func ParseInput() (*Args, error) {
	args := &Args{UseTickCount: false}
	text, _ := reader.ReadString('\n')
	s := strings.Split(strings.TrimSuffix(text, "\n"), " ")

	if len(s) == 0 || len(s) > 2 {
		return nil, fmt.Errorf("Invalid number of arguments. %v provided, 1 or 2 expected", len(s))
	}

	if strings.ToUpper(s[0]) == "EXIT" {
		os.Exit(0)
	}

	if strings.ToUpper(s[0]) == "BENCHMARK" {
		Benchmark()
		return nil, errors.New("Benchmark Complete")
	}

	if strings.ToUpper(s[0]) == "BENCHMARK5" {
		BenchmarkTickFive()
		return nil, errors.New("Benchmark Complete")
	}

	ballCount, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing ball count: %v", err.Error())
	}
	args.BallCount = int(ballCount)

	if len(s) == 2 {
		args.UseTickCount = true

		tmp, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing tick count: %v", err.Error())
		}
		args.TickCount = int(tmp)
	}

	return args, nil
}

//CycleClock causes the clock to run until it is back in it's original configuration and returns the
func CycleClock(c *ballclock.Clock) int {
	for i := 0; i < fiveMinutesPerDay; i++ {
		c.TickFive()
	}

	s := c.GetTrackState().Main

	return 1 + CalculateBallCycle(s)
}

//CalculateBallCycle calculates the ball position after 24 hrs
func CalculateBallCycle(s []int) int {
	tmp := make([]int, len(s), len(s))
	mapping := make([]int, len(s), len(s))

	copy(tmp, s)
	copy(mapping, s)

	var c int
	for c = 1; ; c++ {
		for k, v := range mapping {
			s[k] = tmp[v]
		}

		original := true
		for i, k := range s {
			if i == k {
				continue
			}

			original = false
			break
		}

		if original {
			break
		}

		copy(tmp, s)
	}

	return c
}

//RunForTickCount ticks the clock for the amount specified
func RunForTickCount(c *ballclock.Clock, tickCount int) {
	for i := 0; i < tickCount; i++ {
		c.Tick()
	}
}

func Benchmark() {
	for i := ballclock.MinBalls; i <= ballclock.MaxBalls; i++ {
		c, _ := ballclock.NewClock(i)

		start := time.Now()

		t := CycleClock(c)

		duration := time.Since(start)

		fmt.Printf("Ballclock with %v balls took %s; %v days\n", i, duration, t)
	}
}

func BenchmarkTickFive() {
	for i := ballclock.MinBalls; i <= ballclock.MaxBalls; i++ {
		c, _ := ballclock.NewClock(i)

		start := time.Now()

		for t := 0; t < fiveMinutesPerDay; t++ {
			c.TickFive()
		}

		duration := time.Since(start)

		fmt.Printf("Ballclock with %v balls took %s;\n", i, duration)
	}
}
