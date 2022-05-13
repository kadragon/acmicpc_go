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

	var a, b int64
	fmt.Fscan(rd, &a, &b)

	if a > b {
		a, b = b, a
	}

	fmt.Fprintf(wr, "%d\n", (a+b)*(b-a+1)/2)
}
