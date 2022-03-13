package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var x, y int
	fmt.Fscan(rd, &x, &y)

	fmt.Fprintf(wr, "%d\n", rev(rev(x)+rev(y)))
}

func rev(x int) int {
	t := 0
	for x > 0 {
		t *= 10
		t += x % 10
		x /= 10
	}

	return t
}
