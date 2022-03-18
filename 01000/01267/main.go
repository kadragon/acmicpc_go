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

	var n, x, y int
	fmt.Fscan(rd, &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 0; i < n; i++ {
		x += (d[i]/30 + 1) * 10
		y += (d[i]/60 + 1) * 15
	}

	if x == y {
		fmt.Fprintf(wr, "Y M %d\n", x)
	} else if x < y {
		fmt.Fprintf(wr, "Y %d\n", x)
	} else {
		fmt.Fprintf(wr, "M %d\n", y)
	}
}
