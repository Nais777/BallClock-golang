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

	return CalculateBallCycle(c.BallQueue)
}

//CalculateBallCycle calculates the ball position after 24 hrs
func CalculateBallCycle(s []int) int {
	visited := make([]bool, len(s), len(s))
	mapping := make([]int, len(s), len(s))

	for k, v := range s {
		mapping[v] = k
	}

	days := 1

	for k, v := range visited {
		if v {
			continue
		}

		tmp := 1
		start := k
		for ; ; tmp++ {
			next := mapping[k]
			if next == start {
				break
			}

			visited[k] = true
			k = next
		}

		if days%tmp != 0 {
			days *= tmp
		}
	}

	return days
}

//RunForTickCount ticks the clock for the amount specified
func RunForTickCount(c *ballclock.Clock, tickCount int) {
	for i := 0; i < fiveMinutesPerDay; i++ {
		if tickCount < 5 {
			break
		}

		c.TickFive()
		tickCount -= 5
	}

	if days := tickCount / minutesPerDay; days > 0 {
		PositionAfterDays(c.BallQueue, days)

		tickCount %= minutesPerDay
	}

	for ; tickCount >= 5; tickCount -= 5 {
		c.TickFive()
	}

	for ; tickCount > 0; tickCount-- {
		c.Tick()
	}
}

func PositionAfterDays(q []int, days int) {
	visited := make([]int, len(q), len(q))
	mapping := make([]int, len(q), len(q))
	rotDays := make(map[int]int)

	for k, v := range q {
		mapping[v] = k
	}

	grp := 0
	for k, v := range visited {
		if v != 0 {
			continue
		}

		if mapping[k] == k {
			continue
		}

		grp++

		tmp := 0
		start := k
		for ; ; tmp++ {
			next := mapping[k]
			visited[k] = grp
			k = next

			if k == start {
				break
			}
		}

		rotDays[grp] = tmp + 1
	}

	for grp, fullRot := range rotDays {
		if fullRot == 0 {
			continue
		}

		var start int
		for k, v := range visited {
			if v == grp {
				start = k
				break
			}
		}

		for rot := days % fullRot; rot > 0; rot-- {
			priorTmp := q[start]
			orgin := start

			for {
				dest := mapping[orgin]
				tmp := q[dest]

				q[dest] = priorTmp

				priorTmp = tmp
				orgin = dest

				if orgin == start {
					break
				}
			}
		}
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
