package main

import (
	"os"
	"strconv"
	"time"
)

func printErr(err string) {
	os.Stderr.Write([]byte(err))
}

func main() {
	if len(os.Args) != 2 {
		printErr("Please provide a number of seconds to sleep.\n")
		os.Exit(1)
	}

	var delay = os.Args[1]
	i, err := strconv.Atoi(delay)
	if err != nil {
		printErr("Please give an integer-valued sleep delay.\n")
		os.Exit(1)
	}

	time.Sleep(time.Duration(i) * time.Second)
}
