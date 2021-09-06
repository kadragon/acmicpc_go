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

func f(a, b int) int {
	if a > b {
		return f(a/2, b)
	}
	if a < b {
		return f(a, b/2)
	}
	return a
}

func main() {
	defer wr.Flush()

	var n, a, b int
	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a, &b)
		fmt.Fprintf(wr, "%d\n", f(a, b)*10)
	}
}
