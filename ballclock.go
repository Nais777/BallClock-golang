package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/Nais777/BallClock-golang/clock"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Commands are expected to be in the format # #, with the second number being optional. In the case of a single number (ex: 30), the clock will be configured witht ehe specified amount of balls and the output will be the number of 24 hr periods untill the clock returns to its original configuration. If two numbers are entered on the same line, the clock is configured to have the number of balls specified by the first argument, and runs for the amount of minutes specified by the second. The output of the second input mode will be the ball configuration of the clock after the specified minutes have passed. ")

	for {
		text, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSuffix(text, "\n"), " ");
		if len(s) == 0 || len(s) > 2 {
			fmt.Println("INVALID NUMBER OF ARGUMENTS. Must be either 1 or 2 numbers")
			continue
		}

		ballCount, err1 := strconv.ParseInt(s[0], 10, 64)
		if err1 != nil {
			fmt.Println("ERROR PARSING FIRST ARGUMENT. ERROR: " + err1.Error())
			continue
		}

		if ballCount > 127 || ballCount < 27 {
			fmt.Println("INVALID BALL COUNT. The first paramater must be between 27 and 127 inclusive.")
			continue
		}

		minCount := int64(-1)
		if len(s) == 2 {
			tmp, err2 := strconv.ParseInt(s[1], 10, 64)
			
			if err2 != nil {
				fmt.Println("ERROR PARSING SECOND ARGUMENT. ERROR: " + err2.Error())
				continue
			}
			minCount = tmp
		}

		c, err := clock.New(int(ballCount))
		if err != "" {
			fmt.Println("ERROR CREATING CLOCK. ERROR: " + err)
			continue
		}

		i, config := c.Run(minCount)
		if minCount == -1 {
			fmt.Printf("%d balls cycle after %d days\n", ballCount, i/1440)
		} else {
			data, _ := json.Marshal(config)
			fmt.Println(string(data))
		}
	}
}
