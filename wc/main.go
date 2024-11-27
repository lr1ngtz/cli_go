package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countlines bool, countbytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countlines {
		scanner.Split(bufio.ScanWords)
	}
	if countbytes {
		scanner.Split(bufio.ScanBytes)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	fmt.Println(*lines, *bytes)
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
