# BallClock-golang
Ball Clock written in go

The simplist way to import it is by issuing the command:

	 go get github.com/Nais777/BallClock-golang

That will automatically load it into your src directory in your workspace.

You can then issue either of the following commands to run the program:

	1) go run $GOPATH/src/github.com/Nais777/BallClock-golang/ballclock.go

	2) go install $GOPATH/src/github.com/Nais777/BallClock-golang
	   $GOPATH/bin/BallClock-golang

The program expects an input in the form of a single number, or two numbers 
seperated by a space. The first number indicates how many balls should be in the clock,
while the second number indicates how many minutes the clock should run. The first
number in required while the second is optional.

In the case that only one number is entered, the program will output how many days
it takes for the balls to cycle so that they are in their original order.

In the case that the second number is supplied, the program will output what the current
ball configuration of the clock is.

To exit, type exit and press enter
