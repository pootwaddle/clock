//Clock is the controlling framework to run the ball clock simulation.
package main

import (
	"clock/packages/ballclock"
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Println(`clock <int between 27 and 127> <int (optional, invokes mode 2)>

Clock is a simulation of a ball clock.  It works in one of two modes:
1 - Cycle Days.
Cycle Days expects one integer parameter between 27 and 127.
The simulation runs until the main queue of balls rotates around
to match the original order. This mode reports the number of days
the cycle would take, and the program execution time.
2 - Clock State.
Clock State expects two integer parameters. The first parameter,
like mode 1, is expected to be in the range of 27 to 127.
The second parameter is the number of "minutes" for the simulation
to run, and then report the state of each of the queues.`)
	os.Exit(-1)
}

func main() {
	start := time.Now()

	paramCount := len(os.Args)
	if paramCount < 2 || paramCount > 3 {
		usage()
	}

	ballcount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	var cycles int
	if paramCount == 3 {
		cycles, err = strconv.Atoi(os.Args[2])
		if err != nil {
			usage()
		}
	}

	if ballcount < 27 || ballcount > 127 || cycles < 0 {
		usage()
	}

	//Valid input, run our simulation
	if paramCount == 3 {
		t1 := ballclock.NewClock(ballcount)
		r := t1.RunMode2(cycles)
		fmt.Println(r)
	} else {
		t1 := ballclock.NewClock(ballcount)
		r := t1.RunMode1()
		fmt.Println(r)
		elapsed := time.Since(start)
		fmt.Printf("Completed in %d milliseconds (%1.3f seconds)", elapsed/time.Millisecond, elapsed.Seconds())
	}
}
