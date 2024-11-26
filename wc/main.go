package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countlines bool) int {
	scanner := bufio.NewScanner(r)
	if !countlines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}
