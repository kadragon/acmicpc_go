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

	var a, b float64
	fmt.Fscan(rd, &a, &b)

	t := a * (100 - b) / 100

	if t >= 100 {
		fmt.Fprintf(wr, "0")
	} else {
		fmt.Fprintf(wr, "1")
	}
}
