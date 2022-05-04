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

	var t, a, b int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		fmt.Fscan(rd, &a, &b)
		fmt.Fprintf(wr, "%d\n", (a/b)*(a/b))
	}
}
