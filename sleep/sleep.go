package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	if len(os.Args) != 2 {
		usageError()
	}

	count, err := humanize.ParseBytes(os.Args[1])
	if err != nil {
		usageError()
	}

	time.Sleep(time.Duration(count) * time.Millisecond)
}

func usageError() {
	fmt.Fprintf(os.Stderr, "Usage: %s <count>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Sleep for <count> milliseconds.")
	os.Exit(-1)
}
