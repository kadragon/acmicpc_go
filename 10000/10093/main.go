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

	var a, b int
	fmt.Fscan(rd, &a, &b)

	if a > b {
		a, b = b, a
	}

	if a == b {
		fmt.Fprintf(wr, "0")
		return
	}

	fmt.Fprintf(wr, "%d\n", b-a-1)
	for i := a + 1; i < b; i++ {
		fmt.Fprintf(wr, "%d ", i)
	}
}
