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

	var a, b, c, x, y float64
	fmt.Fscan(rd, &a, &b, &c)
	x = a * b / c
	y = a / b * c

	if x > y {
		fmt.Fprintf(wr, "%d\n", int64(x))
	} else {
		fmt.Fprintf(wr, "%d\n", int64(y))
	}
}
