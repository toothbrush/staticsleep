package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Please provide a number of seconds to sleep.")
		os.Exit(1)
	}

	var delay = os.Args[1]
	i, err := strconv.Atoi(delay)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please give an integer-valued sleep delay.")
		os.Exit(1)
	}

	time.Sleep(time.Duration(i) * time.Second)
}
