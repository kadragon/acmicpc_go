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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(b, a int) int {
	switch {
	case b <= a:
		return a - b
	case b == 1:
		return 1
	case b%2 == 1:
		return 1 + min(f(b-1, a), f(b+1, a))
	default:
		return min(b-a, f(b/2, a))
	}
}

func main() {
	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	fmt.Fprintf(wr, "%d\n", f(k, n))
}
